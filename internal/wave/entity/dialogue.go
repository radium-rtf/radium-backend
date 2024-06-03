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
		// Settings *DialogueSettings `bun:"rel:has-one"`
	}

	DialogueSettings struct {
		bun.BaseModel `bun:"table:wave.dialogue_settings"`
		DBModel
	}
)
