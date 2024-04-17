package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Conference struct {
	db *bun.DB
}

func NewConferenceRepo(pg *postgres.Postgres) Conference {
	return Conference{db: pg.DB}
}

func (r Conference) Create() error {
	return nil
}
