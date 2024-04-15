package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"time"
)

type (
	Notification struct {
		Id uuid.UUID `json:"id"`

		Review *ReviewNotification `json:"review"`

		Type      string    `json:"type"`
		Read      bool      `json:"read"`
		CreatedAt time.Time `json:"createdAt"`
	}

	ReviewNotification struct {
		Reviewer   *User     `json:"reviewer"`
		ModuleName string    `json:"moduleName"`
		ModuleId   uuid.UUID `json:"moduleId"`
		CourseId   uuid.UUID `json:"courseId"`
		PageId     uuid.UUID `json:"pageId"`
		PageName   string    `json:"pageName"`
		MaxScore   uint      `json:"maxScore"`
		Score      uint      `json:"score"`
	}
)

func NewNotification(notification entity.Notification) Notification {
	answer := notification.Answer
	section := answer.Section
	page := section.Page
	module := page.Module

	reviewer := answer.Review.Reviewer

	sectionDto := NewSection(section, answer, 0)

	reviewNotification := ReviewNotification{
		Reviewer: NewUser(reviewer),

		ModuleId:   module.Id,
		ModuleName: module.Name,

		CourseId: module.CourseId,

		PageId:   page.Id,
		PageName: page.Name,

		Score:    sectionDto.Score,
		MaxScore: sectionDto.MaxScore,
	}

	return Notification{
		Id:        notification.Id,
		Type:      string(notification.Type),
		Review:    &reviewNotification,
		CreatedAt: notification.CreatedAt,
		Read:      notification.Read,
	}
}

func NewNotifications(notifications []entity.Notification) []Notification {
	dtos := make([]Notification, 0, len(notifications))
	for _, notification := range notifications {
		dtos = append(dtos, NewNotification(notification))
	}
	return dtos
}
