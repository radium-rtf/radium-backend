package postgres

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Dialogue struct {
	db *bun.DB
}

func NewDialogueRepo(pg *postgres.Postgres) Dialogue {
	return Dialogue{db: pg.DB}
}

func (r Dialogue) Create() error {
	return nil
}

func (r Dialogue) Get(ctx context.Context) (*entity.Dialogue, error) {
	dialogue := new(entity.Dialogue)
	return dialogue, nil
}
