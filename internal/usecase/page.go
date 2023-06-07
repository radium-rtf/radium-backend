package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type PageUseCase struct {
	pageRepo   repo.PageRepo
	moduleRepo repo.ModuleRepo
}

func NewPageUseCase(pg *db.Query) PageUseCase {
	return PageUseCase{pageRepo: repo.NewPageRepo(pg), moduleRepo: repo.NewModuleRepo(pg)}
}

func (uc PageUseCase) CreatePage(ctx context.Context, page entity.PageRequest) (*entity.Page, error) {
	return uc.pageRepo.Create(ctx, page)
}

func (uc PageUseCase) GetByID(ctx context.Context, id string) (*entity.Page, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return uc.pageRepo.GetByID(ctx, uid)
}

// func (uc SlideUseCase) GetSlides(ctx context.Context, slide entity.SlidesRequest) (entity.ModuleSlides, error) {
// 	moduleId, err := uc.moduleRepo.GetModuleId(ctx, slide.CourseId, slide.ModuleNameEng)
// 	if err != nil {
// 		return entity.ModuleSlides{}, err
// 	}
// 	return uc.slideRepo.Get(ctx, moduleId)
// }

// func (uc SlideUseCase) GetSlideSections(ctx context.Context, slide entity.SlideSectionsRequest) (entity.SlideSections, error) {
// 	sections, err := uc.slideRepo.GetSections(ctx, slide.SlideId)
// 	if err != nil {
// 		return entity.SlideSections{}, err
// 	}
// 	return sections, nil
// }
