package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Dialogue struct {
		Id           uuid.UUID         `json:"id"`
		FirstUserId  uuid.UUID         `json:"firstUserId"`
		SecondUserId uuid.UUID         `json:"secondUserId"`
		Settings     *DialogueSettings `json:"settings"`
	}

	DialogueSettings struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewDialogue(dialogue *entity.Dialogue) Dialogue {
	if dialogue == nil {
		return Dialogue{}
	}

	return Dialogue{
		Id:           dialogue.Id,
		FirstUserId:  dialogue.FirstUserId,
		SecondUserId: dialogue.SecondUserId,
		Settings:     nil,
	}
}
