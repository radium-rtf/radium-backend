package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"strings"
)

type SectionUseCase struct {
	sectionRepo repo.SectionRepo
}

func NewSectionUseCase(pg *postgres.Postgres) SectionUseCase {
	return SectionUseCase{sectionRepo: repo.NewSectionRepo(pg)}
}

func (uc SectionUseCase) CreateText(ctx context.Context, post entity.SectionTextPost) (entity.SectionText, error) {
	id, err := uc.sectionRepo.CreateText(ctx, post)
	if err != nil {
		return entity.SectionText{}, err
	}
	section := entity.SectionText{
		Id:       id,
		SlideId:  post.SlideId,
		OrderBy:  post.OrderBy,
		Markdown: post.Markdown,
	}

	return section, err
}

func (uc SectionUseCase) CreateQuestion(ctx context.Context, post entity.SectionQuestionPost) (
	entity.SectionQuestionDto, error) {
	id, err := uc.sectionRepo.CreateQuestion(ctx, post)
	if err != nil {
		return entity.SectionQuestionDto{}, err
	}

	section := entity.SectionQuestionDto{
		Id:            id,
		SlideId:       post.SlideId,
		OrderBy:       post.OrderBy,
		Cost:          post.Cost,
		CaseSensitive: post.CaseSensitive,
		Question:      post.Question,
	}

	return section, nil
}

func (uc SectionUseCase) CreateQuestionAnswer(ctx context.Context, post entity.SectionQuestionAnswerPost, userId string) (
	entity.SectionQuestionAnswerDto, error) {
	var answer entity.SectionQuestionAnswer
	question, err := uc.sectionRepo.GetQuestionById(ctx, post.SectionId)
	if err != nil {
		return entity.SectionQuestionAnswerDto{}, err
	}
	ok := false
	if question.CaseSensitive {
		ok = question.Answer == post.Answer
	} else {
		ok = strings.ToLower(question.Answer) == strings.ToLower(post.Answer)
	}
	verdict := entity.SectionAnswerWA
	if ok {
		verdict = entity.SectionAnswerOK
	}

	answer = entity.SectionQuestionAnswer{
		UserId:    userId,
		SectionId: post.SectionId,
		Answer:    post.Answer,
		Verdict:   verdict,
	}

	id, err := uc.sectionRepo.CreateQuestionAnswer(ctx, answer)
	answer.Id = id
	return entity.NewSectionQuestionAnswerToDto(answer), err
}

func (uc SectionUseCase) CreateChoice(ctx context.Context, post entity.SectionChoicePost) (
	entity.SectionChoiceDto, error) {
	section, err := entity.NewSectionChoicePostToSection(post)
	if err != nil {
		return entity.SectionChoiceDto{}, err
	}

	id, err := uc.sectionRepo.CreateChoice(ctx, section)
	if err != nil {
		return entity.SectionChoiceDto{}, err
	}
	section.Id = id
	sectionDto := entity.NewSectionChoiceToDto(section)
	return sectionDto, nil
}

func (uc SectionUseCase) CreateChoiceAnswer(ctx context.Context, post entity.SectionChoiceAnswerPost, userId string) (
	entity.SectionChoiceAnswerDto, error) {
	choice, err := uc.sectionRepo.GetChoiceById(ctx, post.SectionId)
	if err != nil {
		return entity.SectionChoiceAnswerDto{}, err
	}
	ok := false
	if post.Answer == choice.Answer {
		ok = true
	}
	verdict := entity.SectionAnswerWA
	if ok {
		verdict = entity.SectionAnswerOK
	}

	answer := entity.SectionChoiceAnswer{
		UserId:    userId,
		SectionId: post.SectionId,
		Answer:    post.Answer,
		Verdict:   verdict,
	}
	id, err := uc.sectionRepo.CreateChoiceAnswer(ctx, answer)
	answer.Id = id

	return entity.NewSectionChoiceAnswerToDto(answer), err
}
