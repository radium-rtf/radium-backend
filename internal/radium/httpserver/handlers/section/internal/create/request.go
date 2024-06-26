package create

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
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
		MappingSection     *MappingSection     `json:"mapping,omitempty"`
		FileSection        *FileSection        `json:"file,omitempty"`
		MediaSection       *MediaSection       `json:"media,omitempty"`
	}

	TextSection struct {
		Content string `validate:"required,max=4096"`
	}

	MediaSection struct {
		Url string `validate:"required,url"`
	}

	ChoiceSection struct {
		Question string   `validate:"required,max=4096"`
		Answer   string   `validate:"required,max=256"`
		Variants []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
	}

	MultiChoiceSection struct {
		Question string   `validate:"required,max=4096"`
		Answer   []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
		Variants []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
	}

	ShortAnswerSection struct {
		Question string `validate:"required,max=4096"`
		Answer   string `validate:"required,max=256"`
	}

	AnswerSection struct {
		Question string `validate:"required,max=4096"`
	}

	CodeSection struct {
		Question string `validate:"required,max=4096"`
	}

	PermutationSection struct {
		Question string   `validate:"required,max=4096"`
		Answer   []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
	}

	MappingSection struct {
		Question string `validate:"required,max=4096"`

		Keys   []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
		Answer []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
	}

	FileSection struct {
		Question  string   `validate:"required,max=4096"`
		FileTypes []string `validate:"required,min=1,dive,startswith=."`
	}
)

func (r Section) ToSection() (*entity.Section, error) {
	// TODO: хз хз хз. хочется валидацию разную для разных секций и без вот этого
	maxAttempts := sql.NullInt16{Int16: r.MaxAttempts, Valid: r.MaxAttempts != 0}

	switch {
	case r.MappingSection != nil:
		keys := r.MappingSection.Keys
		answer := r.MappingSection.Answer
		if len(keys) != len(answer) {
			return nil, errors.New("секция с сопоставлением должна иметь одинаковое количество строк и обоих столбцах")
		}
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.MappingSection.Question,
			"", answer, answer, keys, nil, entity.MappingType, "")

	case r.PermutationSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.PermutationSection.Question,
			"", []string{}, r.PermutationSection.Answer, []string{}, nil, entity.PermutationType, "")

	case r.ChoiceSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.ChoiceSection.Question,
			r.ChoiceSection.Answer, r.ChoiceSection.Variants, []string{}, []string{}, nil, entity.ChoiceType, "")

	case r.ShortAnswerSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.ShortAnswerSection.Question,
			r.ShortAnswerSection.Answer, []string{}, []string{}, []string{}, nil, entity.ShortAnswerType, "")

	case r.MultiChoiceSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.MultiChoiceSection.Question,
			"", r.MultiChoiceSection.Variants, r.MultiChoiceSection.Answer, []string{}, nil, entity.MultiChoiceType, "")

	case r.TextSection != nil:
		return entity.NewSection(sql.NullInt16{}, r.PageId, r.Order, r.MaxScore, r.TextSection.Content,
			"", []string{}, []string{}, []string{}, nil, entity.TextType, "")

	case r.CodeSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.CodeSection.Question,
			"", []string{}, []string{}, []string{}, nil, entity.CodeType, "")

	case r.AnswerSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.AnswerSection.Question,
			"", []string{}, []string{}, []string{}, nil, entity.AnswerType, "")
	case r.FileSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, r.FileSection.Question,
			"", []string{}, []string{}, []string{}, r.FileSection.FileTypes, entity.FileType, "")
	case r.MediaSection != nil:
		return entity.NewSection(maxAttempts, r.PageId, r.Order, r.MaxScore, "",
			"", []string{}, []string{}, []string{}, []string{}, entity.MediaType, r.MediaSection.Url)
	default:
		return nil, errors.New("не удалось создать секцию")
	}
}
