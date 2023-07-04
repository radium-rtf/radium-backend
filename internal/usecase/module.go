package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type ModuleUseCase struct {
	moduleRepo repo.ModuleRepo
}

func NewModuleUseCase(pg *db.Query) ModuleUseCase {
	return ModuleUseCase{moduleRepo: repo.NewModuleRepo(pg)}
}

func (uc ModuleUseCase) CreateModule(ctx context.Context, moduleRequest *entity.Module) (*entity.Module, error) {
	return uc.moduleRepo.Create(ctx, moduleRequest)
}

func (uc ModuleUseCase) Delete(ctx context.Context, destroy *entity.Destroy) error {
	return uc.moduleRepo.Delete(ctx, destroy)
}

// func (uc ModuleUseCase) GetCourseModules(ctx context.Context, courseId int) (entity.CourseModules, error) {
// 	return uc.moduleRepo.GetModules(ctx, courseId)
// }
