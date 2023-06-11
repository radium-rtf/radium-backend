package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const (
	VerdictOK = Verdict("OK")
	VerdictWA = Verdict("WA")
)

type (
	Verdict string

	SectionPost struct {
		PageId uuid.UUID
		Order  uint

		TextSection        *TextSection        `gorm:"polymorphic:Owner" json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner" json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner" json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner" json:"shortanswer,omitempty"`
	}
	Section struct {
		ID     uuid.UUID `gorm:"default:gen_random_uuid()"`
		PageId uuid.UUID
		Order  uint

		TextSection        *TextSection        `gorm:"polymorphic:Owner" json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `gorm:"polymorphic:Owner" json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `gorm:"polymorphic:Owner" json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `gorm:"polymorphic:Owner" json:"shortanswer,omitempty"`
	}

	SectionDto struct {
		ID                 uuid.UUID              `json:"id"`
		PageId             uuid.UUID              `json:"pageId"`
		Order              uint                   `json:"order"`
		TextSection        *TextSectionDto        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionDto      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionDto `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionDto `json:"shortanswer,omitempty"`
	}

	TextSection struct {
		ID        uuid.UUID `gorm:"default:gen_random_uuid()"`
		Content   string
		OwnerID   uuid.UUID
		OwnerType string `gorm:"default:main_section_test"`
	}

	TextSectionDto struct {
		Content string `json:"content"`
	}

	ChoiceSection struct {
		ID        uuid.UUID `gorm:"default:gen_random_uuid()"`
		MaxScore  uint
		Question  string
		Answer    string
		Variants  pq.StringArray `gorm:"type:text[]" swaggertype:"array,string"`
		OwnerID   uuid.UUID
		OwnerType string `gorm:"default:main_section_test"`
	}

	ChoiceSectionDto struct {
		MaxScore uint     `json:"maxScore"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
		Score    uint     `json:"score"`
		Verdict  Verdict  `json:"verdict"`
	}

	MultiChoiceSection struct {
		ID        uuid.UUID `gorm:"default:gen_random_uuid()"`
		MaxScore  uint
		Question  string
		Answer    pq.StringArray `gorm:"type:text[]" swaggertype:"array,string"`
		Variants  pq.StringArray `gorm:"type:text[]" swaggertype:"array,string"`
		OwnerID   uuid.UUID
		OwnerType string `gorm:"default:main_section_test"`
	}

	MultiChoiceSectionDto struct {
		MaxScore uint     `json:"maxScore"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
		Score    uint     `json:"score"`
		Verdict  Verdict  `json:"verdict"`
	}

	ShortAnswerSection struct {
		ID        uuid.UUID `gorm:"default:gen_random_uuid()"`
		MaxScore  uint
		Question  string
		Answer    string
		OwnerID   uuid.UUID
		OwnerType string `gorm:"default:main_section_test"`
	}

	ShortAnswerSectionDto struct {
		MaxScore uint    `json:"maxScore"`
		Question string  `json:"question"`
		Score    uint    `json:"score"`
		Verdict  Verdict `json:"verdict"`
	}
)

func NewSectionPostToSection(post SectionPost) Section {
	if post.ChoiceSection != nil {
		variants := make([]string, 0, len(post.ChoiceSection.Variants))
		variants = append(variants, post.ChoiceSection.Variants...)
		post.ChoiceSection.Variants = variants
	}
	return Section{
		PageId:             post.PageId,
		Order:              post.Order,
		TextSection:        post.TextSection,
		ChoiceSection:      post.ChoiceSection,
		MultiChoiceSection: post.MultiChoiceSection,
		ShortAnswerSection: post.ShortAnswerSection,
	}
}
