package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Page struct {
		Id       uuid.UUID  `json:"id"`
		Slug     string     `json:"slug"`
		Name     string     `json:"name"`
		Order    float64    `json:"order"`
		Sections []*Section `json:"sections"`
	}
)

func NewPage(page *entity.Page, answers map[uuid.UUID]*entity.Answer) *Page {
	sectionsDto := NewSections(page.Sections, answers)
	return &Page{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Order:    page.Order,
		Sections: sectionsDto,
	}
}

func NewPages(pages []*entity.Page, answers map[uuid.UUID]*entity.Answer) []*Page {
	dtos := make([]*Page, 0, len(pages))
	for _, page := range pages {
		dtos = append(dtos, NewPage(page, answers))
	}
	return dtos
}
