package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	ans "github.com/radium-rtf/radium-backend/internal/lib/answer"
)

type AnswerUseCase struct {
	sectionRepo postgres.Section
	answerRepo  postgres.Answer
	checker     ans.Checker
}

func NewAnswerUseCase(sectionRepo postgres.Section, answerRepo postgres.Answer) AnswerUseCase {
	return AnswerUseCase{sectionRepo: sectionRepo, answerRepo: answerRepo}
}

func (uc AnswerUseCase) Create(ctx context.Context, answer *entity.Answer) (*entity.Answer, error) {
	section, err := uc.sectionRepo.GetSectionById(ctx, answer.SectionId)
	if err != nil {
		return nil, err
	}

	verdict, err := uc.checker.Check(section, answer)
	if err != nil {
		return nil, err
	}
	answer.Verdict = verdict.Verdict

	return answer, uc.answerRepo.Create(ctx, answer)
}
