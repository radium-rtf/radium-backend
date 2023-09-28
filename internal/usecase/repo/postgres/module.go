package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Module struct {
}

func NewModuleRepo(pg *postgres.Postgres) Module {
	return Module{}
}

func (r Module) Create(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	panic("not implemented")
}

func (r Module) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	panic("not implemented")
}

func (r Module) Update(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	panic("not implemented")
}

func (r Module) GetById(ctx context.Context, id uuid.UUID) (*entity.Module, error) {
	panic("not implemented")
}
