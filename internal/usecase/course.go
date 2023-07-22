package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type CourseUseCase struct {
	courseRepo postgres.Course
}

func NewCourseUseCase(courseRepo postgres.Course) CourseUseCase {
	return CourseUseCase{courseRepo: courseRepo}
}

func (uc CourseUseCase) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	course, err := uc.courseRepo.Create(ctx, course)
	if err != nil {
		return &entity.Course{}, err
	}
	return course, nil
}

func (uc CourseUseCase) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	return uc.courseRepo.GetCourses(ctx)
}

func (uc CourseUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	return uc.courseRepo.GetFullById(ctx, id)
}

func (uc CourseUseCase) GetBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	return uc.courseRepo.GetFullBySlug(ctx, slug)
}

func (uc CourseUseCase) Join(ctx context.Context, userId uuid.UUID, courseId uuid.UUID) (*entity.Course, error) {
	err := uc.courseRepo.Join(ctx, userId, courseId)
	if err != nil {
		return &entity.Course{}, err
	}
	return uc.courseRepo.GetById(ctx, courseId)
}

func (uc CourseUseCase) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	return uc.courseRepo.Delete(ctx, id, isSoft)
}
