package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type PageUseCase struct {
	pageRepo   postgres.Page
	answerRepo postgres.Answer
}

func NewPageUseCase(pageRepo postgres.Page, answerRepo postgres.Answer) PageUseCase {
	return PageUseCase{pageRepo: pageRepo, answerRepo: answerRepo}
}

func (uc PageUseCase) Create(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	return uc.pageRepo.Create(ctx, page)
}

func (uc PageUseCase) GetById(ctx context.Context, id uuid.UUID, userId *uuid.UUID) (*model.Page, error) {
	page, err := uc.pageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userId == nil {
		return model.NewPage(page, map[uuid.UUID]*entity.Answer{}), err
	}
	sectionsIds := make([]uuid.UUID, 0, len(page.Sections))
	for _, section := range page.Sections {
		sectionsIds = append(sectionsIds, section.Id)
	}
	answers, err := uc.answerRepo.Get(ctx, *userId, sectionsIds) // TODO: это надо перенести в answerUseCase ?!?!
	if err != nil {
		return nil, err
	}
	p := model.NewPage(page, answers)
	return p, err
}

func (uc PageUseCase) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	return uc.pageRepo.Delete(ctx, id, isSoft)
}
