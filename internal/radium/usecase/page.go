package usecase

import (
	"context"
	"errors"
	"slices"

	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"

	"github.com/google/uuid"
)

type PageUseCase struct {
	page   postgres2.Page
	module postgres2.Module
	user   postgres2.User
}

func NewPageUseCase(pageRepo postgres2.Page, module postgres2.Module, user postgres2.User) PageUseCase {
	return PageUseCase{page: pageRepo, module: module, user: user}
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
	if err != nil && !errors.Is(err, repoerr.NotFound) {
		return nil, err
	}

	page.Order = 1

	if !errors.Is(err, repoerr.NotFound) {
		page.Order = last.Order + 1
	}

	return uc.page.Create(ctx, page)
}

func (uc PageUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	return uc.page.GetById(ctx, id)
}

func (uc PageUseCase) GetByIdWithUserAnswers(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity.Page, error) {
	page, err := uc.page.GetByIdWithUserAnswers(ctx, id, userId)
	if err != nil {
		return nil, err
	}

	_ = uc.user.SaveLastVisitedPage(ctx, page, userId)
	return page, nil
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

func (uc PageUseCase) UpdateOrder(ctx context.Context, id, editorId uuid.UUID, order uint) (*entity.Page, error) {
	canEditErr := uc.canEdit(ctx, id, editorId)
	if canEditErr != nil {
		return nil, canEditErr
	}
	page, err := uc.page.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	last, err := uc.page.GetLastPage(ctx, page.ModuleId)
	if float64(order) > last.Order {
		return nil, errors.New("новое местоположение не может быть больше местоположения последнего элемента")
	}
	return uc.page.UpdateOrder(ctx, page, order)
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

func (uc PageUseCase) GetBySlug(ctx context.Context, slug string) (*entity.Page, error) {
	return uc.page.GetBySlug(ctx, slug)
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
