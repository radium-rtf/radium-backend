package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Module struct {
	db *bun.DB
}

func NewModuleRepo(pg *postgres.Postgres) Module {
	return Module{db: pg.DB}
}

func (r Module) Create(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	_, err := r.db.NewInsert().Model(module).Exec(ctx)
	return module, err
}

func (r Module) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	var query = r.db.NewDelete().
		Model(&entity.Module{}).
		Where("id = ?", id)
	if !isSoft {
		query = query.ForceDelete()
	}
	_, err := query.Exec(ctx)
	return err
}

func (r Module) Update(ctx context.Context, module *entity.Module) (*entity.Module, error) {
	info, err := r.db.NewUpdate().
		Model(module).
		WherePK().
		OmitZero().
		Exec(ctx)

	n, _ := info.RowsAffected()
	if err == nil && n == 0 {
		return nil, repoerr.ModuleNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, module.Id)
}

func (r Module) GetById(ctx context.Context, id uuid.UUID) (*entity.Module, error) {
	var module = new(entity.Module)
	err := r.db.NewSelect().Model(module).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return module, repoerr.ModuleNotFound
	}
	return module, err
}
