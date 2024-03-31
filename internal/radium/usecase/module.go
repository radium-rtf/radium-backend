package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
)

type ModuleUseCase struct {
	module postgres2.Module
	course postgres2.Course
}

func NewModuleUseCase(moduleRepo postgres2.Module, courseRepo postgres2.Course) ModuleUseCase {
	return ModuleUseCase{module: moduleRepo, course: courseRepo}
}

func (uc ModuleUseCase) Create(ctx context.Context, module *entity.Module, editorId uuid.UUID) (*entity.Module, error) {
	course, err := uc.course.GetFullById(ctx, module.CourseId)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}

	last, err := uc.module.GetLastModule(ctx, module.CourseId)
	if err != nil && !errors.Is(err, repoerr.NotFound) {
		return nil, err
	}

	module.Order = 1
	if !errors.Is(err, repoerr.NotFound) {
		module.Order = last.Order + 1
	}

	return uc.module.Create(ctx, module)
}

func (uc ModuleUseCase) Delete(ctx context.Context, id, editorId uuid.UUID, isSoft bool) error {
	canEditErr := uc.canEdit(ctx, id, editorId)
	if canEditErr != nil {
		return canEditErr
	}
	return uc.module.Delete(ctx, id, isSoft)
}

func (uc ModuleUseCase) Update(ctx context.Context, module *entity.Module, editorId uuid.UUID) (*entity.Module, error) {
	canEditErr := uc.canEdit(ctx, module.Id, editorId)
	if canEditErr != nil {
		return nil, canEditErr
	}
	return uc.module.Update(ctx, module)
}

func (uc ModuleUseCase) UpdateOrder(ctx context.Context, id, editorId uuid.UUID, order uint) (*entity.Module, error) {
	canEditErr := uc.canEdit(ctx, id, editorId)
	if canEditErr != nil {
		return nil, canEditErr
	}
	module, err := uc.module.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	last, err := uc.module.GetLastModule(ctx, module.CourseId)
	if float64(order) > last.Order {
		return nil, errors.New("новое местоположение не может быть больше местоположения последнего элемента")
	}
	return uc.module.UpdateOrder(ctx, module, order)
}
func (uc ModuleUseCase) canEdit(ctx context.Context, id, editorId uuid.UUID) error {
	course, err := uc.module.GetCourseByModuleId(ctx, id)
	if err != nil {
		return err
	}

	if !course.CanEdit(editorId) {
		return cantEditCourse
	}
	return nil
}

// func (uc ModuleUseCase) GetCourseModules(ctx context.Context, courseId int) (entity.CourseModules, error) {
// 	return uc.module.GetModules(ctx, courseId)
// }
