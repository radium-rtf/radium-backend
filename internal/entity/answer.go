package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	Answer struct {
		Id        uuid.UUID
		Verdict   Verdict
		UserId    uuid.UUID `gorm:"index:,unique,composite:answer"`
		SectionId uuid.UUID `gorm:"index:,unique,composite:answer"`

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
		Answer pq.StringArray `json:"answer" swaggertype:"array,string"  gorm:"type:text[]"`
	}

	ChoiceSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ShortAnswerSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ChoiceSectionAnswer struct {
		Id        uuid.UUID `gorm:"default:gen_random_uuid()"`
		OwnerID   uuid.UUID `json:"ownerID"`
		OwnerType string    `json:"ownerType" gorm:"default:main_section_test"`
		Answer    string    `json:"answer"`
	}

	MultichoiceSectionAnswer struct {
		Id        uuid.UUID      `gorm:"default:gen_random_uuid()"`
		OwnerID   uuid.UUID      `json:"ownerID"`
		OwnerType string         `json:"ownerType" gorm:"default:main_section_test"`
		Answer    pq.StringArray `json:"answer" swaggertype:"array,string"  gorm:"type:text[]"`
	}

	ShortAnswerSectionAnswer struct {
		Id        uuid.UUID `gorm:"default:gen_random_uuid()"`
		OwnerID   uuid.UUID `json:"ownerID"`
		OwnerType string    `json:"ownerType" gorm:"default:main_section_test"`
		Answer    string    `json:"answer"`
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
		Id:          uuid.New(),
		Verdict:     VerdictEMPTY,
		SectionId:   post.SectionId,
		UserId:      userId,
		Choice:      choice,
		MultiChoice: multichoice,
		ShortAnswer: shortAnswer,
	}
}
