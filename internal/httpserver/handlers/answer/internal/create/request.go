package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type (
	Answer struct {
		SectionId   uuid.UUID      `json:"id"`
		Choice      *Choice        `json:"choice,omitempty"`
		MultiChoice *MultiChoice   `json:"multiChoice,omitempty"`
		ShortAnswer *ShortAnswer   `json:"shortAnswer,omitempty"`
		Answer      *AnswerSection `json:"answer,omitempty"`
		Code        *Code          `json:"code,omitempty"`
		Permutation *Permutation   `json:"permutation,omitempty"`
	}

	MultiChoice struct {
		Answer []string `json:"answer" swaggertype:"array,string"`
	}

	AnswerSection struct {
		Answer string `json:"answer"`
	}

	Choice struct {
		Answer string `json:"answer"`
	}

	ShortAnswer struct {
		Answer string `json:"answer"`
	}

	Code struct {
		Answer   string `json:"answer"`
		Language string `json:"lang"`
	}

	Permutation struct {
		Answer []string `swaggertype:"array,string" validate:"required,max=8,dive,required,max=100"`
	}
)

func (r *Answer) ToAnswer(userId uuid.UUID) *entity.Answer {
	var (
		choice      *entity.ChoiceSectionAnswer
		multichoice *entity.MultichoiceSectionAnswer
		shortAnswer *entity.ShortAnswerSectionAnswer
		answer      *entity.AnswerSectionAnswer
		code        *entity.CodeSectionAnswer
		permutatuon *entity.PermutationSectionAnswer
	)

	switch {
	case r.Choice != nil:
		choice = &entity.ChoiceSectionAnswer{Answer: r.Choice.Answer}
	case r.MultiChoice != nil:
		multichoice = &entity.MultichoiceSectionAnswer{Answer: r.MultiChoice.Answer}
	case r.ShortAnswer != nil:
		shortAnswer = &entity.ShortAnswerSectionAnswer{Answer: r.ShortAnswer.Answer}
	case r.Answer != nil:
		answer = &entity.AnswerSectionAnswer{Answer: r.Answer.Answer}
	case r.Code != nil:
		code = &entity.CodeSectionAnswer{Answer: r.Code.Answer, Language: r.Code.Language}
	case r.Permutation != nil:
		permutatuon = &entity.PermutationSectionAnswer{Answer: r.Permutation.Answer}
	}

	return &entity.Answer{
		Verdict:     verdict.EMPTY,
		SectionId:   r.SectionId,
		UserId:      userId,
		Choice:      choice,
		MultiChoice: multichoice,
		ShortAnswer: shortAnswer,
		Answer:      answer,
		Code:        code,
		Permutation: permutatuon,
	}
}
