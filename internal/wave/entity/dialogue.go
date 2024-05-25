package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Dialogue struct {
		bun.BaseModel `bun:"table:wave.dialogues"`
		Id            uuid.UUID `bun:",unique"`
		FirstUserId   uuid.UUID `bun:",pk"`
		SecondUserId  uuid.UUID `bun:",pk"`
		// Messages       []*Message        `bun:"m2m:wave.dialogue_message,join:Dialogue=Message"`
		// Settings *DialogueSettings `bun:"rel:has-one"`
		// PinnedMessages []*Message        `bun:"m2m:wave.dialogue_pinned,join:Dialogue=Message"`
	}

	DialogueSettings struct {
		bun.BaseModel `bun:"table:wave.dialogue_settings"`
		DBModel
	}

	DialogueMessage struct {
		bun.BaseModel `bun:"table:wave.dialogue_message"`

		DialogueId uuid.UUID `bun:",pk"`
		Dialogue   *Dialogue `bun:"rel:belongs-to,join:dialogue_id=id"`

		MessageId uuid.UUID `bun:",pk"`
		Message   *Message  `bun:"rel:belongs-to,join:message_id=id"`
	}

	DialoguePinnedMessage struct {
		bun.BaseModel `bun:"table:wave.dialogue_pinned"`

		DialogueId uuid.UUID `bun:",pk"`
		Dialogue   *Dialogue `bun:"rel:belongs-to,join:dialogue_id=id"`

		MessageId uuid.UUID `bun:",pk"`
		Message   *Message  `bun:"rel:belongs-to,join:message_id=id"`
	}
)
