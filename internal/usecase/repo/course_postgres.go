package repo

import (
	"context"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"strings"
)

type CourseRepo struct {
	pg *postgres.Postgres
}

func NewCourseRepo(pg *postgres.Postgres) CourseRepo {
	return CourseRepo{pg: pg}
}

func (r CourseRepo) Create(ctx context.Context, course entity.Course) error {
	sql, args, err := r.pg.Builder.
		Insert("courses").
		Columns("name", "description", "logo", "type", "author_id").
		Values(course.Name, course.Description, course.Logo, course.Type, course.AuthorId).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r CourseRepo) GetByName(ctx context.Context, name string) (entity.Course, error) {
	courses, err := r.get(ctx, sq.Eq{"name": name})
	if err != nil {
		return entity.Course{}, err
	}
	if len(courses) == 0 {
		return entity.Course{}, entity.CourseNotFoundErr
	}
	return courses[0], nil
}

func (r CourseRepo) GetCourses(ctx context.Context) ([]entity.Course, error) {
	return r.get(ctx, sq.Eq{})
}

func (r CourseRepo) get(ctx context.Context, where sq.Eq) ([]entity.Course, error) {
	courses := make([]entity.Course, 0)
	sql, args, err := r.pg.Builder.
		Select("id", "name", "description", "type", "logo", "author_id").
		Where(where).
		From("courses").
		ToSql()
	if err != nil {
		return courses, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return courses, err
	}
	defer rows.Close()
	for rows.Next() {
		var course entity.Course
		err = rows.Scan(&course.Id, &course.Name, &course.Description, &course.Type, &course.Logo, &course.AuthorId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r CourseRepo) GetFullById(ctx context.Context, id int) (entity.CourseTitle, error) {
	var courseTitle entity.CourseTitle
	sql, args, err := r.pg.Builder.
		Select("row_to_json(row)").
		Where(sq.Eq{"id": id}).
		Limit(1).From("courses_title_view as row").
		ToSql()
	if err != nil {
		return courseTitle, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return courseTitle, err
	}
	defer rows.Close()
	if !rows.Next() {
		return courseTitle, entity.CourseNotFoundErr
	}
	courseJson := ""
	err = rows.Scan(&courseJson)
	if err != nil {
		return entity.CourseTitle{}, err
	}
	return courseTitle, json.NewDecoder(strings.NewReader(courseJson)).Decode(&courseTitle)
}

func (r CourseRepo) CreateLink(ctx context.Context, link entity.Link) error {
	sql, args, err := r.pg.Builder.
		Insert("course_links").
		Columns("id", "name", "link", "course_id").
		Values(link.Id, link.Name, link.Link, link.CourseId).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r CourseRepo) CreateCollaborator(ctx context.Context, collaborator entity.CourseCollaborator) error {
	sql, args, err := r.pg.Builder.
		Insert("course_collaborators").
		Columns("id", "user_email", "course_id").
		Values(collaborator.Id, collaborator.UserEmail, collaborator.CourseId).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r CourseRepo) GetByStudent(ctx context.Context, userId string) ([]entity.Course, error) {
	courses := make([]entity.Course, 0)
	sql, args := r.pg.Builder.
		Select("courses.id", "name", "description", "type", "logo", "author_id").
		From("course_student").
		Where(sq.Eq{"user_id": userId}).
		Join("courses on courses.id = course_student.course_id").MustSql()

	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return courses, err
	}
	defer rows.Close()
	for rows.Next() {
		var course entity.Course
		err = rows.Scan(&course.Id, &course.Name, &course.Description, &course.Type, &course.Logo, &course.AuthorId)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, rows.Err()
}

func (r CourseRepo) Join(ctx context.Context, userId, courseId string) error {
	sql, args := r.pg.Builder.Insert("course_student").
		Columns("user_id", "course_id").
		Values(userId, courseId).MustSql()

	_, err := r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r CourseRepo) GetById(ctx context.Context, id string) (entity.Course, error) {
	courses, err := r.get(ctx, sq.Eq{"id": id})
	if err != nil {
		return entity.Course{}, err
	}
	if len(courses) == 0 {
		return entity.Course{}, entity.CourseNotFoundErr
	}
	return courses[0], nil
}
