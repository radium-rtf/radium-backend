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
)
