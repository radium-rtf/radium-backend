package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Chat struct {
		bun.BaseModel `bun:"table:wave.chats"`

		Id   uuid.UUID `bun:",pk"`
		Name string    `bun:",notnull"`
		Type string    `bun:",notnull"`

		Dialogue *Dialogue `bun:"rel:belongs-to,join:id=id"`
		Group    *Group    `bun:"rel:belongs-to,join:id=id"`

		Messages []*Message `bun:"m2m:wave.chat_message,join:Chat=Message"`
	}

	ChatMessage struct {
		bun.BaseModel `bun:"table:wave.chat_message"`

		ChatId uuid.UUID `bun:",pk"`
		Chat   *Chat     `bun:"rel:belongs-to,join:chat_id=id"`

		MessageId uuid.UUID `bun:",pk"`
		Message   *Message  `bun:"rel:belongs-to,join:message_id=id"`

		IsPinned bool `bun:",notnull"`
	}
)
