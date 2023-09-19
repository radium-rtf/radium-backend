package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type PageUseCase struct {
	pageRepo postgres.Page
}

func NewPageUseCase(pageRepo postgres.Page) PageUseCase {
	return PageUseCase{pageRepo: pageRepo}
}

func (uc PageUseCase) Create(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	return uc.pageRepo.Create(ctx, page)
}

func (uc PageUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	return uc.pageRepo.GetByID(ctx, id)
}

func (uc PageUseCase) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	return uc.pageRepo.Delete(ctx, id, isSoft)
}

func (uc PageUseCase) Update(ctx context.Context, page *entity.Page, userId uuid.UUID) (*entity.Page, error) {
	return uc.pageRepo.Update(ctx, page)
}
