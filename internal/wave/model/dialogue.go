package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Dialogue struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewDialogue(dialogue *entity.Dialogue) Dialogue {
	return Dialogue{
		Id: dialogue.Id,
	}
}
