package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
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
	if err != nil {
		return nil, err
	}

	n, _ := info.RowsAffected()
	if n == 0 {
		return nil, repoerr.ModuleNotFound
	}
	return r.GetById(ctx, module.Id)
}

func (r Module) GetById(ctx context.Context, id uuid.UUID) (*entity.Module, error) {
	var module = new(entity.Module)
	err := r.db.NewSelect().
		Model(module).
		Where("id = ?", id).
		Relation("Pages", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Pages.Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return module, repoerr.ModuleNotFound
	}
	return module, err
}

func (r Module) GetCourseByModuleId(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	var module = new(entity.Module)
	err := r.db.NewSelect().
		Model(module).
		Where("module.id = ?", id).
		Relation("Course").
		Relation("Course.Authors").
		Relation("Course.Coauthors").
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.ModuleNotFound
	}
	if err != nil {
		return nil, err
	}
	return module.Course, nil
}

func (r Module) GetLastModule(ctx context.Context, courseId uuid.UUID) (*entity.Module, error) {
	var module = new(entity.Module)
	err := r.db.NewSelect().Model(module).
		Where("course_id = ?", courseId).
		Order("order desc").
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.ModuleNotFound
	}
	return module, err
}
