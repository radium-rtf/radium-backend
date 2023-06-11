package mapper

import "github.com/radium-rtf/radium-backend/internal/entity"

type Answer struct {
}

func (Answer) PostToAnswer(post *entity.AnswerPost, verdict entity.Verdict, score uint) *entity.Answer {
	return &entity.Answer{
		Verdict:     verdict,
		Score:       score,
		ShortAnswer: post.ShortAnswer,
		Choice:      post.Choice,
		MultiChoice: post.MultiChoice,
	}
}

func (a Answer) Answer(answer *entity.Answer) entity.AnswerDto {
	return entity.AnswerDto{Score: answer.Score, Verdict: answer.Verdict}
}
