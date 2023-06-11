package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	Answer struct {
		Verdict     Verdict
		Score       uint
		Choice      *ChoiceSectionAnswer
		MultiChoice *MultichoiceSectionAnswer
		ShortAnswer *ShortAnswerSectionAnswer
	}

	AnswerDto struct {
		Verdict Verdict `json:"verdict"`
		Score   uint    `json:"score"`
	}

	AnswerPost struct {
		Choice      *ChoiceSectionAnswer      `json:"choice,omitempty"`
		MultiChoice *MultichoiceSectionAnswer `json:"multiChoice,omitempty"`
		ShortAnswer *ShortAnswerSectionAnswer `json:"shortAnswer,omitempty"`
	}

	ChoiceSectionAnswer struct {
		ID     uuid.UUID `json:"id"`
		Answer string    `json:"answer"`
	}

	MultichoiceSectionAnswer struct {
		ID     uuid.UUID      `json:"id"`
		Answer pq.StringArray `json:"answer" swaggertype:"array,string"`
	}

	ShortAnswerSectionAnswer struct {
		ID     uuid.UUID `json:"id"`
		Answer string    `json:"answer"`
	}
)
