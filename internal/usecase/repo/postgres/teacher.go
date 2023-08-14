package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Teacher struct {
	pg *db.Query
}

func NewTeacherRepo(pg *db.Query) Teacher {
	return Teacher{pg: pg}
}

func (t Teacher) GetByUserId(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	teacher := t.pg.Teacher
	courses := teacher.Courses
	return teacher.WithContext(ctx).
		Preload(courses, courses.Course, courses.Group).
		Where(teacher.UserId.Eq(id)).
		First()
}

func (t Teacher) Create(ctx context.Context, teacher *entity.Teacher) (*entity.Teacher, error) {
	courses := t.pg.Teacher.Courses
	err := t.pg.Teacher.WithContext(ctx).
		Preload(courses, courses.Group, courses.Course).Create(teacher)
	if err != nil {
		return nil, err
	}
	return t.GetByUserId(ctx, teacher.UserId)
}
