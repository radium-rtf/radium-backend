package postgres

import (
	"context"
	"fmt"

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

// available sorters dict
var sorters = map[string]string{
	"date": "message.created_at",
}

// available orders dict
var orders = map[string]string{
	"asc":  "ASC",
	"desc": "DESC",
}

func (r Message) GetMessagesFrom(
	ctx context.Context, chatId string, page, pageSize int, sort, order string,
) ([]*entity.Message, error) {
	sortParam, ok := sorters[sort]
	if !ok {
		return nil, fmt.Errorf("invalid sort parameter: %s", sort)
	}

	sortOrder, ok := orders[order]
	if !ok {
		return nil, fmt.Errorf("invalid order parameter: %s", order)
	}

	var messageLinks []*entity.DialogueMessage
	err := r.db.NewSelect().Model(&messageLinks).
		Relation("Dialogue").
		Relation("Message").
		Relation("Message.Content").
		Where("dialogue_id = ?", chatId).
		Order(sortParam + " " + sortOrder).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, 0, len(messageLinks))
	for _, link := range messageLinks {
		message := link.Message
		messages = append(messages, message)
	}

	return messages, err
}

func (r Message) Get(ctx context.Context) (*entity.Message, error) {
	message := new(entity.Message)
	return message, nil
}
