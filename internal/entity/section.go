package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	SectionPost struct {
		PageId uuid.UUID `gorm:"type:uuid; not null"`
		Order  uint      `gorm:"not null"`

		TextSection        *TextSectionPost        `gorm:"polymorphic:Owner" json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionPost      `gorm:"polymorphic:Owner" json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionPost `gorm:"polymorphic:Owner" json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionPost `gorm:"polymorphic:Owner" json:"shortanswer,omitempty"`
	}

	TextSectionPost struct {
		Content string
	}

	ChoiceSectionPost struct {
		MaxScore uint
		Question string
		Answer   string
		Variants []string `swaggertype:"array,string"`
	}

	MultiChoiceSectionPost struct {
		MaxScore uint
		Question string
		Answer   []string `swaggertype:"array,string"`
		Variants []string `swaggertype:"array,string"`
	}

	ShortAnswerSectionPost struct {
		MaxScore uint
		Question string
		Answer   string
	}

	Section struct {
		DBModel
		PageId uuid.UUID `gorm:"type:uuid; not null"`
		Order  uint      `gorm:"not null"`

		TextSection        *TextSection        `gorm:"polymorphic:Owner" json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner" json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner" json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner" json:"shortanswer,omitempty"`
	}

	SectionDto struct {
		Id                 uuid.UUID              `json:"id"`
		PageId             uuid.UUID              `json:"pageId"`
		Order              uint                   `json:"order"`
		TextSection        *TextSectionDto        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionDto      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionDto `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionDto `json:"shortanswer,omitempty"`
	}

	TextSection struct {
		DBModel
		Content   string    `gorm:"type:text; not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	TextSectionDto struct {
		Content string `json:"content"`
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

	ChoiceSectionDto struct {
		MaxScore uint     `json:"maxScore"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
		Verdict  Verdict  `json:"verdict"`
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

	MultiChoiceSectionDto struct {
		MaxScore uint     `json:"maxScore"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
		Score    uint     `json:"score"`
		Verdict  Verdict  `json:"verdict"`
	}

	ShortAnswerSection struct {
		DBModel
		MaxScore  uint      `gorm:"not null"`
		Question  string    `gorm:"not null"`
		Answer    string    `gorm:"not null"`
		OwnerID   uuid.UUID `gorm:"type:uuid; not null"`
		OwnerType string    `gorm:"not null"`
	}

	ShortAnswerSectionDto struct {
		MaxScore uint    `json:"maxScore"`
		Question string  `json:"question"`
		Verdict  Verdict `json:"verdict"`
	}
)
