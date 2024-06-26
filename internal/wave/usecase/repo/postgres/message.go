package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (r Message) Delete(ctx context.Context, messageId uuid.UUID) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().Model(&entity.ChatMessage{MessageId: messageId}).
			Where("message_id = ?", messageId).
			Exec(ctx)
		return err
	})
}

var sorters = map[string]string{
	"date": "message.created_at",
}
var orders = map[string]string{
	"asc":  "ASC",
	"desc": "DESC",
}

func (r Message) GetMessagesFrom(
	ctx context.Context,
	chatId string,
	page, pageSize int,
	sort, order string,
) ([]*entity.Message, error) {
	sortParam, ok := sorters[sort]
	if !ok {
		return nil, fmt.Errorf("invalid sort parameter: %s", sort)
	}

	sortOrder, ok := orders[order]
	if !ok {
		return nil, fmt.Errorf("invalid order parameter: %s", order)
	}

	var messageLinks []*entity.ChatMessage
	err := r.db.NewSelect().Model(&messageLinks).
		Relation("Chat").
		Relation("Message").
		Relation("Message.Content").
		Where("chat_id = ?", chatId).
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

func (r Message) GetChatFromMessage(ctx context.Context, messageId uuid.UUID) *entity.Chat {
	chatLink := new(entity.ChatMessage)
	err := r.db.NewSelect().Model(chatLink).
		Relation("Chat").
		Where("message_id = ?", messageId).
		Scan(ctx)
	if err != nil {
		return nil
	}
	return chatLink.Chat
}

func (r Message) Get(ctx context.Context, id uuid.UUID) (*entity.Message, *entity.Chat, error) {
	messageLink := new(entity.ChatMessage)
	err := r.db.NewSelect().Model(messageLink).
		Relation("Chat").
		Relation("Message").
		Relation("Message.Content").
		Where("message_id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, nil, err
	}
	return messageLink.Message, messageLink.Chat, nil
}

func (r Message) Pin(ctx context.Context, messageId, chatId uuid.UUID, chatType string, status bool) error {
	message := new(entity.Message)

	_, err := r.db.NewUpdate().Model(message).
		Set("is_pinned = ?", status).
		Where("id = ?", messageId, chatId).
		Exec(ctx)

	return err
}

func (r Message) GetPinnedMessages(ctx context.Context, chatId uuid.UUID) ([]*entity.Message, error) {
	pinnedMessages := make([]*entity.Message, 0)
	err := r.db.NewSelect().Model(&pinnedMessages).
		Relation("Content").
		Join("JOIN wave.chat_message ON wave.chat_message.message_id = message.id").
		Where("wave.chat_message.chat_id = ? AND message.is_pinned = ?", chatId, true).
		Order("message.created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return pinnedMessages, nil
}

func (r Message) LinkToChat(ctx context.Context, messageId uuid.UUID, chatId uuid.UUID) error {
	chatMessage := entity.ChatMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}
	_, err := r.db.NewInsert().Model(&chatMessage).Exec(ctx)
	return err
}
