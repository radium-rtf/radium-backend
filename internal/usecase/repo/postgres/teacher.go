package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Teacher struct {
}

func NewTeacherRepo(pg *postgres.Postgres) Teacher {
	return Teacher{}
}

func (t Teacher) GetByUserId(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	panic("not implemented")
}

func (t Teacher) Create(ctx context.Context, teacher *entity.Teacher) (*entity.Teacher, error) {
	panic("not implemented")
}
