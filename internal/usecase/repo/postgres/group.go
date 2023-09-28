package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"

	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Group struct {
}

func NewGroupRepo(pg *postgres.Postgres) Group {
	return Group{}
}

func (r Group) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	panic("not implemented")
}

func (r Group) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	panic("not implemented")
}

func (r Group) GetByInviteCode(ctx context.Context, code string) (*entity.Group, error) {
	panic("not implemented")
}

func (r Group) JoinStudent(ctx context.Context, studentId uuid.UUID, code string) error {
	panic("not implemented")
}

func (r Group) Get(ctx context.Context) ([]*entity.Group, error) {
	panic("not implemented")
}
