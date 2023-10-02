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
