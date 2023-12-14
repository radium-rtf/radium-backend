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
	err := r.getPageByIdQuery(page).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return page, repoerr.PageNotFound
	}
	return page, err
}

func (r Page) getPageByIdQuery(page *entity.Page) *bun.SelectQuery {
	return r.db.NewSelect().Model(page).
		Relation("Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		})
}

func (r Page) GetByIdWithUserAnswers(ctx context.Context, id, userId uuid.UUID) (*entity.Page, error) {
	var page = new(entity.Page)
	err := r.getPageByIdQuery(page).
		Where("id = ?", id).
		Relation("Sections.UsersAnswers", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("answer.user_id = ?", userId).Order("answer.created_at desc")
		}).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
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

func (r Page) GetCourseByPageId(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	var page = new(entity.Page)
	err := r.db.NewSelect().Model(page).
		Where("page.id = ?", id).
		Relation("Module").
		Relation("Module.Course").
		Relation("Module.Course.Authors").
		Relation("Module.Course.Coauthors").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return page.Module.Course, nil
}

func (r Page) GetLastPage(ctx context.Context, moduleId uuid.UUID) (*entity.Page, error) {
	var page = new(entity.Page)
	err := r.db.NewSelect().Model(page).
		Where("module_id = ?", moduleId).
		Order("order desc").
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.PageNotFound
	}
	return page, err
}

func (r Page) GetModulesByPageId(ctx context.Context, id uuid.UUID) ([]*entity.Module, error) {
	var page = new(entity.Page)
	err := r.db.NewSelect().Model(page).
		Where("page.id = ?", id).
		Relation("Module").
		Relation("Module.Course").
		Relation("Module.Course.Modules", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Module.Course.Modules.Pages", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return page.Module.Course.Modules, nil
}

func (r Page) UpdateOrder(ctx context.Context, page *entity.Page, order uint) (*entity.Page, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		where := "page.order >= ? and page.order <= ?"
		set := "\"order\" = page.order + 1"
		if page.Order < float64(order) {
			where = "page.order <= ? and page.order >= ?"
			set = "\"order\" = page.order - 1"
		}

		_, err := tx.NewUpdate().
			Model(&entity.Page{}).
			Where(where, order, page.Order).
			Set(set).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewUpdate().
			Model(&entity.Page{}).
			Where("uuid_eq(page.id, ?)", page.Id).
			Set("\"order\" = ?", order).
			Exec(ctx)

		return err
	})

	page.Order = float64(order)
	return page, err
}
