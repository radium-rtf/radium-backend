package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Group struct {
		Id        uuid.UUID     `json:"id"`
		Name      string        `json:"name"`
		AvatarUrl string        `json:"avatar_url"`
		Members   []uuid.UUID   `json:"members"`
		Settings  GroupSettings `json:"settings"`
	}

	GroupSettings struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewGroup(groupChat *entity.Group) Group {
	members := make([]uuid.UUID, 0, len(groupChat.Members))
	for _, member := range groupChat.Members {
		members = append(members, member.Id)
	}
	return Group{
		Id:        groupChat.Id,
		Name:      groupChat.Name,
		AvatarUrl: groupChat.AvatarUrl,
		Members:   members,
		Settings:  GroupSettings{},
	}
}
