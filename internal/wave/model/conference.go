package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Conference struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewConference(conference *entity.Conference) Conference {
	return Conference{
		Id: conference.Id,
	}
}
