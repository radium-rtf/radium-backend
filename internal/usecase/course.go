package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type CourseUseCase struct {
	storage    filestorage.Storage
	courseRepo repo.CourseRepo
}

func NewCourseUseCase(pg *postgres.Postgres, storage filestorage.Storage) CourseUseCase {
	return CourseUseCase{storage: storage, courseRepo: repo.NewCourseRepo(pg)}
}

func (uc CourseUseCase) CreateCourse(ctx context.Context, courseRequest entity.CourseRequest) (entity.Course, error) {
	var course entity.Course
	contentType := courseRequest.Header.Header["Content-Type"][0]
	image, err := uc.storage.PutImage(ctx, courseRequest.Logo, courseRequest.Header.Size, contentType)
	if err != nil {
		return course, err
	}
	course = entity.NewCourse(0, courseRequest.Name, courseRequest.Description,
		image.Location, courseRequest.Chat, courseRequest.Type)
	_, err = uc.courseRepo.GetByName(ctx, course.Name)
	if err == nil {
		return course, errors.New("Курс с таким названием уже существует")
	}
	if err != entity.CourseNotFoundErr {
		return course, err
	}
	err = uc.courseRepo.Create(ctx, course)
	if err != nil {
		return course, err
	}
	return uc.courseRepo.GetByName(ctx, course.Name)
}

func (uc CourseUseCase) GetCourses(ctx context.Context) ([]entity.Course, error) {
	return uc.courseRepo.GetCourses(ctx)
}
