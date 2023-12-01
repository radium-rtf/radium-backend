package model

import (
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

	GroupAnswers struct {
		Id         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		InviteCode string    `json:"inviteCode"`
		Courses    []*Course `json:"courses"`

		UserAnswers []*UserAnswers `json:"students"`
	}
)

func NewGroup(group *entity.Group) *Group {
	courses := NewCourses(group.Courses, uuid.UUID{})
	students := NewUsers(group.Students)
	return &Group{
		Id:         group.Id,
		InviteCode: group.InviteCode,
		Name:       group.Name,
		Students:   students,
		Courses:    courses,
	}
}

func NewGroups(groups []*entity.Group) []*Group {
	groupsDto := make([]*Group, 0, len(groups))
	for _, group := range groups {
		groupsDto = append(groupsDto, NewGroup(group))
	}
	return groupsDto
}

func NewGroupAnswers(group *entity.Group) GroupAnswers {
	return GroupAnswers{
		Id:          group.Id,
		Name:        group.Name,
		InviteCode:  group.InviteCode,
		Courses:     NewCourses(group.Courses, uuid.UUID{}),
		UserAnswers: NewUserAnswers(group.Students),
	}
}
