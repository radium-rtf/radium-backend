//go:build ignore
// +build ignore

package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type SlideUseCase struct {
	slideRepo  repo.SlideRepo
	moduleRepo repo.ModuleRepo
}

func NewSlideUseCase(pg *db.Query) SlideUseCase {
	return SlideUseCase{slideRepo: repo.NewSlideRepo(pg), moduleRepo: repo.NewModuleRepo(pg)}
}

func (uc SlideUseCase) CreateSlide(ctx context.Context, createSlide entity.SlideRequest) (entity.SlideDto, error) {
	// moduleId, err := uc.moduleRepo.GetModuleId(ctx, createSlide.CourseId, createSlide.ModuleNameEng)
	// if err != nil {
	// 	return entity.SlideDto{}, err
	// }
	slide := entity.Slide{
		NameEng:  translit.RuEn(createSlide.Name),
		Name:     createSlide.Name,
		ModuleId: createSlide.ModuleId,
	}
	slideId, err := uc.slideRepo.Create(ctx, slide)
	if err != nil {
		return entity.SlideDto{}, err
	}
	dto := entity.SlideDto{
		Id:      slideId,
		Name:    slide.Name,
		NameEng: slide.NameEng}
	return dto, nil
}

func (uc SlideUseCase) GetSlides(ctx context.Context, slide entity.SlidesRequest) (entity.ModuleSlides, error) {
	moduleId, err := uc.moduleRepo.GetModuleId(ctx, slide.CourseId, slide.ModuleNameEng)
	if err != nil {
		return entity.ModuleSlides{}, err
	}
	return uc.slideRepo.Get(ctx, moduleId)
}

func (uc SlideUseCase) GetSlideSections(ctx context.Context, slide entity.SlideSectionsRequest) (entity.SlideSections, error) {
	sections, err := uc.slideRepo.GetSections(ctx, slide.SlideId)
	if err != nil {
		return entity.SlideSections{}, err
	}
	return sections, nil
}
