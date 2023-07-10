package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/gen"
)

type CourseRepo struct {
	pg *db.Query
}

func NewCourseRepo(pg *db.Query) CourseRepo {
	return CourseRepo{pg: pg}
}

func (r CourseRepo) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
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

func (r CourseRepo) GetByName(ctx context.Context, name string) (*entity.Course, error) {
	courses, err := r.get(ctx, r.pg.Course.Name.Eq(name))
	if err != nil {
		return &entity.Course{}, err
	}
	if len(courses) == 0 {
		return &entity.Course{}, entity.ErrCourseNotFound
	}
	return courses[0], nil
}

func (r CourseRepo) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	return r.get(ctx, nil)
}

func (r CourseRepo) get(ctx context.Context, where ...gen.Condition) ([]*entity.Course, error) {
	c := r.pg.Course
	s := c.Modules.Pages.Sections
	return c.WithContext(ctx).Debug().
		Preload(c.Links, c.Authors).
		Preload(c.Modules.Pages.Sections, s.Order(r.pg.Section.Order)).
		Where(where...).
		Find()
}

func (r CourseRepo) GetById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	c := r.pg.Course
	sections := c.Modules.Pages.Sections
	course, err := c.WithContext(ctx).Debug().
		Preload(c.Links, c.Authors, c.Modules).
		Preload(sections, sections.Order(r.pg.Section.Order)).
		Where(c.Id.Eq(id)).Take()
	return course, err
}

func (r CourseRepo) GetFullById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	c := r.pg.Course
	s := c.Modules.Pages.Sections
	course, err := c.WithContext(ctx).Debug().
		Preload(c.Links, c.Authors, c.Modules.Pages).
		Preload(s, s.Order(r.pg.Section.Order)).
		Preload(s.ChoiceSection, s.MultiChoiceSection, s.TextSection, s.ShortAnswerSection, s.AnswerSection).
		Where(c.Id.Eq(id)).
		Take()
	return course, err
}

func (r CourseRepo) GetFullBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	c := r.pg.Course
	course, err := c.WithContext(ctx).Debug().
		Preload(c.Links, c.Authors, c.Modules.Pages).
		Where(c.Slug.Eq(slug)).
		Take()
	return course, err
}

func (r CourseRepo) Join(ctx context.Context, userId, courseId uuid.UUID) error {
	course := entity.Course{DBModel: entity.DBModel{Id: courseId}}
	user := &entity.User{DBModel: entity.DBModel{Id: userId}}
	err := r.pg.Course.Students.WithContext(ctx).Model(&course).Append(user)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", course)

	return nil
}

// func (r CourseRepo) CreateLink(ctx context.Context, link entity.Link) error {
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

// func (r CourseRepo) CreateCollaborator(ctx context.Context, collaborator entity.CourseCollaborator) error {
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

func (r CourseRepo) GetByStudent(ctx context.Context, userId uuid.UUID) ([]*entity.Course, error) {
	c := r.pg.Course
	u := r.pg.User
	courses, err := c.WithContext(ctx).Preload(c.Students.On(u.Id.Eq(userId))).Find()
	if err != nil {
		return []*entity.Course{}, err
	}

	return courses, nil
}

func (r CourseRepo) Delete(ctx context.Context, destroy *entity.Destroy) error {
	c := r.pg.Course.WithContext(ctx)
	if !destroy.IsSoft {
		c = c.Unscoped()
	}
	info, err := c.Where(r.pg.Course.Id.Eq(destroy.Id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
