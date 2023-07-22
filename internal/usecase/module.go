package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type ModuleUseCase struct {
	moduleRepo postgres.Module
}

func NewModuleUseCase(moduleRepo postgres.Module) ModuleUseCase {
	return ModuleUseCase{moduleRepo: moduleRepo}
}

func (uc ModuleUseCase) Create(ctx context.Context, moduleRequest *entity.Module) (*entity.Module, error) {
	return uc.moduleRepo.Create(ctx, moduleRequest)
}

func (uc ModuleUseCase) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	return uc.moduleRepo.Delete(ctx, id, isSoft)
}

// func (uc ModuleUseCase) GetCourseModules(ctx context.Context, courseId int) (entity.CourseModules, error) {
// 	return uc.moduleRepo.GetModules(ctx, courseId)
// }
