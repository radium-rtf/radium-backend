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

func (uc SectionUseCase) CreateSection(ctx context.Context, section *entity.Section) (*entity.Section, error) {
	return uc.sectionRepo.CreateSection(ctx, section)
}
