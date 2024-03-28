package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Content struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewContent(content *entity.Content) Content {
	return Content{
		Id: content.Id,
	}
}
