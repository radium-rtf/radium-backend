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
		if err != nil {
			return err
		}
		_, err = tx.NewDelete().Model(&entity.Message{DBModel: entity.DBModel{Id: messageId}}).
			Where("id = ?", messageId).
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
	dialogueMessage := new(entity.ChatMessage)

	_, err := r.db.NewUpdate().Model(dialogueMessage).
		Set("is_pinned = ?", status).
		Where("message_id = ? AND chat_id = ?", messageId, chatId).
		Exec(ctx)

	return err
}

func (r Message) IsPinned(ctx context.Context, messageId uuid.UUID) (bool, error) {
	dialoguePinned := new(entity.ChatMessage)
	err := r.db.NewSelect().Model(dialoguePinned).
		Where("message_id = ? AND is_pinned = ?", messageId, true).
		Scan(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r Message) LinkToChat(ctx context.Context, messageId uuid.UUID, chatId uuid.UUID) error {
	chatMessage := entity.ChatMessage{
		ChatId:    chatId,
		MessageId: messageId,
	}

	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(&chatMessage).Exec(ctx)
		return err
	})
}
