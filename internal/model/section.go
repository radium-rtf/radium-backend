package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

const (
	ChoiceType      = SectionType("choice")
	MultiChoiceType = SectionType("multiChoice")
	TextType        = SectionType("text")
	ShortAnswerType = SectionType("shortAnswer")
	AnswerType      = SectionType("answer")
)

type (
	SectionType string

	Section struct {
		Id       uuid.UUID    `json:"id"`
		PageId   uuid.UUID    `json:"pageId"`
		Order    uint         `json:"order"`
		Type     SectionType  `json:"type"`
		Score    uint         `json:"score"`
		Answer   string       `json:"answer"`
		Answers  []string     `json:"answers" swaggertype:"array,string"`
		Content  string       `json:"content"`
		MaxScore uint         `json:"maxScore"`
		Variants []string     `json:"variants"`
		Verdict  verdict.Type `json:"verdict" swaggertype:"string"`
	}
)

func NewSections(sections []*entity.Section, answers map[uuid.UUID]*entity.Answer) []*Section {
	dtos := make([]*Section, 0, len(sections))
	for _, section := range sections {
		var (
			verdictType = verdict.EMPTY
			score       = uint(0)
			answerStr   = ""
			answersArr  []string
		)

		answer, ok := answers[section.Id]
		if ok {
			verdictType = answer.Verdict
			score = answer.Score(section)
			answerStr = answer.AnswerStr()
			answersArr = answer.Answers()
		}

		dto := NewSection(section, verdictType, score, answerStr, answersArr)
		dtos = append(dtos, dto)
	}
	return dtos
}

func NewSection(section *entity.Section, verdict verdict.Type,
	score uint, answer string, answers []string) *Section {
	var sectionType SectionType
	if section.MultiChoiceSection != nil {
		sectionType = MultiChoiceType
	} else if section.ChoiceSection != nil {
		sectionType = ChoiceType
	} else if section.TextSection != nil {
		sectionType = TextType
	} else if section.ShortAnswerSection != nil {
		sectionType = ShortAnswerType
	} else if section.AnswerSection != nil {
		sectionType = AnswerType
	}
	return &Section{
		Id:       section.Id,
		PageId:   section.PageId,
		Order:    section.Order,
		Content:  section.Content(),
		MaxScore: section.MaxScore(),
		Verdict:  verdict,
		Variants: section.Variants(),
		Type:     sectionType,
		Score:    score,
		Answers:  answers,
		Answer:   answer,
	}
}
