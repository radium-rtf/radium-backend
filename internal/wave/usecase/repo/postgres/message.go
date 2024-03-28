package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Message struct {
	db *bun.DB
}

func NewMessageRepo(pg *postgres.Postgres) Message {
	return Message{db: pg.DB}
}

func (r Message) Create() error {
	return nil
}
