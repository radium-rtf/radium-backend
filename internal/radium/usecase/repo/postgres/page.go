package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Page struct {
	db *bun.DB
}

func NewPageRepo(pg *postgres.Postgres) Page {
	return Page{db: pg.DB}
}

func (r Page) Create(ctx context.Context, page *entity2.Page) (*entity2.Page, error) {
	_, err := r.db.NewInsert().Model(page).Exec(ctx)
	return page, err
}

func (r Page) GetById(ctx context.Context, id uuid.UUID) (*entity2.Page, error) {
	return r.get(ctx, columnValue{column: "id", value: id})
}

func (r Page) GetBySlug(ctx context.Context, slug string) (*entity2.Page, error) {
	return r.get(ctx, columnValue{column: "slug", value: slug})
}

func (r Page) get(ctx context.Context, where columnValue) (*entity2.Page, error) {
	var page = new(entity2.Page)
	err := r.getPageQuery(page).
		Where(where.column+" = ?", where.value).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return page, repoerr.NotFound
	}
	return page, err
}

func (r Page) getPageQuery(page *entity2.Page) *bun.SelectQuery {
	return r.db.NewSelect().Model(page).
		Relation("Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Sections.File")
}

func (r Page) GetByIdWithUserAnswers(ctx context.Context, id, userId uuid.UUID) (*entity2.Page, error) {
	var page = new(entity2.Page)
	err := r.getPageQuery(page).
		Where("id = ?", id).
		Relation("Sections.UsersAnswers", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("answer.user_id = ?", userId).Order("answer.created_at desc")
		}).
		Relation("Sections.UsersAnswers.Review").
		Relation("Sections.UsersAnswers.Review.Reviewer").
		Relation("Sections.UsersAnswers.File").
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return page, repoerr.NotFound
	}
	return page, err
}

func (r Page) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	var query = r.db.NewDelete().
		Model(&entity2.Page{}).
		Where("id = ?", id)
	if !isSoft {
		query = query.ForceDelete()
	}
	_, err := query.Exec(ctx)
	return err
}

func (r Page) Update(ctx context.Context, page *entity2.Page) (*entity2.Page, error) {
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
		return nil, repoerr.NotFound
	}
	return r.GetById(ctx, page.Id)
}

func (r Page) GetCourseByPageId(ctx context.Context, id uuid.UUID) (*entity2.Course, error) {
	var page = new(entity2.Page)
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

func (r Page) GetLastPage(ctx context.Context, moduleId uuid.UUID) (*entity2.Page, error) {
	var page = new(entity2.Page)
	err := r.db.NewSelect().Model(page).
		Where("module_id = ?", moduleId).
		Order("order desc").
		Limit(1).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.NotFound
	}
	return page, err
}

func (r Page) GetModulesByPageId(ctx context.Context, id uuid.UUID) ([]*entity2.Module, error) {
	var page = new(entity2.Page)
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

func (r Page) UpdateOrder(ctx context.Context, page *entity2.Page, order uint) (*entity2.Page, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		where := "page.order >= ? and page.order < ? and page.module_id = ?"
		set := "\"order\" = page.order + 1"
		if page.Order < float64(order) {
			where = "page.order <= ? and page.order > ? and page.module_id = ?"
			set = "\"order\" = page.order - 1"
		}

		_, err := tx.NewUpdate().
			Model(&entity2.Page{}).
			Where(where, order, page.Order, page.ModuleId).
			Set(set).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewUpdate().
			Model(&entity2.Page{}).
			Where("uuid_eq(page.id, ?)", page.Id).
			Set("\"order\" = ?", order).
			Exec(ctx)

		return err
	})

	page.Order = float64(order)
	return page, err
}
