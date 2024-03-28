package postgres

import (
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
