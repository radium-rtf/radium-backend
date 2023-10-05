package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"github.com/uptrace/bun"
	"math"
)

type (
	Answer struct {
		bun.BaseModel `bun:"table:answers"`
		DBModel

		UserId    uuid.UUID
		SectionId uuid.UUID

		Type    SectionType
		Verdict verdict.Type
		Answer  string
		Answers pq.StringArray

		Language string

		Review *Review `bun:"rel:has-one,join:id=answer_id"`
	}

	UsersAnswersCollection struct {
		Users           []*User
		AnswersByUserId map[uuid.UUID]*AnswersCollection
	}

	AnswersCollection struct {
		AnswerBySectionId map[uuid.UUID]*Answer
	}
)

func (a Answer) Score(section *Section) uint {
	maxScore := section.GetMaxScore()

	if (a.Type == AnswerType || a.Type == CodeType) && a.Review != nil {
		return uint(math.Round(float64(maxScore) * a.Review.Score))
	}

	if a.Verdict == verdict.OK {
		return maxScore
	}
	return 0
}

func NewAnswersCollection(answers []*Answer) *AnswersCollection {
	var answerBySectionId = make(map[uuid.UUID]*Answer, len(answers))

	for _, answer := range answers {
		answerBySectionId[answer.SectionId] = answer
	}

	return &AnswersCollection{AnswerBySectionId: answerBySectionId}
}

func NewUsersAnswersCollection(users []*User, answers []*Answer) *UsersAnswersCollection {
	result := &UsersAnswersCollection{
		Users:           users,
		AnswersByUserId: make(map[uuid.UUID]*AnswersCollection, len(users)),
	}

	for _, answer := range answers {
		if _, ok := result.AnswersByUserId[answer.UserId]; !ok {
			result.AnswersByUserId[answer.UserId] = &AnswersCollection{
				AnswerBySectionId: make(map[uuid.UUID]*Answer),
			}
		}
		result.AnswersByUserId[answer.UserId].AnswerBySectionId[answer.SectionId] = answer
	}

	return result
}
