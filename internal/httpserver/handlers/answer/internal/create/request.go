package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type (
	Answer struct {
		SectionId   uuid.UUID                     `json:"id"`
		Choice      *ChoiceSectionAnswerPost      `json:"choice,omitempty"`
		MultiChoice *MultichoiceSectionAnswerPost `json:"multiChoice,omitempty"`
		ShortAnswer *ShortAnswerSectionAnswerPost `json:"shortAnswer,omitempty"`
		Answer      *AnswerSectionAnswerPost      `json:"answer,omitempty"`
		Code        *CodeSectionAnswerPost        `json:"code,omitempty"`
	}

	MultichoiceSectionAnswerPost struct {
		Answer []string `json:"answer" swaggertype:"array,string"`
	}

	AnswerSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ChoiceSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	ShortAnswerSectionAnswerPost struct {
		Answer string `json:"answer"`
	}

	CodeSectionAnswerPost struct {
		Answer   string `json:"answer"`
		Language string `json:"lang"`
	}
)

func (r *Answer) ToAnswer(userId uuid.UUID) *entity.Answer {
	var (
		choice      *entity.ChoiceSectionAnswer
		multichoice *entity.MultichoiceSectionAnswer
		shortAnswer *entity.ShortAnswerSectionAnswer
		answer      *entity.AnswerSectionAnswer
		code        *entity.CodeSectionAnswer
	)

	if r.Choice != nil {
		choice = &entity.ChoiceSectionAnswer{Answer: r.Choice.Answer}
	}

	if r.MultiChoice != nil {
		multichoice = &entity.MultichoiceSectionAnswer{Answer: r.MultiChoice.Answer}
	}

	if r.ShortAnswer != nil {
		shortAnswer = &entity.ShortAnswerSectionAnswer{Answer: r.ShortAnswer.Answer}
	}

	if r.Answer != nil {
		answer = &entity.AnswerSectionAnswer{Answer: r.Answer.Answer}
	}

	if r.Code != nil {
		code = &entity.CodeSectionAnswer{Answer: r.Code.Answer, Language: r.Code.Language}
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
	}
}
