package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type SectionUseCase struct {
	sectionRepo postgres.Section
}

func NewSectionUseCase(sectionRepo postgres.Section) SectionUseCase {
	return SectionUseCase{sectionRepo: sectionRepo}
}

func (uc SectionUseCase) Create(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	return uc.sectionRepo.Create(ctx, section)
}

func (uc SectionUseCase) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	return uc.sectionRepo.Delete(ctx, id, isSoft)
}

func (uc SectionUseCase) Update(ctx context.Context, section *entity.Section, userId uuid.UUID) (*entity.Section, error) {
	return uc.sectionRepo.Update(ctx, section)
}
