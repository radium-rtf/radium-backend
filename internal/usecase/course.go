package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type CourseUseCase struct {
	storage    filestorage.Storage
	courseRepo repo.CourseRepo
}

func NewCourseUseCase(pg *db.Query, storage filestorage.Storage) CourseUseCase {
	return CourseUseCase{storage: storage, courseRepo: repo.NewCourseRepo(pg)}
}

func (uc CourseUseCase) CreateCourse(ctx context.Context, courseRequest entity.CourseRequest) (*entity.Course, error) {
	course, err := uc.courseRepo.Create(ctx, courseRequest)
	if err != nil {
		return &entity.Course{}, err
	}
	return course, nil
}

func (uc CourseUseCase) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	return uc.courseRepo.GetCourses(ctx)
}

func (uc CourseUseCase) GetCourseById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	return uc.courseRepo.GetFullById(ctx, id)
}

func (uc CourseUseCase) GetCourseBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	return uc.courseRepo.GetFullBySlug(ctx, slug)
}

func (uc CourseUseCase) Join(ctx context.Context, userId uuid.UUID, courseId uuid.UUID) (*entity.Course, error) {
	err := uc.courseRepo.Join(ctx, userId, courseId)
	if err != nil {
		return &entity.Course{}, err
	}
	return uc.courseRepo.GetById(ctx, courseId)
}
