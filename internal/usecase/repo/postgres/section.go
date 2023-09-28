package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"gorm.io/gen/field"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Section struct {
	pg *db.Query
}

func NewSectionRepo(pg *postgres.Postgres) Section {
	return Section{pg: pg.Q}
}

func (r Section) CreateSection(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	err := r.pg.Section.
		WithContext(ctx).
		Preload(field.Associations).
		Create(section)

	if err != nil {
		return nil, err
	}

	return section, nil
}

func (r Section) GetSectionById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	return r.pg.Section.
		WithContext(ctx).
		Preload(field.Associations).
		Where(r.pg.Section.Id.Eq(id)).
		First()
}

func (r Section) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	s := r.pg.Section.WithContext(ctx)
	if !isSoft {
		s = s.Unscoped()
	}
	info, err := s.Where(r.pg.Section.Id.Eq(id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
