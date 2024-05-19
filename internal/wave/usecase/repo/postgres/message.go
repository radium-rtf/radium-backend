package postgres

import (
	"context"
	"database/sql"
	"errors"

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

func (r Message) Create(ctx context.Context, message *entity.Message) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(message).Exec(ctx)
		if err != nil {
			return err
		}
		return err
	})
}

func (r Message) GetMessagesFrom(ctx context.Context, chatId string) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.db.NewSelect().Model(&messages).
		Where("chat_id = ?", chatId).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, nil
	}
	return messages, err
}

func (r Message) Get(ctx context.Context) (*entity.Message, error) {
	message := new(entity.Message)
	return message, nil
}
