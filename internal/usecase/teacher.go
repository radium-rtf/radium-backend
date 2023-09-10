package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type TeacherUseCase struct {
	teacher postgres.Teacher
}

func NewTeacherUseCase(teacherRepo postgres.Teacher) TeacherUseCase {
	return TeacherUseCase{teacher: teacherRepo}
}

func (uc TeacherUseCase) GetByUserId(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	return uc.teacher.GetByUserId(ctx, id)
}

func (uc TeacherUseCase) Create(ctx context.Context, teacher *entity.Teacher) (*entity.Teacher, error) {
	return uc.teacher.Create(ctx, teacher)
}
