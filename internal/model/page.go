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
		Score    uint       `json:"score"`
		MaxScore uint       `json:"maxScore"`
		Sections []*Section `json:"sections"`
	}
)

func NewPage(page *entity.Page, answers map[uuid.UUID]*entity.Answer) *Page {
	sectionsDto, score, maxScore := NewSections(page.Sections, answers)

	return &Page{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Order:    page.Order,
		Score:    score,
		MaxScore: maxScore,
		Sections: sectionsDto,
	}
}

func NewPages(pages []*entity.Page, answers map[uuid.UUID]*entity.Answer) ([]*Page, uint, uint) {
	dtos := make([]*Page, 0, len(pages))
	var maxScore, score uint = 0, 0

	for _, page := range pages {
		dto := NewPage(page, answers)
		maxScore += dto.MaxScore
		score += dto.Score

		dtos = append(dtos, dto)
	}

	return dtos, score, maxScore
}
