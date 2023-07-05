package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	ChoiceType      = SectionType("choice")
	MultiChoiceType = SectionType("multiChoice")
	TextType        = SectionType("text")
	ShortAnswerType = SectionType("shortAnswer")
)

type (
	SectionType string

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
		Id       uuid.UUID   `json:"id"`
		PageId   uuid.UUID   `json:"pageId"`
		Order    uint        `json:"order"`
		Type     SectionType `json:"type"`
		Content  string      `json:"content"`
		MaxScore uint        `json:"maxScore"`
		Variants []string    `json:"variants"`
		Verdict  Verdict     `json:"verdict"`
	}

	TextSection struct {
		DBModel
		Content   string    `gorm:"type:text; not null"`
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
