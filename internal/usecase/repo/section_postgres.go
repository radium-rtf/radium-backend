package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (r SectionRepo) Delete(ctx context.Context, destroy *entity.Destroy) error {
	s := r.pg.Section.WithContext(ctx)
	if !destroy.IsSoft {
		s = s.Unscoped()
	}
	info, err := s.Where(r.pg.Section.Id.Eq(destroy.Id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
