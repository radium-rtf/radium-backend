package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
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
		_, err := tx.NewDelete().Model(&entity.DialogueMessage{MessageId: messageId}).
			Where("message_id = ?", messageId).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.NewDelete().Model(&entity.DialoguePinnedMessage{MessageId: messageId}).
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

func (r Message) GetDialogueFromMessage(ctx context.Context, messageId uuid.UUID) *model.Chat {
	dialogueLink := new(entity.DialogueMessage)
	err := r.db.NewSelect().Model(dialogueLink).
		Relation("Dialogue").
		Where("message_id = ?", messageId).
		Scan(ctx)
	if err != nil {
		return nil
	}
	dialogue := dialogueLink.Dialogue
	chat := model.NewChat(dialogue.Id, dialogue.Id.String(), "dialogue", nil)
	return &chat
}

func (r Message) Get(ctx context.Context, id uuid.UUID) (*entity.Message, *model.Chat, error) {
	message := new(entity.Message)
	err := r.db.NewSelect().Model(message).
		Relation("Content").
		Where("\"message\".\"id\" = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, nil, err
	}
	chat := r.GetDialogueFromMessage(ctx, id)
	if chat == nil {
		return message, nil, fmt.Errorf("chat not found")
	}
	return message, chat, err
}

func (r Message) Pin(ctx context.Context, messageId, chatId uuid.UUID, chatType string, status bool) error {
	dialoguePinned := new(entity.DialoguePinnedMessage)

	if chatType != "dialogue" && chatType != "" { // TODO: разделить получше
		return fmt.Errorf("unsupported chat type: %s", chatType)
	}

	var err error
	if status {
		_, err = r.db.NewInsert().Model(&entity.DialoguePinnedMessage{
			DialogueId: chatId,
			MessageId:  messageId,
		}).Exec(ctx)
	} else {
		_, err = r.db.NewDelete().Model(dialoguePinned).
			Where("message_id = ?", messageId).
			Exec(ctx)
	}

	return err
}

func (r Message) IsPinned(ctx context.Context, messageId uuid.UUID, chatType string) (bool, error) {
	if chatType != "dialogue" && chatType != "" { // TODO: разделить получше
		return false, fmt.Errorf("unsupported chat type: %s", chatType)
	}

	dialoguePinned := new(entity.DialoguePinnedMessage)
	err := r.db.NewSelect().Model(dialoguePinned).
		Where("message_id = ?", messageId).
		Scan(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
