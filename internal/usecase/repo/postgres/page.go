package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Page struct {
}

func NewPageRepo(pg *postgres.Postgres) Page {
	return Page{}
}

func (r Page) Create(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	panic("not implemented")
}

func (r Page) GetByID(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	panic("not implemented")
}

func (r Page) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	panic("not implemented")
}

func (r Page) Update(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	panic("not implemented")
}
