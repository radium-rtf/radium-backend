package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Channel struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewChannel(channel *entity.Channel) Channel {
	return Channel{
		Id: channel.Id,
	}
}
