package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type State string

var (
	None     State = "none"
	Assigned State = "assigned"
	Editor   State = "editor"
)

type Card struct {
	Id              uuid.UUID  `json:"id"`
	Slug            string     `json:"slug"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	LogoUrl         string     `json:"logoUrl"`
	State           State      `json:"state"`
	Score           uint       `json:"score "`
	MaxScore        uint       `json:"maxScore"`
	ModuleCount     int        `json:"moduleCount"`
	LastVisitedPage *uuid.UUID `json:"lastVisitedPage"`
}

type Main struct {
	TopCard       *Card   `json:"topCard"`
	AuthorCards   []*Card `json:"authorCards"`
	AssignedCards []*Card `json:"assignedCards"`
	Recommended   []*Card `json:"recommended"`
}

func NewCard(course *entity.Course, state State, userId uuid.UUID) *Card {
	if course == nil {
		return nil
	}
	dto := NewCourse(course, map[uuid.UUID][]*entity.Answer{}, userId)

	var lastVisitedPage *uuid.UUID
	if len(course.LastVisitedPage) != 0 {
		lastVisitedPage = &course.LastVisitedPage[0].PageId
	}

	return &Card{
		Id:              dto.Id,
		Slug:            dto.Slug,
		Name:            dto.Name,
		Description:     dto.ShortDescription,
		LogoUrl:         dto.Logo,
		State:           state,
		MaxScore:        dto.MaxScore,
		Score:           dto.Score,
		ModuleCount:     len(dto.Modules),
		LastVisitedPage: lastVisitedPage,
	}
}

func NewCards(courses []*entity.Course, state State, userId uuid.UUID) []*Card {
	cards := make([]*Card, 0, len(courses))
	for _, course := range courses {
		cards = append(cards, NewCard(course, state, userId))
	}
	return cards
}

func NewMainCard(top *entity.LastVisitedPage, assigned, authorship, recommended []*entity.Course, userId uuid.UUID) Main {
	state := Assigned

	canEdit := false
	if top != nil && top.Course != nil {
		canEdit = top.Course.CanEdit(userId)
	}
	if canEdit {
		state = Editor
	}

	var topCard *Card
	if top != nil && top.Course != nil {
		topCard = NewCard(top.Course, state, userId)
		topCard.LastVisitedPage = &top.PageId
	}

	return Main{
		TopCard:       topCard,
		AssignedCards: NewCards(assigned, Assigned, userId),
		AuthorCards:   NewCards(authorship, Editor, userId),
		Recommended:   NewCards(recommended, None, userId),
	}
}
