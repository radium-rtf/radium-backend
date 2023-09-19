package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/utils"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Module struct {
	pg *db.Query
}

func NewModuleRepo(pg *db.Query) Module {
	return Module{pg: pg}
}

func (r Module) Create(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	err := r.pg.Module.WithContext(ctx).Create(module)
	return module, err
}

func (r Module) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	m := r.pg.Module.WithContext(ctx)
	if !isSoft {
		m = m.Unscoped()
	}
	info, err := m.Where(r.pg.Module.Id.Eq(id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}

func (r Module) Update(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	m := utils.RemoveEmptyFields(module)

	info, err := r.pg.Module.WithContext(ctx).
		Where(r.pg.Module.Id.Eq(module.Id)).
		Updates(m)
	if err != nil {
		return nil, err
	}

	if info.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return r.GetById(ctx, module.Id)
}

func (r Module) GetById(ctx context.Context, id uuid.UUID) (*entity.Module, error) {
	m := r.pg.Module
	return m.WithContext(ctx).
		Where(m.Id.Eq(id)).
		Preload(m.Pages).
		First()
}
