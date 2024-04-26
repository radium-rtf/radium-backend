package postgres

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/wave/entity"
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

func (r Message) Get(ctx context.Context) (*entity.Message, error) {
	message := new(entity.Message)
	return message, nil
}
