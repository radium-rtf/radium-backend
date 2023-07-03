package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/mapper"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type PageUseCase struct {
	pageRepo   repo.PageRepo
	moduleRepo repo.ModuleRepo
	answerRepo repo.AnswerRepo
	mapper     mapper.Page
}

func NewPageUseCase(pg *db.Query) PageUseCase {
	return PageUseCase{pageRepo: repo.NewPageRepo(pg), moduleRepo: repo.NewModuleRepo(pg), answerRepo: repo.NewAnswerRepo(pg)}
}

func (uc PageUseCase) CreatePage(ctx context.Context, page entity.PageRequest) (*entity.Page, error) {
	return uc.pageRepo.Create(ctx, page)
}

func (uc PageUseCase) GetByID(ctx context.Context, id uuid.UUID, userId *uuid.UUID) (*entity.PageDto, error) {
	page, err := uc.pageRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userId == nil {
		return uc.mapper.Page(page, map[uuid.UUID]*entity.Answer{}), err
	}
	sectionsIds := make([]uuid.UUID, 0, len(page.Sections))
	for _, section := range page.Sections {
		sectionsIds = append(sectionsIds, section.Id)
	}
	answers, err := uc.answerRepo.Get(ctx, *userId, sectionsIds)
	if err != nil {
		return nil, err
	}
	p := uc.mapper.Page(page, answers)
	return p, err
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
