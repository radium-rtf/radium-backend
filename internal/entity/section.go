package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"math/rand"
)

type (
	Section struct {
		DBModel
		PageId   uuid.UUID `gorm:"type:uuid; not null"`
		Order    uint      `gorm:"not null"`
		MaxScore uint      `gorm:"not null; default:0"`

		TextSection        *TextSection        `gorm:"polymorphic:Owner"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner"`
		AnswerSection      *AnswerSection      `gorm:"polymorphic:Owner"`
		CodeSection        *CodeSection        `gorm:"polymorphic:Owner"`
		PermutationSection *PermutationSection `gorm:"polymorphic:Owner"`
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

	PermutationSection struct {
		DBModel
		Question  string         `gorm:"not null"`
		Answer    pq.StringArray `gorm:"type:text[]; not null"`
		OwnerID   uuid.UUID      `gorm:"type:uuid; not null"`
		OwnerType string         `gorm:"not null"`
	}
)

func (s Section) Content() string {
	switch {
	case s.ChoiceSection != nil:
		return s.ChoiceSection.Question
	case s.ShortAnswerSection != nil:
		return s.ShortAnswerSection.Question
	case s.TextSection != nil:
		return s.TextSection.Content
	case s.AnswerSection != nil:
		return s.AnswerSection.Question
	case s.CodeSection != nil:
		return s.CodeSection.Question
	case s.PermutationSection != nil:
		return s.PermutationSection.Question
	case s.MultiChoiceSection != nil:
		return s.MultiChoiceSection.Question
	default:
		panic("")
	}
}

func (s Section) Variants() []string {
	if s.ChoiceSection != nil {
		return s.ChoiceSection.Variants
	}
	if s.MultiChoiceSection != nil {
		return s.MultiChoiceSection.Variants
	}
	if s.PermutationSection != nil {
		variants := []string(s.PermutationSection.Answer)
		rand.Shuffle(len(variants), func(i, j int) {
			variants[i], variants[j] = variants[j], variants[i]
		})
		return variants
	}
	return []string{}
}

func (s Section) GetMaxScore() uint {
	if s.TextSection != nil {
		return 0
	}
	return s.MaxScore
}
