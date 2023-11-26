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
		Next     *uuid.UUID `json:"next"`
		Previous *uuid.UUID `json:"previous"`
	}

	NextAndPreviousPage struct {
		Next     *uuid.UUID `json:"next"`
		Previous *uuid.UUID `json:"previous"`
	}
)

func NewPage(page *entity.Page, answers map[uuid.UUID][]*entity.Answer, nextAndPrevious *NextAndPreviousPage) *Page {
	sectionsDto, score, maxScore := NewSections(page.Sections, answers)

	var next, previous *uuid.UUID
	if nextAndPrevious != nil {
		previous = nextAndPrevious.Previous
		next = nextAndPrevious.Next
	}

	return &Page{
		Id:       page.Id,
		Slug:     page.Slug,
		Name:     page.Name,
		Order:    page.Order,
		Score:    score,
		MaxScore: maxScore,
		Sections: sectionsDto,
		Next:     next,
		Previous: previous,
	}
}

func NewPages(pages []*entity.Page, answers map[uuid.UUID][]*entity.Answer) ([]*Page, uint, uint) {
	dtos := make([]*Page, 0, len(pages))
	var maxScore, score uint = 0, 0

	for _, page := range pages {
		dto := NewPage(page, answers, nil)
		maxScore += dto.MaxScore
		score += dto.Score

		dtos = append(dtos, dto)
	}

	return dtos, score, maxScore
}

func GetNextAndPreviousPage(moduleIndex, pageIndex int, modules []*entity.Module) *NextAndPreviousPage {
	var prev *uuid.UUID
	for module := moduleIndex; module >= 0; module-- {
		page := len(modules[module].Pages) - 1
		if module == moduleIndex {
			page = pageIndex - 1
		}
		if page >= 0 {
			prev = &modules[module].Pages[page].Id
			break
		}
	}

	var next *uuid.UUID
	for module := moduleIndex; module < len(modules); module++ {
		page := 0
		if module == moduleIndex {
			page = pageIndex + 1
		}
		if page < len(modules[module].Pages) {
			next = &modules[module].Pages[page].Id
			break
		}
	}
	return &NextAndPreviousPage{Next: next, Previous: prev}
}
