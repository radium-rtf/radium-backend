package update

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Section struct {
		MaxScore    uint  `json:"maxScore,omitempty" validate:"min=0,max=300"`
		MaxAttempts uint8 `json:"maxAttempts,omitempty" validate:"min=0,max=200"`

		TextSection        *TextSection        `json:"text,omitempty"`
		ChoiceSection      *ChoiceSection      `json:"choice,omitempty"`
		MultiChoiceSection *MultiChoiceSection `json:"multichoice,omitempty"`
		ShortAnswerSection *ShortAnswerSection `json:"shortanswer,omitempty"`
		AnswerSection      *AnswerSection      `json:"answer,omitempty"`
		CodeSection        *CodeSection        `json:"code,omitempty"`
		PermutationSection *PermutationSection `json:"permutation,omitempty"`
	}

	TextSection struct {
		Content string `validate:"max=10000"`
	}

	ChoiceSection struct {
		Question string   `validate:"max=1000"`
		Answer   string   `validate:"max=100"`
		Variants []string `swaggertype:"array,string" validate:"min=2,max=6,dive,required,max=100"`
	}

	MultiChoiceSection struct {
		Question string   `validate:"max=1000"`
		Answer   []string `swaggertype:"array,string" validate:"max=6,dive,required,max=100"`
		Variants []string `swaggertype:"array,string" validate:"min=2,max=6,dive,required,max=100"`
	}

	ShortAnswerSection struct {
		Question string `validate:"required,max=200"`
		Answer   string `validate:"required,max=50"`
	}

	AnswerSection struct {
		Question string `validate:"max=3000"`
	}

	CodeSection struct {
		Question string `validate:"max=5000"`
	}

	PermutationSection struct {
		Question string   `validate:"max=500"`
		Answer   []string `swaggertype:"array,string" validate:"max=8,dive,required,max=100"`
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
			"", []string{}, r.PermutationSection.Answer, entity.PermutationType, []string{})

	case r.ChoiceSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.ChoiceSection.Question,
			r.ChoiceSection.Answer, r.ChoiceSection.Variants, []string{}, entity.ChoiceType, []string{})

	case r.ShortAnswerSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.ShortAnswerSection.Question,
			r.ShortAnswerSection.Answer, []string{}, []string{}, entity.ShortAnswerType, []string{})

	case r.MultiChoiceSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.MultiChoiceSection.Question,
			"", r.MultiChoiceSection.Variants, r.MultiChoiceSection.Answer, entity.MultiChoiceType, []string{})

	case r.TextSection != nil:
		section, err = entity.NewSection(sql.NullInt16{}, pageId, 0, r.MaxScore, r.TextSection.Content,
			"", []string{}, []string{}, entity.TextType, []string{})

	case r.CodeSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.CodeSection.Question,
			"", []string{}, []string{}, entity.CodeType, []string{})

	case r.AnswerSection != nil:
		section, err = entity.NewSection(maxAttempts, pageId, 0, r.MaxScore, r.AnswerSection.Question,
			"", []string{}, []string{}, entity.AnswerType, []string{})
	default:
	}
	section.Id = sectionId
	section.MaxScore = r.MaxScore
	return section, err
}