package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Chat struct {
	db *bun.DB
}

func NewChatRepo(pg *postgres.Postgres) Chat {
	return Chat{db: pg.DB}
}

func (r Chat) Create(ctx context.Context, chat *entity.Chat) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(chat).Exec(ctx)
		return err
	})
}

func (r Chat) Get(ctx context.Context, chatId uuid.UUID) (*entity.Chat, error) {
	var chat entity.Chat
	err := r.db.NewSelect().Model(&chat).
		Where("id = ?", chatId).
		Scan(ctx)
	return &chat, err
}

func (r Chat) GetByMessageId(ctx context.Context, messageId uuid.UUID) (*entity.Chat, error) {
	var chatLink entity.ChatMessage
	err := r.db.NewSelect().Model(&chatLink).
		Relation("Chat").
		Where("message_id = ?", messageId).
		Scan(ctx)
	return chatLink.Chat, err
}

func (r Chat) GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.Chat, error) {
	var chats []*entity.Chat
	err := r.db.NewSelect().Model(&chats).
		Relation("Dialogue", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("dialogue.first_user_id = ? OR dialogue.second_user_id = ?", userId, userId)
		}).
		Scan(ctx)
	return chats, err
}
