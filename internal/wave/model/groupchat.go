package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	GroupChat struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewGroupChat(groupChat *entity.GroupChat) GroupChat {
	return GroupChat{
		Id: groupChat.Id,
	}
}
