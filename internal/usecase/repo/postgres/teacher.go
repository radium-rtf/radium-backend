package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Teacher struct {
	db *bun.DB
}

func NewTeacherRepo(pg *postgres.Postgres) Teacher {
	return Teacher{db: pg.DB}
}

func (t Teacher) GetCoursesByTeacherId(ctx context.Context, id uuid.UUID) ([]*entity.TeacherCourseGroup, error) {
	var courses []*entity.TeacherCourseGroup
	err := t.db.NewSelect().
		Model(&courses).
		Where("user_id = ?", id).
		Relation("Group").Relation("Course").
		Scan(ctx)
	return courses, err
}

func (t Teacher) Create(ctx context.Context, courses []*entity.TeacherCourseGroup) ([]*entity.TeacherCourseGroup, error) {
	_, err := t.db.NewInsert().Model(&courses).Exec(ctx)
	return courses, err
}
