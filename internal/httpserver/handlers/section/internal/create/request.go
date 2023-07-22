package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Request struct {
		PageId uuid.UUID `gorm:"type:uuid; not null"`
		Order  uint      `gorm:"not null"`

		TextSection        *TextSectionPost        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionPost      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionPost `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionPost `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSectionPost      `json:"answer,omitempty"`
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

	AnswerSectionPost struct {
		MaxScore uint
		Question string
	}
)

func (r Request) PostToSection() *entity.Section {
	return &entity.Section{
		PageId:             r.PageId,
		Order:              r.Order,
		TextSection:        r.postToText(r.TextSection),
		ChoiceSection:      r.postToChoice(r.ChoiceSection),
		MultiChoiceSection: r.postToMultiChoice(r.MultiChoiceSection),
		ShortAnswerSection: r.postToShortAnswer(r.ShortAnswerSection),
		AnswerSection:      r.postToAnswer(r.AnswerSection),
	}
}

func (r Request) postToText(post *TextSectionPost) *entity.TextSection {
	if post == nil {
		return nil
	}
	return &entity.TextSection{
		Content: post.Content,
	}
}

func (r Request) postToChoice(post *ChoiceSectionPost) *entity.ChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.ChoiceSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (r Request) postToMultiChoice(post *MultiChoiceSectionPost) *entity.MultiChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.MultiChoiceSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (r Request) postToShortAnswer(post *ShortAnswerSectionPost) *entity.ShortAnswerSection {
	if post == nil {
		return nil
	}
	return &entity.ShortAnswerSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Question: post.Question,
	}
}

func (r Request) postToAnswer(post *AnswerSectionPost) *entity.AnswerSection {
	if post == nil {
		return nil
	}
	return &entity.AnswerSection{
		Question: post.Question,
		MaxScore: post.MaxScore,
	}
}
