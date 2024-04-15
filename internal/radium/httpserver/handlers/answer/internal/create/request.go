package create

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
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
		Mapping     *Mapping       `json:"mapping,omitempty"`
		File        *File          `json:"file,omitempty"`
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

	File struct {
		Answer string `json:"answer" validate:"required,url"`
	}

	Code struct {
		Answer   string `json:"answer"`
		Language string `json:"lang"`
	}

	Permutation struct {
		Answer []string `swaggertype:"array,string"`
	}

	Mapping struct {
		Answer []string `swaggertype:"array,string"`
	}
)

func (r *Answer) ToAnswer(userId uuid.UUID) (*entity2.Answer, error) {
	answer := &entity2.Answer{
		DBModel:   entity2.DBModel{Id: uuid.New()},
		SectionId: r.SectionId,
		UserId:    userId,
	}

	switch {
	case r.Choice != nil:
		answer.Answer = r.Choice.Answer
		answer.Type = entity2.ChoiceType
	case r.MultiChoice != nil:
		answer.Answers = r.MultiChoice.Answer
		answer.Type = entity2.MultiChoiceType
	case r.ShortAnswer != nil:
		answer.Answer = r.ShortAnswer.Answer
		answer.Type = entity2.ShortAnswerType
	case r.Answer != nil:
		answer.Answer = r.Answer.Answer
		answer.Type = entity2.AnswerType
	case r.Code != nil:
		answer.Answer = r.Code.Answer
		answer.Type = entity2.CodeType
		answer.Language = r.Code.Language
	case r.Permutation != nil:
		answer.Type = entity2.PermutationType
		answer.Answers = r.Permutation.Answer
	case r.Mapping != nil:
		answer.Type = entity2.MappingType
		answer.Answers = r.Mapping.Answer
	case r.File != nil:
		answer.Type = entity2.FileType
		answer.FileUrl = sql.NullString{String: r.File.Answer, Valid: r.File.Answer == ""}
	default:
		return nil, errors.New("create.Answer - toAnswer - не удалось создать ответ")
	}

	return answer, nil
}
