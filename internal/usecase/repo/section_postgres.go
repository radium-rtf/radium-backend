package repo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gen/field"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type SectionRepo struct {
	pg *db.Query
}

func NewSectionRepo(pg *db.Query) SectionRepo {
	return SectionRepo{pg: pg}
}

func (r SectionRepo) CreateSection(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	err := r.pg.Section.
		WithContext(ctx).
		Preload(field.Associations).
		Create(section)

	if err != nil {
		return nil, err
	}

	return section, nil
}

func (r SectionRepo) GetSectionById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	return r.pg.Section.
		WithContext(ctx).
		Preload(field.Associations).
		Where(r.pg.Section.Id.Eq(id)).
		First()
}
