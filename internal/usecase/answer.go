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
	answerRepo  repo.AnswerRepo
	mapper      mapper.Answer
}

func NewAnswerUseCase(pg *db.Query) AnswerUseCase {
	return AnswerUseCase{sectionRepo: repo.NewSectionRepo(pg), answerRepo: repo.NewAnswerRepo(pg)}
}

func (uc AnswerUseCase) Answer(ctx context.Context, answer *entity.Answer) (*entity.Answer, error) {
	section, err := uc.sectionRepo.GetSectionById(ctx, answer.SectionId)
	if err != nil {
		return nil, err
	}
	var verdict entity.Verdict

	switch {
	case answer.MultiChoice != nil:
		verdict = uc.multiChoice(answer.MultiChoice, section.MultiChoiceSection)
	case answer.Choice != nil:
		verdict = uc.choice(answer.Choice, section.ChoiceSection)
	case answer.ShortAnswer != nil:
		verdict = uc.shortAnswer(answer.ShortAnswer, section.ShortAnswerSection)
	default:
		return nil, errors.New("ответы должны быть")
	}

	answer.Verdict = verdict

	return answer, uc.answerRepo.CreateOrUpdate(ctx, answer)
}

func (uc AnswerUseCase) multiChoice(answer *entity.MultichoiceSectionAnswer, section *entity.MultiChoiceSection) entity.Verdict {
	answerArr := []string(answer.Answer)
	solutionArr := []string(section.Answer)
	verdict := entity.VerdictOK
	ok := reflect.DeepEqual(answerArr, solutionArr)
	if !ok {
		verdict = entity.VerdictWA
	}
	return verdict
}

func (uc AnswerUseCase) choice(answer *entity.ChoiceSectionAnswer, section *entity.ChoiceSection) entity.Verdict {
	verdict := entity.VerdictOK
	ok := answer.Answer == section.Answer
	if !ok {
		verdict = entity.VerdictWA
	}
	return verdict
}

func (uc AnswerUseCase) shortAnswer(answer *entity.ShortAnswerSectionAnswer, section *entity.ShortAnswerSection) entity.Verdict {
	verdict := entity.VerdictOK
	ok := answer.Answer == section.Answer
	if !ok {
		verdict = entity.VerdictWA
	}
	return verdict
}
