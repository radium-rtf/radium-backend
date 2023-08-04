package entity

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	Section struct {
		DBModel
		PageId   uuid.UUID `gorm:"type:uuid; not null"`
		Order    uint      `gorm:"not null"`
		MaxScore uint      `gorm:"not null; default:10"`

		TextSection        *TextSection        `gorm:"polymorphic:Owner"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner"`
		AnswerSection      *AnswerSection      `gorm:"polymorphic:Owner"`
		CodeSection        *CodeSection        `gorm:"polymorphic:Owner"`
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
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	ChoiceSection struct {
		DBModel
		Question  string         `gorm:"not null"`
		Answer    string         `gorm:"not null"`
		Variants  pq.StringArray `gorm:"type:text[]; not null"`
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
	}

	MultiChoiceSection struct {
		DBModel
		Question  string         `gorm:"not null"`
		Answer    pq.StringArray `gorm:"type:text[]; not null"`
		Variants  pq.StringArray `gorm:"type:text[]; not null"`
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
	}

	ShortAnswerSection struct {
		DBModel
		Question  string    `gorm:"not null"`
		Answer    string    `gorm:"not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	CodeSection struct {
		DBModel
		Question  string    `gorm:"not null"`
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
	fmt.Printf("%+v", s)
	if s.CodeSection != nil {
		return s.CodeSection.Question
	}
	return ""
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
