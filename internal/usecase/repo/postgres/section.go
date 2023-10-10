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

type Section struct {
	db *bun.DB
}

func NewSectionRepo(pg *postgres.Postgres) Section {
	return Section{db: pg.DB}
}

func (r Section) Create(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	_, err := r.db.NewInsert().Model(section).Exec(ctx)
	return section, err
}

func (r Section) GetById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	var section = new(entity.Section)
	err := r.db.NewSelect().Model(section).Where("id = ?", id).Scan(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
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
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		info, err := r.db.NewUpdate().
			Model(section).
			WherePK().
			OmitZero().
			Exec(ctx)

		n, _ := info.RowsAffected()
		if err == nil && n == 0 {
			return repoerr.SectionNotFound
		}

		if section.Answer == "" && len(section.Answers) == 0 {
			return nil
		}
		_, err = r.db.NewDelete().
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
