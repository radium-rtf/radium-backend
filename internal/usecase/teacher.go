package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type TeacherUseCase struct {
	teacherRepo repo.TeacherRepo
}

func NewTeacherUseCase(pg *db.Query) TeacherUseCase {
	return TeacherUseCase{teacherRepo: repo.NewTeacherRepo(pg)}
}

func (uc TeacherUseCase) GetByUserId(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	return uc.teacherRepo.GetByUserId(ctx, id)
}

func (uc TeacherUseCase) Create(ctx context.Context, teacher *entity.Teacher) (*entity.Teacher, error) {
	return uc.teacherRepo.Create(ctx, teacher)
}
