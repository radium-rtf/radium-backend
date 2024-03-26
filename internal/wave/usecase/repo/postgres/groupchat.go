package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type GroupChat struct {
	db *bun.DB
}

func NewGroupChatRepo(pg *postgres.Postgres) GroupChat {
	return GroupChat{db: pg.DB}
}

func (r GroupChat) Create() error {
	return nil
}
