package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/radium-rtf/radium-backend/pkg/translit"
)

type ModuleUseCase struct {
	moduleRepo repo.ModuleRepo
}

func NewModuleUseCase(pg *postgres.Postgres) ModuleUseCase {
	return ModuleUseCase{moduleRepo: repo.NewModuleRepo(pg)}
}

func (uc ModuleUseCase) CreateModule(ctx context.Context, moduleRequest entity.ModuleRequest) (
	entity.ModuleDto, error) {
	module := entity.Module{
		Name:     moduleRequest.Name,
		NameEng:  translit.RuEn(moduleRequest.Name),
		CourseId: moduleRequest.CourseId,
	}
	moduleDto := entity.ModuleDto{
		NameEng: module.NameEng,
		Name:    module.Name,
	}
	return moduleDto, uc.moduleRepo.Create(ctx, module)
}

func (uc ModuleUseCase) GetCourseModules(ctx context.Context, courseId int) (entity.CourseModules, error) {
	return uc.moduleRepo.GetModules(ctx, courseId)
}
