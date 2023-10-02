package create

import (
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Section struct {
		PageId   uuid.UUID `json:"pageId"`
		Order    float64   `json:"order" validate:"number"`
		MaxScore uint      `json:"maxScore,omitempty" validate:"min=0,max=300"`

		TextSection        *TextSectionPost        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSectionPost      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSectionPost `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSectionPost `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSectionPost      `json:"answer,omitempty"`
		CodeSection        *CodeSection            `json:"code,omitempty"`
		PermutationSection *PermutationSection     `json:"permutation,omitempty"`
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
		Question string `validate:"required,max=200"`
		Answer   string `validate:"required,max=50"`
	}

	AnswerSectionPost struct {
		Question string `validate:"required,max=3000"`
	}

	CodeSection struct {
		Question string `validate:"required,max=5000"`
	}

	PermutationSection struct {
		Question string   `validate:"required,max=500"`
		Answer   []string `swaggertype:"array,string" validate:"required,max=8,dive,required,max=100"`
	}
)

func (r Section) toSection() (*entity.Section, error) {
	// TODO: хз хз хз. хочется валидацию разную для разных секций и без вот этого

	switch {
	case r.PermutationSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.PermutationSection.Question,
			"", []string{}, r.PermutationSection.Answer, entity.PermutationType)

	case r.ChoiceSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.ChoiceSection.Question,
			r.ChoiceSection.Answer, r.ChoiceSection.Variants, []string{}, entity.ChoiceType)

	case r.ShortAnswerSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.ShortAnswerSection.Question,
			r.ShortAnswerSection.Answer, []string{}, []string{}, entity.ShortAnswerType)

	case r.MultiChoiceSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.MultiChoiceSection.Question,
			"", r.MultiChoiceSection.Variants, r.MultiChoiceSection.Answer, entity.MultiChoiceType)

	case r.TextSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.TextSection.Content,
			"", []string{}, []string{}, entity.TextType)

	case r.CodeSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.CodeSection.Question,
			"", []string{}, []string{}, entity.CodeType)

	case r.AnswerSection != nil:
		return entity.NewSection(r.PageId, r.Order, r.MaxScore, r.AnswerSection.Question,
			"", []string{}, []string{}, entity.AnswerType)
	default:
		return nil, errors.New("не удалось создать секцию")
	}
}
