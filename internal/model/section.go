package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type (
	Section struct {
		Id        uuid.UUID          `json:"id"`
		PageId    uuid.UUID          `json:"pageId"`
		Order     float64            `json:"order"`
		Type      entity.SectionType `json:"type" enums:"choice,multiChoice,text,shortAnswer,answer,code"`
		Score     uint               `json:"score"`
		Answer    string             `json:"answer"`
		Answers   []string           `json:"answers" swaggertype:"array,string"`
		Content   string             `json:"content"`
		MaxScore  uint               `json:"maxScore"`
		Variants  []string           `json:"variants"`
		Verdict   verdict.Type       `json:"verdict" enums:"OK,WA,WAIT,"`
		Keys      []string           `json:"keys"`
		Review    *Review            `json:"review"`
		Attempts  int                `json:"attempts"`
		FileTypes []string           `json:"fileTypes"`
		File      *File              `json:"file"`
	}
)

func NewSections(sections []*entity.Section, answers map[uuid.UUID][]*entity.Answer) ([]*Section, uint, uint) {
	dtos := make([]*Section, 0, len(sections))
	var sumMaxScore, sumScore uint = 0, 0

	for _, section := range sections {
		var (
			attempts = int(section.MaxAttempts.Int16)
			answer   *entity.Answer
			score    uint
		)

		answers, ok := answers[section.Id]
		if ok && len(answers) >= 1 {
			answer = answers[0]
			attempts = max(int(section.MaxAttempts.Int16)-len(answers), 0)
			score = answer.Score(section)
		}

		sumMaxScore += section.GetMaxScore()
		sumScore += score

		dto := NewSection(section, answer, attempts)
		dtos = append(dtos, dto)
	}

	return dtos, sumScore, sumMaxScore
}

func NewSection(section *entity.Section, answer *entity.Answer, attempts int) *Section {
	var (
		verdict   = verdict.EMPTY
		score     = uint(0)
		answerStr = ""
		answers   []string
		review    *Review
		file      *File
	)

	if answer != nil {
		verdict = answer.Verdict
		score = answer.Score(section)
		answerStr = answer.Answer
		answers = answer.Answers
		review = NewReview(answer.Review)
		file = NewFile(answer.File)
	}

	return &Section{
		Id:        section.Id,
		PageId:    section.PageId,
		Order:     section.Order,
		Content:   section.Content,
		MaxScore:  section.GetMaxScore(),
		Verdict:   verdict,
		Variants:  section.GetVariants(),
		Type:      section.Type,
		Score:     score,
		Answers:   answers,
		Answer:    answerStr,
		Keys:      section.Keys,
		Attempts:  attempts,
		Review:    review,
		FileTypes: section.FileTypes,
		File:      file,
	}
}
