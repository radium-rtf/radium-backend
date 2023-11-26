package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"
	"slices"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type PageUseCase struct {
	page   postgres.Page
	module postgres.Module
}

func NewPageUseCase(pageRepo postgres.Page, module postgres.Module) PageUseCase {
	return PageUseCase{page: pageRepo, module: module}
}

func (uc PageUseCase) Create(ctx context.Context, page *entity.Page, editorId uuid.UUID) (*entity.Page, error) {
	course, err := uc.module.GetCourseByModuleId(ctx, page.ModuleId)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}

	last, err := uc.page.GetLastPage(ctx, page.ModuleId)
	if err != nil && !errors.Is(err, repoerr.PageNotFound) {
		return nil, err
	}

	page.Order = 1
	if !errors.Is(err, repoerr.PageNotFound) {
		page.Order = last.Order + 1
	}

	return uc.page.Create(ctx, page)
}

func (uc PageUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	return uc.page.GetById(ctx, id)
}

func (uc PageUseCase) Delete(ctx context.Context, id, editorId uuid.UUID, isSoft bool) error {
	canEditErr := uc.canEdit(ctx, id, editorId)
	if canEditErr != nil {
		return canEditErr
	}
	return uc.page.Delete(ctx, id, isSoft)
}

func (uc PageUseCase) Update(ctx context.Context, page *entity.Page, editorId uuid.UUID) (*entity.Page, error) {
	canEditErr := uc.canEdit(ctx, page.Id, editorId)
	if canEditErr != nil {
		return nil, canEditErr
	}
	return uc.page.Update(ctx, page)
}

func (p PageUseCase) UpdateOrder(ctx context.Context, id, editorId uuid.UUID, order uint) (*entity.Page, error) {
	//TODO implement me
	panic("implement me")
}

func (uc PageUseCase) GetNextAndPrevious(ctx context.Context, page *entity.Page) (*model.NextAndPreviousPage, error) {
	modules, err := uc.page.GetModulesByPageId(ctx, page.Id)
	if err != nil {
		return nil, err
	}
	moduleIndex := slices.IndexFunc(modules, func(module *entity.Module) bool {
		return module.Id == page.ModuleId
	})
	pageIndex := slices.IndexFunc(modules[moduleIndex].Pages, func(p *entity.Page) bool {
		return p.Id == page.Id
	})

	var nextAndPrevious = model.GetNextAndPreviousPage(moduleIndex, pageIndex, modules)
	return nextAndPrevious, nil
}

func (uc PageUseCase) canEdit(ctx context.Context, id, editorId uuid.UUID) error {
	course, err := uc.page.GetCourseByPageId(ctx, id)
	if err != nil {
		return err
	}

	if !course.CanEdit(editorId) {
		return cantEditCourse
	}
	return nil
}
