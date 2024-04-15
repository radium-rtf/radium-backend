package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	Link struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Link string    `json:"link"`
	}
)

func NewLinks(links []*entity.Link) []Link {
	dtos := make([]Link, 0, len(links))
	for _, link := range links {
		dtos = append(dtos, NewLink(link))
	}
	return dtos
}

func NewLink(link *entity.Link) Link {
	return Link{
		Id:   link.Id,
		Name: link.Name,
		Link: link.Link,
	}
}
