package update

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/httpserver/handlers/section/internal/create"
)

type (
	Section struct {
		MaxScore    uint  `json:"maxScore,omitempty" validate:"min=0,max=300"`
		MaxAttempts uint8 `json:"maxAttempts,omitempty" validate:"min=0,max=200"`

		TextSection        *create.TextSection `json:"text,omitempty"`
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

	ChoiceSection struct {
		Question string   `validate:"max=4096"`
		Answer   string   `validate:"max=256"`
		Variants []string `swaggertype:"array,string" validate:"min=1,max=10,dive,required,max=256"`
	}

	MediaSection struct {
		Url string `validate:"required,url"`
	}

	MultiChoiceSection struct {
		Question string   `validate:"max=4096"`
		Answer   []string `swaggertype:"array,string" validate:"min=1,max=10,dive,required,max=256"`
		Variants []string `swaggertype:"array,string" validate:"min=1,max=10,dive,required,max=256"`
	}

	ShortAnswerSection struct {
		Question string `validate:"max=4096"`
		Answer   string `validate:"max=256"`
	}

	AnswerSection struct {
		Question string `validate:"max=4096"`
	}

	CodeSection struct {
		Question string `validate:"max=4096"`
	}

	PermutationSection struct {
		Question string   `validate:"max=4096"`
		Answer   []string `swaggertype:"array,string" validate:"min=1,max=10,dive,required,max=256"`
	}

	MappingSection struct {
		Question string `validate:"max=4096"`

		Keys   []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
		Answer []string `swaggertype:"array,string" validate:"required,min=1,max=10,dive,required,max=256"`
	}

	FileSection struct {
		Question  string   `validate:"max=4096"`
		FileTypes []string `validate:"min=1,dive,startswith=."`
	}
)

func (r Section) toSection(sectionId uuid.UUID) (*entity.Section, error) {
	// TODO: хз хз хз. хочется валидацию разную для разных секций и без вот этого
	var (
		section = &entity.Section{}
		pageId  uuid.UUID
		err     error
	)
	maxAttempts := sql.NullInt16{Int16: int16(r.MaxAttempts), Valid: r.MaxAttempts != 0}

	switch {
	case r.PermutationSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.PermutationSection.Question,
			"", []string{}, r.PermutationSection.Answer, []string{}, nil, entity.PermutationType, "")

	case r.ChoiceSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.ChoiceSection.Question,
			r.ChoiceSection.Answer, r.ChoiceSection.Variants, []string{}, []string{}, nil, entity.ChoiceType, "")

	case r.ShortAnswerSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.ShortAnswerSection.Question,
			r.ShortAnswerSection.Answer, []string{}, []string{}, []string{}, nil, entity.ShortAnswerType, "")

	case r.MultiChoiceSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.MultiChoiceSection.Question,
			"", r.MultiChoiceSection.Variants, r.MultiChoiceSection.Answer, []string{}, nil, entity.MultiChoiceType, "")

	case r.TextSection != nil:
		section, err = entity.NewSection(sql.NullInt16{}, pageId, 0, r.MaxScore, r.TextSection.Content,
			"", []string{}, []string{}, []string{}, nil, entity.TextType, "")

	case r.CodeSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.CodeSection.Question,
			"", []string{}, []string{}, []string{}, nil, entity.CodeType, "")

	case r.AnswerSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.AnswerSection.Question,
			"", []string{}, []string{}, []string{}, nil, entity.AnswerType, "")

	case r.MappingSection != nil:
		keys := r.MappingSection.Keys
		answer := r.MappingSection.Answer
		if len(keys) != len(answer) {
			return nil, errors.New("секция с сопоставлением должна иметь одинаковое количество строк и обоих столбцах")
		}
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.MappingSection.Question,
			"", answer, answer, keys, nil, entity.MappingType, "")

	case r.FileSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.FileSection.Question,
			"", []string{}, []string{}, []string{}, r.FileSection.FileTypes, entity.FileType, "")
	case r.MediaSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, "",
			"", []string{}, []string{}, []string{}, []string{}, entity.MediaType, r.MediaSection.Url)
	default:
	}
	section.Id = sectionId
	section.MaxScore = r.MaxScore
	return section, err
}
