package model

import (
	"github.com/dranikpg/dto-mapper"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Group struct {
		Id         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		InviteCode string    `json:"inviteCode"`
		Courses    []*Course `json:"courses"`
		Students   []*User   `json:"students"`
	}
)

func NewGroup(group *entity.Group) *Group {
	groupDto := Group{}
	dto.Map(&groupDto, group)
	return &groupDto
}

func NewGroups(groups []*entity.Group) []*Group {
	groupsDto := make([]*Group, 0, len(groups))
	for _, group := range groups {
		groupsDto = append(groupsDto, NewGroup(group))
	}
	return groupsDto
}
