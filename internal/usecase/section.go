package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type SectionUseCase struct {
	sectionRepo repo.SectionRepo
}

func NewSectionUseCase(pg *postgres.Postgres) SectionUseCase {
	return SectionUseCase{sectionRepo: repo.NewSectionRepo(pg)}
}

func (uc SectionUseCase) CreateText(ctx context.Context, sectionPost entity.SectionTextPost) (entity.SectionText, error) {
	id, err := uc.sectionRepo.CreateText(ctx, sectionPost)
	if err != nil {
		return entity.SectionText{}, err
	}
	section := entity.SectionText{
		Id:       id,
		SlideId:  sectionPost.SlideId,
		OrderBy:  sectionPost.OrderBy,
		Markdown: sectionPost.Markdown,
	}

	return section, err
}
