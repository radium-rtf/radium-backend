package repo

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
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
		Columns("name", "description", "chat", "logo", "type").
		Values(course.Name, course.Description, course.Chat, course.Logo, course.Type).
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
		return entity.Course{}, errors.New("курс не найден")
	}
	return courses[0], nil
}

func (r CourseRepo) GetCourses(ctx context.Context) ([]entity.Course, error) {
	return r.get(ctx, sq.Eq{})
}

func (r CourseRepo) get(ctx context.Context, where sq.Eq) ([]entity.Course, error) {
	courses := make([]entity.Course, 0)
	sql, args, err := r.pg.Builder.
		Select("id", "name", "description", "type", "chat", "logo").
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
		err = rows.Scan(&course.Id, &course.Name, &course.Description, &course.Type, &course.Chat, &course.Logo)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
