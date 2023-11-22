package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type ModuleUseCase struct {
	module postgres.Module
	course postgres.Course
}

func NewModuleUseCase(moduleRepo postgres.Module, courseRepo postgres.Course) ModuleUseCase {
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
	if err != nil && !errors.Is(err, repoerr.ModuleNotFound) {
		return nil, err
	}

	module.Order = 1
	if !errors.Is(err, repoerr.ModuleNotFound) {
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
	//TODO implement me
	panic("implement me")
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
