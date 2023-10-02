package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type (
	Section struct {
		Id       uuid.UUID          `json:"id"`
		PageId   uuid.UUID          `json:"pageId"`
		Order    float64            `json:"order"`
		Type     entity.SectionType `json:"type" enums:"choice,multiChoice,text,shortAnswer,answer,code"`
		Score    uint               `json:"score"`
		Answer   string             `json:"answer"`
		Answers  []string           `json:"answers" swaggertype:"array,string"`
		Content  string             `json:"content"`
		MaxScore uint               `json:"maxScore"`
		Variants []string           `json:"variants"`
		Verdict  verdict.Type       `json:"verdict" enums:"OK,WA,WAIT,"`
	}
)

func NewSections(sections []*entity.Section, answers map[uuid.UUID]*entity.Answer) ([]*Section, uint, uint) {
	dtos := make([]*Section, 0, len(sections))
	var sumMaxScore, sumScore uint = 0, 0

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

		sumMaxScore += section.GetMaxScore()
		sumScore += score

		dto := NewSection(section, verdictType, score, answerStr, answersArr)
		dtos = append(dtos, dto)
	}

	return dtos, sumScore, sumMaxScore
}

func NewSection(section *entity.Section, verdict verdict.Type,
	score uint, answer string, answers []string) *Section {
	return &Section{
		Id:       section.Id,
		PageId:   section.PageId,
		Order:    section.Order,
		Content:  section.Content,
		MaxScore: section.GetMaxScore(),
		Verdict:  verdict,
		Variants: section.GetVariants(),
		Type:     section.Type,
		Score:    score,
		Answers:  answers,
		Answer:   answer,
	}
}
