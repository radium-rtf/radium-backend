package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Dialogue struct {
		bun.BaseModel `bun:"table:dialogues"`
		DBModel

		FirstUserId    uuid.UUID
		SecondUserId   uuid.UUID
		Messages       []*Message        `bun:"m2m:dialogue_messages,join:Dialogue=Message"`
		Settings       *DialogueSettings `bun:"rel:has-one"`
		PinnedMessages []*Message        `bun:"m2m:dialogue_pinned_messages,join:Dialogue=Message"`
	}

	DialogueSettings struct {
		bun.BaseModel `bun:"table:dialogue_settings"`
		DBModel
	}

	DialogueMessage struct {
		bun.BaseModel `bun:"table:dialogue_messages"`

		DialogueId uuid.UUID `bun:",pk"`
		Dialogue   *Dialogue `bun:"rel:belongs-to,join:dialogue_id=id"`
		MessageId  uuid.UUID `bun:",pk"`
		Message    *Message  `bun:"rel:belongs-to,join:message_id=id"`
	}

	DialoguePinnedMessage struct {
		bun.BaseModel `bun:"table:dialogue_pinned_messages"`

		DialogueId uuid.UUID `bun:",pk"`
		Dialogue   *Dialogue `bun:"rel:belongs-to,join:dialogue_id=id"`
		MessageId  uuid.UUID `bun:",pk"`
		Message    *Message  `bun:"rel:belongs-to,join:message_id=id"`
	}
)
