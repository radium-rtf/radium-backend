package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/gen"
)

type Course struct {
	pg *db.Query
}

func NewCourseRepo(pg *db.Query) Course {
	return Course{pg: pg}
}

func (r Course) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	err := r.pg.Course.WithContext(ctx).Preload(r.pg.Course.Authors).Create(course)
	if err != nil {
		return &entity.Course{}, err
	}
	course, err = r.GetFullById(ctx, course.Id)
	if err != nil {
		return &entity.Course{}, err
	}
	return course, err
}

func (r Course) GetByName(ctx context.Context, name string) (*entity.Course, error) {
	courses, err := r.get(ctx, r.pg.Course.Name.Eq(name))
	if err != nil {
		return &entity.Course{}, err
	}
	if len(courses) == 0 {
		return &entity.Course{}, entity.ErrCourseNotFound
	}
	return courses[0], nil
}

func (r Course) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	return r.get(ctx, nil)
}

// TODO: спросить про необходимость показывать курс полностью в списке всех курсов
func (r Course) get(ctx context.Context, where ...gen.Condition) ([]*entity.Course, error) {
	c := r.pg.Course
	return c.WithContext(ctx).Debug().
		Preload(c.Links, c.Authors).
		Preload(c.Modules).
		Where(where...).
		Find()
}

func (r Course) GetById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	c := r.pg.Course
	sections := c.Modules.Pages.Sections
	course, err := c.WithContext(ctx).Debug().
		Preload(c.Modules, c.Modules.Order(r.pg.Module.Order)).
		Preload(c.Links, c.Authors).
		Preload(sections, sections.Order(r.pg.Section.Order)).
		Where(c.Id.Eq(id)).Take()

	return course, err
}

func (r Course) GetFullById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	return r.getFull(ctx, r.pg.Course.Id.Eq(id))
}

func (r Course) GetFullBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	return r.getFull(ctx, r.pg.Course.Slug.Eq(slug))
}

func (r Course) getFull(ctx context.Context, where ...gen.Condition) (*entity.Course, error) {
	c := r.pg.Course
	m := c.Modules
	p := m.Pages
	s := p.Sections

	course, err := c.WithContext(ctx).Debug().
		Where(where...).
		Preload(c.Links, c.Authors).
		Preload(m, m.Order(r.pg.Module.Order)).
		Preload(p, p.Order(r.pg.Page.Order)).
		Preload(s, s.Order(r.pg.Section.Order)).
		Preload(s.ChoiceSection, s.MultiChoiceSection, s.TextSection, s.ShortAnswerSection, s.AnswerSection).
		First()

	return course, err
}

func (r Course) Join(ctx context.Context, userId, courseId uuid.UUID) error {
	course := entity.Course{DBModel: entity.DBModel{Id: courseId}}
	user := &entity.User{DBModel: entity.DBModel{Id: userId}}
	err := r.pg.Course.Students.WithContext(ctx).Model(&course).Append(user)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", course)

	return nil
}

// func (r Course) CreateLink(ctx context.Context, link entity.Link) error {
// 	sql, args, err := r.pg.Builder.
// 		Insert("course_links").
// 		Columns("id", "name", "link", "course_id").
// 		Values(link.Id, link.Name, link.Link, link.CourseId).
// 		ToSql()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = r.pg.Pool.Exec(ctx, sql, args...)
// 	return err
// }

// func (r Course) CreateCollaborator(ctx context.Context, collaborator entity.CourseCollaborator) error {
// 	sql, args, err := r.pg.Builder.
// 		Insert("course_collaborators").
// 		Columns("id", "user_email", "course_id").
// 		Values(collaborator.Id, collaborator.UserEmail, collaborator.CourseId).
// 		ToSql()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = r.pg.Pool.Exec(ctx, sql, args...)
// 	return err
// }

func (r Course) GetByStudent(ctx context.Context, userId uuid.UUID) ([]*entity.Course, error) {
	c := r.pg.Course
	u := r.pg.User
	courses, err := c.WithContext(ctx).Preload(c.Students.On(u.Id.Eq(userId))).Find()
	if err != nil {
		return []*entity.Course{}, err
	}

	return courses, nil
}

func (r Course) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	c := r.pg.Course.WithContext(ctx)
	if !isSoft {
		c = c.Unscoped()
	}
	info, err := c.Where(r.pg.Course.Id.Eq(id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}

func (r Course) Update(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	m := structs.Map(course)
	delete(m, "DBModel")

	c := r.pg.Course
	info, err := c.WithContext(ctx).Where(c.Id.Eq(course.Id)).Updates(m)
	if err != nil {
		return nil, err
	}

	if info.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return r.GetById(ctx, course.Id)
}
