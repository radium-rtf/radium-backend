package mapper

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Answer struct {
}

func (a Answer) Answer(answer *entity.Answer) entity.VerdictDto {
	return entity.VerdictDto{Verdict: answer.Verdict}
}
