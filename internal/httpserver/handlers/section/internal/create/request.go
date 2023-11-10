package create

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Section struct {
		PageId      uuid.UUID `json:"pageId"`
		Order       float64   `json:"order" validate:"number"`
		MaxScore    uint      `json:"maxScore,omitempty" validate:"min=0,max=300"`
		MaxAttempts int16     `json:"maxAttempts,omitempty" validate:"min=0,max=200"`

		TextSection        *TextSection        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSection      `json:"answer,omitempty"`
		CodeSection        *CodeSection        `json:"code,omitempty"`
		PermutationSection *PermutationSection `json:"permutation,omitempty"`
		MappingSection     *MappingSection     `json:"mappingSection,omitempty"`
	}

	TextSection struct {
		Content string `validate:"required,max=10000"`
	}

	ChoiceSection struct {
		Question string   `validate:"required,max=1000"`
		Answer   string   `validate:"required,max=100"`
		Variants []string `swaggertype:"array,string" validate:"required,min=2,max=6,dive,required,max=100"`
	}

	MultiChoiceSection struct {
		Question string   `validate:"required,max=1000"`
		Answer   []string `swaggertype:"array,string" validate:"required,max=6,dive,required,max=100"`
		Variants []string `swaggertype:"array,string" validate:"required,min=2,max=6,dive,required,max=100"`
	}

	ShortAnswerSection struct {
		Question string `validate:"required,max=200"`
		Answer   string `validate:"required,max=50"`
	}

	AnswerSection struct {
		Question string `validate:"required,max=3000"`
	}

	CodeSection struct {
		Question string `validate:"required,max=5000"`
	}

	PermutationSection struct {
		Question string   `validate:"required,max=500"`
		Answer   []string `swaggertype:"array,string" validate:"required,max=8,dive,required,max=100"`
	}

	MappingSection struct {
		Question string `validate:"required,max=800"`

		Keys   []string `swaggertype:"array,string" validate:"required,max=10,dive,required,max=150"`
		Answer []string `swaggertype:"array,string" validate:"required,max=10,dive,required,max=150"`
	}
)

func (r Section) ToSection() (*entity.Section, error) {
	// TODO: хз хз хз. хочется валидацию разную для разных секций и без вот этого
	maxAttempts := sql.NullInt16{Int16: r.MaxAttempts, Valid: r.MaxAttempts != 0}

	switch {
	case r.MappingSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.MappingSection.Question,
			"", []string{}, r.MappingSection.Answer, entity.MappingType, r.MappingSection.Keys)

	case r.PermutationSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.PermutationSection.Question,
			"", []string{}, r.PermutationSection.Answer, entity.PermutationType, []string{})

	case r.ChoiceSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.ChoiceSection.Question,
			r.ChoiceSection.Answer, r.ChoiceSection.Variants, []string{}, entity.ChoiceType, []string{})

	case r.ShortAnswerSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.ShortAnswerSection.Question,
			r.ShortAnswerSection.Answer, []string{}, []string{}, entity.ShortAnswerType, []string{})

	case r.MultiChoiceSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.MultiChoiceSection.Question,
			"", r.MultiChoiceSection.Variants, r.MultiChoiceSection.Answer, entity.MultiChoiceType, []string{})

	case r.TextSection != nil:
		return entity.NewSection(sql.NullInt16{}, r.PageId, r.Order, r.MaxScore, r.TextSection.Content,
			"", []string{}, []string{}, entity.TextType, []string{})

	case r.CodeSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.CodeSection.Question,
			"", []string{}, []string{}, entity.CodeType, []string{})

	case r.AnswerSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.AnswerSection.Question,
			"", []string{}, []string{}, entity.AnswerType, []string{})
	default:
		return nil, errors.New("не удалось создать секцию")
	}
}
