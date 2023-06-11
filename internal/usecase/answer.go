package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"reflect"
)

type AnswerUseCase struct {
	sectionRepo repo.SectionRepo
	mapper      mapper.Answer
}

func NewAnswerUseCase(pg *db.Query) AnswerUseCase {
	return AnswerUseCase{sectionRepo: repo.NewSectionRepo(pg)}
}

func (uc AnswerUseCase) Answer(ctx context.Context, post *entity.AnswerPost) (*entity.Answer, error) {
	section, err := uc.getSection(ctx, post)
	if err != nil {
		return nil, err
	}
	var verdict entity.Verdict
	var score uint

	switch {
	case post.MultiChoice != nil:
		verdict, score = uc.multiChoice(post.MultiChoice, section.MultiChoiceSection)
	case post.Choice != nil:
		verdict, score = uc.choice(post.Choice, section.ChoiceSection)
	case post.ShortAnswer != nil:
		verdict, score = uc.shortAnswer(post.ShortAnswer, section.ShortAnswerSection)
	default:
		return nil, errors.New("ответы должны быть")
	}
	return uc.mapper.PostToAnswer(post, verdict, score), nil
}

func (uc AnswerUseCase) multiChoice(answer *entity.MultichoiceSectionAnswer, section *entity.MultiChoiceSection) (entity.Verdict, uint) {
	answerArr := []string(answer.Answer)
	solutionArr := []string(section.Answer)
	verdict := entity.VerdictOK
	score := section.MaxScore
	ok := reflect.DeepEqual(answerArr, solutionArr)
	if !ok {
		verdict = entity.VerdictWA
		score = 0
	}
	return verdict, score
}

func (uc AnswerUseCase) choice(answer *entity.ChoiceSectionAnswer, section *entity.ChoiceSection) (entity.Verdict, uint) {
	verdict := entity.VerdictOK
	score := section.MaxScore
	ok := answer.Answer == section.Answer
	if !ok {
		verdict = entity.VerdictWA
		score = 0
	}
	return verdict, score
}

func (uc AnswerUseCase) shortAnswer(answer *entity.ShortAnswerSectionAnswer, section *entity.ShortAnswerSection) (entity.Verdict, uint) {
	verdict := entity.VerdictOK
	score := section.MaxScore
	ok := answer.Answer == section.Answer
	if !ok {
		verdict = entity.VerdictWA
		score = 0
	}
	return verdict, score
}

func (uc AnswerUseCase) getSection(ctx context.Context, post *entity.AnswerPost) (*entity.Section, error) {
	switch {
	case post.MultiChoice != nil:
		return uc.sectionRepo.GetSectionById(ctx, post.MultiChoice.ID)
	case post.ShortAnswer != nil:
		return uc.sectionRepo.GetSectionById(ctx, post.ShortAnswer.ID)
	case post.Choice != nil:
		return uc.sectionRepo.GetSectionById(ctx, post.Choice.ID)
	}
	return nil, errors.New("ответы должны быть")
}
