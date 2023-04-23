package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
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
		image.Location, courseRequest.AuthorId, courseRequest.Type)
	_, err = uc.courseRepo.GetByName(ctx, course.Name)
	if err == nil {
		return course, errors.New("курс с таким названием уже существует")
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

func (uc CourseUseCase) CreateLink(ctx context.Context, link entity.LinkRequest) (entity.LinkDto, error) {
	courseLink := entity.Link{
		Id:       uuid.NewString(),
		Name:     link.Name,
		Link:     link.Link,
		CourseId: link.CourseId,
	}
	dto := entity.LinkDto{Link: link.Link, Name: link.Name}
	return dto, uc.courseRepo.CreateLink(ctx, courseLink)
}

func (uc CourseUseCase) CreateCollaborator(ctx context.Context, collaborator entity.Collaborator) (entity.Collaborator, error) {
	courseCollaborator := entity.CourseCollaborator{
		CourseId:  collaborator.CourseId,
		UserEmail: collaborator.UserEmail,
		Id:        uuid.NewString(),
	}
	return collaborator, uc.courseRepo.CreateCollaborator(ctx, courseCollaborator)
}

func (uc CourseUseCase) GetCourseById(ctx context.Context, id int) (entity.CourseTitle, error) {
	return uc.courseRepo.GetFullById(ctx, id)
}

func (uc CourseUseCase) Join(ctx context.Context, userId string, courseId string) (entity.Course, error) {
	err := uc.courseRepo.Join(ctx, userId, courseId)
	if err != nil {
		return entity.Course{}, err
	}
	return uc.courseRepo.GetById(ctx, courseId)
}
