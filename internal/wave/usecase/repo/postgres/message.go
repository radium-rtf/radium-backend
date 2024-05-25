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
		return err
	})
}

func (r Message) GetMessagesFrom(ctx context.Context, chatId string) ([]*entity.Message, error) {
	var messageLinks []*entity.DialogueMessage
	err := r.db.NewSelect().Model(&messageLinks).
		Relation("Dialogue").
		Relation("Message").
		Relation("Message.Content").
		Where("dialogue_id = ?", chatId).
		Scan(ctx)
	messages := make([]*entity.Message, 0, len(messageLinks))
	for _, link := range messageLinks {
		message := link.Message
		messages = append(messages, message)
	}
	if errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}
	return messages, err
}

func (r Message) Get(ctx context.Context) (*entity.Message, error) {
	message := new(entity.Message)
	return message, nil
}
