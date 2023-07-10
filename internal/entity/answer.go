package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"math"
)

type (
	Answer struct {
		DBModel
		Verdict   Verdict   `gorm:"not null"`
		UserId    uuid.UUID `gorm:"index:idx_userId_sectionId; type:uuid; not null"`
		SectionId uuid.UUID `gorm:"index:idx_userId_sectionId; type:uuid; not null"`

		Choice      *ChoiceSectionAnswer      `gorm:"polymorphic:Owner"`
		MultiChoice *MultichoiceSectionAnswer `gorm:"polymorphic:Owner"`
		ShortAnswer *ShortAnswerSectionAnswer `gorm:"polymorphic:Owner"`
		Answer      *AnswerSectionAnswer      `gorm:"polymorphic:Owner"`
	}

	AnswerPost struct {
		SectionId   uuid.UUID                     `json:"id"`
		Choice      *ChoiceSectionAnswerPost      `json:"choice,omitempty"`
		MultiChoice *MultichoiceSectionAnswerPost `json:"multiChoice,omitempty"`
		ShortAnswer *ShortAnswerSectionAnswerPost `json:"shortAnswer,omitempty"`
		Answer      *AnswerSectionAnswerPost      `json:"answer,omitempty"`
	}

	MultichoiceSectionAnswerPost struct {
		Answer []string `json:"answer" swaggertype:"array,string"`
	}

	AnswerSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ChoiceSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ShortAnswerSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ChoiceSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID `json:"ownerID" gorm:"type:uuid; not null"`
		OwnerType string    `json:"ownerType" gorm:"not null"`
		Answer    string    `json:"answer" gorm:"not null"`
	}

	MultichoiceSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID      `json:"ownerID" gorm:"type:uuid; not null"`
		OwnerType string         `json:"ownerType" gorm:"not null"`
		Answer    pq.StringArray `json:"answer" swaggertype:"array,string"  gorm:"type:text[]; not null"`
	}

	ShortAnswerSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID `json:"ownerID" gorm:"type:uuid; not null"`
		OwnerType string    `json:"ownerType" gorm:"not null"`
		Answer    string    `json:"answer" gorm:"not null"`
	}

	AnswerSectionAnswer struct {
		DBModel
		OwnerID   uuid.UUID     `json:"ownerID" gorm:"type:uuid; not null"`
		OwnerType string        `json:"ownerType" gorm:"not null"`
		Answer    string        `json:"answer" gorm:"not null"`
		Review    *AnswerReview `json:"review" gorm:"foreignKey:OwnerId"`
	}
)

func (a AnswerSectionAnswer) Score(maxScore uint) uint {
	if a.Review != nil {
		return uint(math.Round(float64(float32(maxScore) * a.Review.Score)))
	}
	return 0
}

func (a Answer) Score(section *Section) uint {
	maxScore := section.MaxScore()
	if a.Answer != nil {
		return a.Answer.Score(maxScore)
	}
	if a.Verdict == VerdictOK {
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

func NewPostToAnswer(post *AnswerPost, userId uuid.UUID) *Answer {
	var (
		choice      *ChoiceSectionAnswer
		multichoice *MultichoiceSectionAnswer
		shortAnswer *ShortAnswerSectionAnswer
		answer      *AnswerSectionAnswer
	)

	if post.Choice != nil {
		choice = &ChoiceSectionAnswer{Answer: post.Choice.Answer}
	}

	if post.MultiChoice != nil {
		multichoice = &MultichoiceSectionAnswer{Answer: post.MultiChoice.Answer}
	}

	if post.ShortAnswer != nil {
		choice = &ChoiceSectionAnswer{Answer: post.ShortAnswer.Answer}
	}

	if post.Answer != nil {
		answer = &AnswerSectionAnswer{Answer: post.Answer.Answer}
	}

	return &Answer{
		Verdict:     VerdictEMPTY,
		SectionId:   post.SectionId,
		UserId:      userId,
		Choice:      choice,
		MultiChoice: multichoice,
		ShortAnswer: shortAnswer,
		Answer:      answer,
	}
}
