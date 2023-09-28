package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Section struct {
}

func NewSectionRepo(pg *postgres.Postgres) Section {
	return Section{}
}

func (r Section) CreateSection(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	panic("not implemented")
}

func (r Section) GetSectionById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	panic("not implemented")
}

func (r Section) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	panic("not implemented")
}
