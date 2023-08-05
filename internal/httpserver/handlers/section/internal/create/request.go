package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Section struct {
		PageId   uuid.UUID `json:"pageId"`
		Order    uint      `json:"order"`
		MaxScore uint      `json:"maxScore,omitempty"`

		TextSection        *TextSectionPost        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionPost      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionPost `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionPost `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSectionPost      `json:"answer,omitempty"`
		CodeSection        *CodeSection            `json:"code,omitempty"`
	}

	TextSectionPost struct {
		Content string
	}

	ChoiceSectionPost struct {
		Question string
		Answer   string
		Variants []string `swaggertype:"array,string"`
	}

	MultiChoiceSectionPost struct {
		Question string
		Answer   []string `swaggertype:"array,string"`
		Variants []string `swaggertype:"array,string"`
	}

	ShortAnswerSectionPost struct {
		Question string
		Answer   string
	}

	AnswerSectionPost struct {
		Question string
	}

	CodeSection struct {
		Question string
	}
)

func (r Section) toSection() *entity.Section {
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
