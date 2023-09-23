package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Section struct {
		PageId   uuid.UUID `json:"pageId"`
		Order    uint      `json:"order" validate:"number"`
		MaxScore uint      `json:"maxScore,omitempty" validate:"min=0,max=300"`

		TextSection        *TextSectionPost        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionPost      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionPost `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionPost `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSectionPost      `json:"answer,omitempty"`
		CodeSection        *CodeSection            `json:"code,omitempty"`
	}

	TextSectionPost struct {
		Content string `validate:"required,max=10000"`
	}

	ChoiceSectionPost struct {
		Question string   `validate:"required,max=1000"`
		Answer   string   `validate:"required,max=100"`
		Variants []string `swaggertype:"array,string" validate:"required,min=2,max=6,dive,required,max=100"`
	}

	MultiChoiceSectionPost struct {
		Question string   `validate:"required,max=1000"`
		Answer   []string `swaggertype:"array,string" validate:"required,max=6,dive,required,max=100"`
		Variants []string `swaggertype:"array,string" validate:"required,min=2,max=6,dive,required,max=100"`
	}

	ShortAnswerSectionPost struct {
		Question string `validate:"required,max=100"`
		Answer   string `validate:"required,max=10"`
	}

	AnswerSectionPost struct {
		Question string `validate:"required,max=3000"`
	}

	CodeSection struct {
		Question string `validate:"required,max=5000"`
	}
)

func (r Section) toSection() *entity.Section {
	if r.TextSection != nil {
		r.MaxScore = 0
	}
	return &entity.Section{
		PageId:             r.PageId,
		Order:              r.Order,
		MaxScore:           r.MaxScore,
		TextSection:        r.postToText(r.TextSection),
		ChoiceSection:      r.postToChoice(r.ChoiceSection),
		MultiChoiceSection: r.postToMultiChoice(r.MultiChoiceSection),
		ShortAnswerSection: r.postToShortAnswer(r.ShortAnswerSection),
		AnswerSection:      r.postToAnswer(r.AnswerSection),
		CodeSection:        r.postToCode(r.CodeSection),
	}
}

func (r Section) postToText(post *TextSectionPost) *entity.TextSection {
	if post == nil {
		return nil
	}
	return &entity.TextSection{
		Content: post.Content,
	}
}

func (r Section) postToChoice(post *ChoiceSectionPost) *entity.ChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.ChoiceSection{
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (r Section) postToMultiChoice(post *MultiChoiceSectionPost) *entity.MultiChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.MultiChoiceSection{
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (r Section) postToShortAnswer(post *ShortAnswerSectionPost) *entity.ShortAnswerSection {
	if post == nil {
		return nil
	}
	return &entity.ShortAnswerSection{
		Answer:   post.Answer,
		Question: post.Question,
	}
}

func (r Section) postToAnswer(post *AnswerSectionPost) *entity.AnswerSection {
	if post == nil {
		return nil
	}
	return &entity.AnswerSection{
		Question: post.Question,
	}
}

func (r Section) postToCode(post *CodeSection) *entity.CodeSection {
	if post == nil {
		return nil
	}
	return &entity.CodeSection{
		Question: post.Question,
	}
}
