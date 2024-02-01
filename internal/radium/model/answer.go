package model

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/answer/verdict"
	"sort"
	"time"
)

type (
	Answer struct {
		Id      uuid.UUID `json:"id"`
		Section *Section  `json:"section"`

		Type    entity2.SectionType `json:"type"`
		Verdict verdict.Type        `json:"verdict"`

		Answer  string   `json:"answer"`
		Answers []string `json:"answers"`
		File    *File    `json:"file"`

		Language string `json:"language"`

		Review *Review `json:"review"`

		CreatedAt time.Time `json:"createdAt"`
	}

	UserAnswers struct {
		User          *User `json:"user"`
		WithoutReview uint  `json:"withoutReview"`

		MaxScore uint `json:"maxScore"`
		Score    uint `json:"score"`

		Answers []Answer `json:"answers"`
	}
)

func NewUserAnswers(students []*entity2.User) []*UserAnswers {
	// TODO: фильтровать лучше в базе
	var userAnswers = make([]*UserAnswers, 0, len(students))

	for _, student := range students {

		withoutReview := 0
		set := make(map[uuid.UUID]bool)
		for _, answer := range student.Answers {
			if answer.Review != nil {
				continue
			}
			if !set[answer.SectionId] {
				withoutReview += 1
				set[answer.SectionId] = true
			}
		}

		reviewed := make(map[uuid.UUID]*entity2.Answer, len(student.Answers))
		for i := withoutReview; i < len(student.Answers); i++ {
			answer := student.Answers[i]
			if _, ok := reviewed[answer.SectionId]; ok {
				continue
			}
			reviewed[answer.SectionId] = answer
		}

		studentAnswers := make([]*entity2.Answer, 0, len(student.Answers))
		answersSet := make(map[uuid.UUID]bool, len(student.Answers))
		reviewedAnswersSet := make(map[uuid.UUID]bool, len(student.Answers))
		withoutReview = 0
		for _, answer := range student.Answers {
			_, isReviewed := reviewed[answer.SectionId]

			if !reviewedAnswersSet[answer.SectionId] && answer.Review != nil {
				reviewedAnswersSet[answer.SectionId] = true
				studentAnswers = append(studentAnswers, answer)
				continue
			}

			if isReviewed && answer.CreatedAt.Before(reviewed[answer.SectionId].CreatedAt) {
				continue
			}

			if answersSet[answer.SectionId] {
				continue
			}

			if answer.Review == nil {
				withoutReview += 1
			}
			answersSet[answer.SectionId] = true
			studentAnswers = append(studentAnswers, answer)
		}

		userAnswer := &UserAnswers{
			User:          NewUser(student),
			Answers:       NewAnswers(studentAnswers),
			WithoutReview: uint(withoutReview),
		}

		userAnswers = append(userAnswers, userAnswer)
	}

	sort.Slice(userAnswers, func(i, j int) bool {
		return userAnswers[i].WithoutReview > userAnswers[j].WithoutReview
	})
	return userAnswers
}

func NewAnswers(answers []*entity2.Answer) []Answer {
	var dtos = make([]Answer, 0, len(answers))

	for _, answer := range answers {
		// todo: ПРИ УДАЛЕНИИ СЕКЦИИ УДАЛЯТЬ ОТВЕТЫ убрать if
		if answer.Section == nil {
			continue
		}
		attempts := max(int(answer.Section.MaxAttempts.Int16)-len(answers), 0)
		dtos = append(dtos, NewAnswer(answer, attempts))
	}
	return dtos
}

func NewAnswer(answer *entity2.Answer, attempts int) Answer {
	return Answer{
		Id:      answer.Id,
		Section: NewSection(answer.Section, answer, attempts),

		Type:    answer.Type,
		Verdict: answer.Verdict,

		Answers: answer.Answers,
		Answer:  answer.Answer,

		Language: answer.Language,

		Review: NewReview(answer.Review),
		File:   NewFile(answer.File),

		CreatedAt: answer.CreatedAt,
	}
}
