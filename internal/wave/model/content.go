package model

import (
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Content struct {
		Text string `json:"text"`
	}
)

func NewContent(content *entity.Content) Content {
	if content == nil {
		return Content{}
	}

	return Content{
		Text: content.Text,
	}
}
