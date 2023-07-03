package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	}

	AnswerPost struct {
		SectionId   uuid.UUID                     `json:"id"`
		Choice      *ChoiceSectionAnswerPost      `json:"choice,omitempty"`
		MultiChoice *MultichoiceSectionAnswerPost `json:"multiChoice,omitempty"`
		ShortAnswer *ShortAnswerSectionAnswerPost `json:"shortAnswer,omitempty"`
	}

	MultichoiceSectionAnswerPost struct {
		Answer []string `json:"answer" swaggertype:"array,string"`
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
)

func NewPostToAnswer(post *AnswerPost, userId uuid.UUID) *Answer {
	var (
		choice      *ChoiceSectionAnswer
		multichoice *MultichoiceSectionAnswer
		shortAnswer *ShortAnswerSectionAnswer
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

	return &Answer{
		Verdict:     VerdictEMPTY,
		SectionId:   post.SectionId,
		UserId:      userId,
		Choice:      choice,
		MultiChoice: multichoice,
		ShortAnswer: shortAnswer,
	}
}
