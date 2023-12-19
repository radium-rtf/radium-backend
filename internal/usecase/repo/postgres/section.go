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
	"slices"
)

type Section struct {
	db     *bun.DB
	answer Answer
}

func NewSectionRepo(pg *postgres.Postgres) Section {
	return Section{db: pg.DB, answer: NewAnswerRepo(pg)}
}

func (r Section) Create(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	_, err := r.db.NewInsert().Model(section).Exec(ctx)
	return section, err
}

func (r Section) GetById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	var section = new(entity.Section)
	err := r.db.NewSelect().Model(section).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.SectionNotFound
	}
	return section, err
}

func (r Section) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	var query = r.db.NewDelete().
		Model(&entity.Section{}).
		Where("id = ?", id)
	if !isSoft {
		query = query.ForceDelete()
	}
	_, err := query.Exec(ctx)
	return err
}

func (r Section) Update(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	before, err := r.GetById(ctx, section.Id)
	if err != nil {
		return nil, err
	}
	err = r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		info, err := tx.NewUpdate().
			Model(section).
			WherePK().
			OmitZero().
			Exec(ctx)
		if err != nil {
			return err
		}

		n, _ := info.RowsAffected()
		if n == 0 {
			return repoerr.SectionNotFound
		}

		if section.Answer == "" && section.Answers == nil {
			return nil
		}
		if slices.Equal(before.Answers, section.Answers) {
			return nil
		}
		_, err = tx.NewDelete().
			Model(&entity.Answer{}).
			Where("section_id = ?", section.Id).
			Exec(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, section.Id)
}

func (r Section) GetByAnswerId(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	answer, err := r.answer.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return answer.Section, nil
}

func (r Section) GetCourseBySectionId(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	var section = new(entity.Section)
	err := r.db.NewSelect().
		Model(section).
		Where("section.id = ?", id).
		Relation("Page").
		Relation("Page.Module").
		Relation("Page.Module.Course").
		Relation("Page.Module.Course.Authors").
		Relation("Page.Module.Course.Coauthors").
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.SectionNotFound
	}
	if err != nil {
		return nil, err
	}
	return section.Page.Module.Course, err
}

func (r Section) GetLastSection(ctx context.Context, pageId uuid.UUID) (*entity.Section, error) {
	var section = new(entity.Section)
	err := r.db.NewSelect().Model(section).
		Where("page_id = ?", pageId).
		Order("order desc").
		Limit(1).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.SectionNotFound
	}
	return section, err
}

func (r Section) UpdateOrder(ctx context.Context, section *entity.Section, order uint) (*entity.Section, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		where := "section.order >= ? and section.order <= ?"
		set := "\"order\" = section.order + 1"
		if section.Order < float64(order) {
			where = "section.order <= ? and section.order >= ?"
			set = "\"order\" = section.order - 1"
		}
		_, err := tx.NewUpdate().
			Model(&entity.Section{}).
			Where(where, order, section.Order).
			Set(set).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewUpdate().
			Model(&entity.Section{}).
			Where("uuid_eq(section.id, ?)", section.Id).
			Set("\"order\" = ?", order).
			Exec(ctx)

		return err
	})

	section.Order = float64(order)
	return section, err
}
