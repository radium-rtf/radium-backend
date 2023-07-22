package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	Section struct {
		DBModel
		PageId uuid.UUID `gorm:"type:uuid; not null"`
		Order  uint      `gorm:"not null"`

		TextSection        *TextSection        `gorm:"polymorphic:Owner" json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner" json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner" json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner" json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSection      `gorm:"polymorphic:Owner" json:"answer,omitempty"`
	}

	TextSection struct {
		DBModel
		Content   string    `gorm:"type:text; not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	AnswerSection struct {
		DBModel
		Question  string    `gorm:"type:text; not null"`
		MaxScore  uint      `gorm:"not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	ChoiceSection struct {
		DBModel
		MaxScore  uint           `gorm:"not null"`
		Question  string         `gorm:"not null"`
		Answer    string         `gorm:"not null"`
		Variants  pq.StringArray `gorm:"type:text[]; not null"`
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
	}

	MultiChoiceSection struct {
		DBModel
		MaxScore  uint           `gorm:"not null"`
		Question  string         `gorm:"not null"`
		Answer    pq.StringArray `gorm:"type:text[]; not null" swaggertype:"array,string"`
		Variants  pq.StringArray `gorm:"type:text[]; not null" swaggertype:"array,string"`
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
	}

	ShortAnswerSection struct {
		DBModel
		MaxScore  uint      `gorm:"not null"`
		Question  string    `gorm:"not null"`
		Answer    string    `gorm:"not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}
)

func (s Section) Content() string {
	if s.ChoiceSection != nil {
		return s.ChoiceSection.Question
	}
	if s.MultiChoiceSection != nil {
		return s.MultiChoiceSection.Question
	}
	if s.ShortAnswerSection != nil {
		return s.ShortAnswerSection.Question
	}
	if s.TextSection != nil {
		return s.TextSection.Content
	}
	if s.AnswerSection != nil {
		return s.AnswerSection.Question
	}
	return ""
}

func (s Section) MaxScore() uint {
	if s.ChoiceSection != nil { // TODO: ПЕРЕНЕСТИ В SECTION
		return s.ChoiceSection.MaxScore
	}
	if s.MultiChoiceSection != nil {
		return s.MultiChoiceSection.MaxScore
	}
	if s.ShortAnswerSection != nil {
		return s.ShortAnswerSection.MaxScore
	}
	if s.TextSection != nil {
		return 0
	}
	if s.AnswerSection != nil {
		return s.AnswerSection.MaxScore
	}
	return 0
}

func (s Section) Variants() []string {
	if s.ChoiceSection != nil {
		return s.ChoiceSection.Variants
	}
	if s.MultiChoiceSection != nil {
		return s.MultiChoiceSection.Variants
	}
	return []string{}
}
