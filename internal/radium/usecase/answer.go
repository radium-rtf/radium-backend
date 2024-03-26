package usecase

import (
	"context"
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	ans "github.com/radium-rtf/radium-backend/internal/radium/lib/answer"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type AnswerUseCase struct {
	section postgres2.Section
	answer  postgres2.Answer
	file    postgres2.File
	checker ans.Checker
}

func NewAnswerUseCase(section postgres2.Section, answer postgres2.Answer, file postgres2.File) AnswerUseCase {
	return AnswerUseCase{section: section, answer: answer, file: file}
}

func (uc AnswerUseCase) Create(ctx context.Context, answer *entity2.Answer) (*entity2.Answer, error) {
	section, err := uc.section.GetById(ctx, answer.SectionId)
	if err != nil {
		return nil, err
	}

	if !section.MaxAttempts.Valid {
		return uc.createAnswer(ctx, section, answer)
	}

	count, err := uc.answer.GetCountBySectionAndUserId(ctx, answer.UserId, section.Id)
	if err != nil {
		return nil, err
	}
	if int(section.MaxAttempts.Int16) == count {
		return nil, errors.New("достигнуто максимальное количество попыток")
	}
	return uc.createAnswer(ctx, section, answer)
}

func (uc AnswerUseCase) createAnswer(ctx context.Context, section *entity2.Section, answer *entity2.Answer) (
	*entity2.Answer, error) {
	verdict, err := uc.checker.Check(section, answer)
	if err != nil {
		return nil, err
	}
	answer.Verdict = verdict.Verdict
	if answer.Type != entity2.FileType {
		return uc.answer.Create(ctx, answer)
	}

	file, err := uc.file.Get(ctx, answer.FileUrl.String)
	if err != nil {
		return nil, err
	}
	isCorrectType := slices.ContainsFunc(section.FileTypes, func(s string) bool {
		return strings.HasSuffix(file.Name, s)
	})
	if !isCorrectType {
		return nil, errors.New("неверный тип файла")
	}
	return uc.answer.Create(ctx, answer)
}

func (uc AnswerUseCase) GetBySections(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (
	*entity2.AnswersCollection, error) {
	return uc.answer.Get(ctx, userId, ids)
}

func (uc AnswerUseCase) GetByUserIdAndSectionId(ctx context.Context, userId, sectionsId uuid.UUID) (
	*entity2.AnswersCollection, error) {
	return uc.answer.GetByUserIdAnsSectionId(ctx, userId, sectionsId)
}
