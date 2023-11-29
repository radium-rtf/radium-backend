package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type AuthorUseCase struct {
	course postgres.Course
}

func NewAuthorUseCase(courseRepo postgres.Course) AuthorUseCase {
	return AuthorUseCase{course: courseRepo}
}

func (uc AuthorUseCase) GetCoursesByAuthorOrCoauthorId(ctx context.Context, id uuid.UUID) ([]*entity.Course, error) {
	return uc.course.GetByAuthorOrCoauthorId(ctx, id)
}
