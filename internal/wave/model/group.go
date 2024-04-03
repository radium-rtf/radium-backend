package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Group struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewGroup(groupChat *entity.Group) Group {
	return Group{
		Id: groupChat.Id,
	}
}
