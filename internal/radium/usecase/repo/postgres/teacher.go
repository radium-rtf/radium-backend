package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Teacher struct {
	db *bun.DB
}

func NewTeacherRepo(pg *postgres.Postgres) Teacher {
	return Teacher{db: pg.DB}
}

func (t Teacher) GetCoursesByTeacherId(ctx context.Context, id uuid.UUID) ([]*entity.Teacher, error) {
	var courses []*entity.Teacher
	err := t.db.NewSelect().
		Model(&courses).
		Where("user_id = ?", id).
		Relation("Group").
		Relation("Course").
		Scan(ctx)
	return courses, err
}

func (t Teacher) Create(ctx context.Context, courses []*entity.Teacher) ([]*entity.Teacher, error) {
	_, err := t.db.NewInsert().Model(&courses).Exec(ctx)
	return courses, err
}
