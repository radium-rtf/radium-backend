package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"slices"
	"time"
)

type (
	Answer struct {
		Id      uuid.UUID `json:"id"`
		Section *Section  `json:"section"`

		Type    entity.SectionType `json:"type"`
		Verdict verdict.Type       `json:"verdict"`

		Answer  string   `json:"answer"`
		Answers []string `json:"answers"`

		Language string `json:"language"`

		Review *Review `json:"review"`

		CreatedAt time.Time `json:"createdAt"`
	}

	UserAnswers struct {
		User          User `json:"user"`
		WithoutReview uint `json:"withoutReview"`

		MaxScore uint `json:"maxScore"`
		Score    uint `json:"score"`

		Answers []Answer `json:"answers"`
	}
)

func NewUserAnswers(students []*entity.User) []*UserAnswers {
	var userAnswers = make([]*UserAnswers, 0, len(students))

	for _, student := range students {

		withoutReview := slices.IndexFunc(student.Answers, func(answer *entity.Answer) bool {
			return answer.Review != nil
		})
		if withoutReview == -1 {
			withoutReview = len(student.Answers)
		}

		reviewed := make(map[uuid.UUID]*entity.Answer, len(student.Answers))
		for i := withoutReview; i < len(student.Answers); i++ {
			answer := student.Answers[i]
			if _, ok := reviewed[answer.SectionId]; ok {
				continue
			}
			reviewed[answer.SectionId] = answer
		}

		studentAnswers := make([]*entity.Answer, 0, len(student.Answers))
		setAnswers := make(map[uuid.UUID]bool, len(student.Answers))
		setReviewedAnswers := make(map[uuid.UUID]bool, len(student.Answers))
		withoutReview = 0
		for _, answer := range student.Answers {
			_, isReviewed := reviewed[answer.SectionId]

			if !setReviewedAnswers[answer.SectionId] && answer.Review != nil {
				setReviewedAnswers[answer.SectionId] = true
				studentAnswers = append(studentAnswers, answer)
				continue
			}

			if isReviewed && answer.CreatedAt.Before(reviewed[answer.SectionId].CreatedAt) {
				continue
			}

			if setAnswers[answer.SectionId] {
				continue
			}

			if answer.Review == nil {
				withoutReview += 1
			}
			setAnswers[answer.SectionId] = true
			studentAnswers = append(studentAnswers, answer)
		}

		userAnswer := &UserAnswers{
			User:          NewUser(student),
			Answers:       NewAnswers(studentAnswers),
			WithoutReview: uint(withoutReview),
		}

		userAnswers = append(userAnswers, userAnswer)
	}

	return userAnswers
}

func NewAnswers(answers []*entity.Answer) []Answer {
	var dtos = make([]Answer, 0, len(answers))

	for _, answer := range answers {
		dtos = append(dtos, NewAnswer(answer))
	}
	return dtos
}

func NewAnswer(answer *entity.Answer) Answer {
	return Answer{
		Id:      answer.Id,
		Section: NewSection(answer.Section, answer.Verdict, answer.Score(answer.Section), answer.Section.Answer, answer.Section.Answers, 0),

		Type:    answer.Type,
		Verdict: answer.Verdict,

		Answers: answer.Answers,
		Answer:  answer.Answer,

		Language: answer.Language,

		Review: NewReview(answer.Review),

		CreatedAt: answer.CreatedAt,
	}
}