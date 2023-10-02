package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Module struct {
		Id       uuid.UUID `json:"id"`
		Slug     string    `json:"slug"`
		Name     string    `json:"name"`
		Order    float64   `json:"order"`
		MaxScore uint      `json:"maxScore"`
		Score    uint      `json:"score"`
		Pages    []*Page   `json:"pages"`
	}
)

func NewModule(module *entity.Module, answers map[uuid.UUID]*entity.Answer) *Module {
	pages, score, maxScore := NewPages(module.Pages, answers)

	return &Module{
		Id:       module.Id,
		Name:     module.Name,
		Slug:     module.Slug,
		Order:    module.Order,
		Score:    score,
		MaxScore: maxScore,
		Pages:    pages,
	}
}

func NewModules(modules []*entity.Module, answers map[uuid.UUID]*entity.Answer) ([]*Module, uint, uint) {
	dtos := make([]*Module, 0, len(modules))
	var score, maxScore uint = 0, 0

	for _, module := range modules {
		dto := NewModule(module, answers)
		score += dto.Score
		maxScore += dto.MaxScore

		dtos = append(dtos, dto)
	}

	return dtos, score, maxScore
}
