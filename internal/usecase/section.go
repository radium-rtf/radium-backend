package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type SectionUseCase struct {
	sectionRepo repo.SectionRepo
}

func NewSectionUseCase(pg *db.Query) SectionUseCase {
	return SectionUseCase{sectionRepo: repo.NewSectionRepo(pg)}
}

func (uc SectionUseCase) CreateSection(ctx context.Context, post entity.SectionPost) (*entity.Section, error) {
	return uc.sectionRepo.CreateSection(ctx, post)
}

// func (uc SectionUseCase) CreateQuestionAnswer(ctx context.Context, post entity.SectionQuestionAnswerPost, userId string) (
// 	entity.SectionQuestionAnswerDto, error) {
// 	var answer entity.SectionQuestionAnswer
// 	question, err := uc.sectionRepo.GetQuestionById(ctx, post.SectionId)
// 	if err != nil {
// 		return entity.SectionQuestionAnswerDto{}, err
// 	}
// 	ok := false
// 	if question.CaseSensitive {
// 		ok = question.Answer == post.Answer
// 	} else {
// 		ok = strings.ToLower(question.Answer) == strings.ToLower(post.Answer)
// 	}
// 	verdict := entity.VerdictWA
// 	if ok {
// 		verdict = entity.VerdictOK
// 	}

// 	answer = entity.SectionQuestionAnswer{
// 		UserId:    userId,
// 		SectionId: post.SectionId,
// 		Answer:    post.Answer,
// 		Verdict:   verdict,
// 	}

// 	id, err := uc.sectionRepo.CreateQuestionAnswer(ctx, answer)
// 	answer.Id = id
// 	return entity.NewSectionQuestionAnswerToDto(answer), err
// }

// func (uc SectionUseCase) CreateChoice(ctx context.Context, post entity.SectionChoicePost) (
// 	entity.SectionChoiceDto, error) {
// 	section, err := entity.NewSectionChoicePostToSection(post)
// 	if err != nil {
// 		return entity.SectionChoiceDto{}, err
// 	}

// 	id, err := uc.sectionRepo.CreateChoice(ctx, section)
// 	if err != nil {
// 		return entity.SectionChoiceDto{}, err
// 	}
// 	section.Id = id
// 	sectionDto := entity.NewSectionChoiceToDto(section)
// 	return sectionDto, nil
// }

// func (uc SectionUseCase) CreateChoiceAnswer(ctx context.Context, post entity.SectionChoiceAnswerPost, userId string) (
// 	entity.SectionChoiceAnswerDto, error) {
// 	choice, err := uc.sectionRepo.GetChoiceById(ctx, post.SectionId)
// 	if err != nil {
// 		return entity.SectionChoiceAnswerDto{}, err
// 	}
// 	ok := false
// 	if post.Answer == choice.Answer {
// 		ok = true
// 	}
// 	verdict := entity.VerdictWA
// 	if ok {
// 		verdict = entity.VerdictOK
// 	}

// 	answer := entity.SectionChoiceAnswer{
// 		UserId:    userId,
// 		SectionId: post.SectionId,
// 		Answer:    post.Answer,
// 		Verdict:   verdict,
// 	}
// 	id, err := uc.sectionRepo.CreateChoiceAnswer(ctx, answer)
// 	answer.Id = id

// 	return entity.NewSectionChoiceAnswerToDto(answer), err
// }

// func (uc SectionUseCase) CreateMultiChoice(ctx context.Context, post entity.SectionMultiChoicePost) (
// 	entity.SectionMultiChoiceDto, error) {
// 	section, err := entity.NewSectionMultiChoicePostToSection(post)
// 	if err != nil {
// 		return entity.SectionMultiChoiceDto{}, err
// 	}

// 	id, err := uc.sectionRepo.CreateMultiChoice(ctx, section)
// 	if err != nil {
// 		return entity.SectionMultiChoiceDto{}, err
// 	}
// 	section.Id = id
// 	sectionDto := entity.NewSectionMultiChoiceToDto(section)
// 	return sectionDto, nil
// }

// func (uc SectionUseCase) CreateMultiChoiceAnswer(ctx context.Context, post entity.SectionMultiChoiceAnswerPost, userId string) (
// 	entity.SectionMultiChoiceAnswerDto, error) {
// 	choice, err := uc.sectionRepo.GetMultiChoiceById(ctx, post.SectionId)
// 	if err != nil {
// 		return entity.SectionMultiChoiceAnswerDto{}, err
// 	}
// 	ok := false
// 	if reflect.DeepEqual(post.Answer, choice.Answer) {
// 		ok = true
// 	}
// 	verdict := entity.VerdictWA
// 	if ok {
// 		verdict = entity.VerdictOK
// 	}

// 	answer := entity.SectionMultiChoiceAnswer{
// 		UserId:    userId,
// 		SectionId: post.SectionId,
// 		Answer:    post.Answer,
// 		Verdict:   verdict,
// 	}
// 	id, err := uc.sectionRepo.CreateMultiChoiceAnswer(ctx, answer)
// 	answer.Id = id

// 	return entity.NewSectionMultiChoiceAnswerToDto(answer), err
// }
