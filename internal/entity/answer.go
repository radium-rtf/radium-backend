package entity

import (
	"math"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type (
	Answer struct {
		DBModel
		Verdict   verdict.Type `gorm:"not null"`
		UserId    uuid.UUID    `gorm:"index:idx_userId_sectionId; type:uuid; not null"`
		SectionId uuid.UUID    `gorm:"index:idx_userId_sectionId; type:uuid; not null"`

		Choice      *ChoiceSectionAnswer      `gorm:"polymorphic:Owner"`
		MultiChoice *MultichoiceSectionAnswer `gorm:"polymorphic:Owner"`
		ShortAnswer *ShortAnswerSectionAnswer `gorm:"polymorphic:Owner"`
		Answer      *AnswerSectionAnswer      `gorm:"polymorphic:Owner"`
		Code        *CodeSectionAnswer        `gorm:"polymorphic:Owner"`
	}

	ChoiceSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
		Answer    string    `gorm:"not null"`
	}

	MultichoiceSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
		Answer    pq.StringArray `gorm:"type:text[]; not null"`
	}

	ShortAnswerSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
		Answer    string    `gorm:"not null"`
	}

	AnswerSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID     `gorm:"type:uuid; not null"`
		OwnerType string        `gorm:"not null"`
		Answer    string        `gorm:"not null"`
		Review    *AnswerReview `gorm:"foreignKey:OwnerId"`
	}

	CodeSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID   `gorm:"type:uuid; not null"`
		OwnerType string      `gorm:"not null"`
		Answer    string      `gorm:"not null"`
		Language  string      `gorm:"not null"`
		Review    *CodeReview `gorm:"foreignKey:OwnerId"`
	}
)

func (a AnswerSectionAnswer) Score(maxScore uint) uint {
	if a.Review != nil {
		return uint(math.Round(float64(float32(maxScore) * a.Review.Score)))
	}
	return 0
}

func (a Answer) Score(section *Section) uint {
	maxScore := section.MaxScore
	if a.Answer != nil {
		return a.Answer.Score(maxScore)
	}
	if a.Verdict == verdict.OK {
		return maxScore
	}
	return 0
}

func (a Answer) Answers() []string {
	if a.MultiChoice != nil {
		return a.MultiChoice.Answer
	}
	return []string{}
}

func (a Answer) AnswerStr() string {
	if a.ShortAnswer != nil {
		return a.ShortAnswer.Answer
	}
	if a.Choice != nil {
		return a.Choice.Answer
	}
	if a.Answer != nil {
		return a.Answer.Answer
	}
	return ""
}
