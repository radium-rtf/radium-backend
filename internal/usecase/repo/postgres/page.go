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

type Page struct {
	db *bun.DB
}

func NewPageRepo(pg *postgres.Postgres) Page {
	return Page{db: pg.DB}
}

func (r Page) Create(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	_, err := r.db.NewInsert().Model(page).Exec(ctx)
	return page, err
}

func (r Page) GetById(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	var page = new(entity.Page)
	err := r.db.NewSelect().Model(page).
		Where("id = ?", id).
		Relation("Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Scan(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return page, repoerr.PageNotFound
	}
	return page, err
}

func (r Page) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	var query = r.db.NewDelete().
		Model(&entity.Page{}).
		Where("id = ?", id)
	if !isSoft {
		query = query.ForceDelete()
	}
	_, err := query.Exec(ctx)
	return err
}

func (r Page) Update(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	info, err := r.db.NewUpdate().
		Model(page).
		WherePK().
		OmitZero().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	n, _ := info.RowsAffected()
	if n == 0 {
		return nil, repoerr.PageNotFound
	}
	return r.GetById(ctx, page.Id)
}
