package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
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
	section, err := uc.sectionRepo.GetById(ctx, answer.SectionId)
	if err != nil {
		return nil, err
	}
	if !section.MaxAttempts.Valid {
		return uc.createAnswer(ctx, section, answer)
	}

	count, err := uc.answerRepo.GetCountBySectionAndUserId(ctx, answer.UserId, section.Id)
	if err != nil {
		return nil, err
	}
	if int(section.MaxAttempts.Int16) == count {
		return nil, errors.New("достигнуто масимальное количество попыток")
	}

	return uc.createAnswer(ctx, section, answer)
}

func (uc AnswerUseCase) createAnswer(ctx context.Context, section *entity.Section, answer *entity.Answer) (
	*entity.Answer, error) {
	verdict, err := uc.checker.Check(section, answer)
	if err != nil {
		return nil, err
	}
	answer.Verdict = verdict.Verdict

	return answer, uc.answerRepo.Create(ctx, answer)
}

func (uc AnswerUseCase) GetBySections(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (
	*entity.AnswersCollection, error) {
	return uc.answerRepo.Get(ctx, userId, ids)
}

func (uc AnswerUseCase) GetByUserIdAndSectionId(ctx context.Context, userId, sectionsId uuid.UUID) (
	*entity.AnswersCollection, error) {
	return uc.answerRepo.GetByUserIdAnsSectionId(ctx, userId, sectionsId)
}
