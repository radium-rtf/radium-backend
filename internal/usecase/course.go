package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"slices"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type CourseUseCase struct {
	courseRepo postgres.Course
	user       postgres.User
}

func NewCourseUseCase(courseRepo postgres.Course, user postgres.User) CourseUseCase {
	return CourseUseCase{courseRepo: courseRepo, user: user}
}

func (uc CourseUseCase) Create(ctx context.Context, course *entity.Course, creatorId uuid.UUID) (*entity.Course, error) {
	creator, err := uc.user.GetById(ctx, creatorId)
	if err != nil {
		return nil, err
	}

	if creator.Roles.IsAuthor {
		return nil, errors.New("только автор может созавать курсы")
	}

	return uc.courseRepo.Create(ctx, course)
}

func (uc CourseUseCase) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	return uc.courseRepo.Get(ctx)
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
	return uc.courseRepo.GetFullById(ctx, courseId)
}

func (uc CourseUseCase) Delete(ctx context.Context, id, deleterId uuid.UUID, isSoft bool) error {
	course, err := uc.courseRepo.GetFullById(ctx, id)
	if err != nil {
		return err
	}
	canDelete := slices.ContainsFunc(course.Authors, func(user *entity.User) bool {
		return user.Id == deleterId
	})
	if !canDelete {
		return errors.New("только автор курса может его удалить")
	}
	return uc.courseRepo.Delete(ctx, id, isSoft)
}

func (uc CourseUseCase) Update(ctx context.Context, update *entity.Course, editorId uuid.UUID) (*entity.Course, error) {
	course, err := uc.courseRepo.GetFullById(ctx, update.Id)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}

	return uc.courseRepo.Update(ctx, update)
}

func (uc CourseUseCase) Publish(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity.Course, error) {
	course, err := uc.courseRepo.GetFullById(ctx, id)
	if err != nil {
		return nil, err
	}
	if !course.IsPublished && validator.Struct(course) != nil {
		return nil, errors.New("курс ещё не готов к публикации")
	}
	isAuthor := slices.ContainsFunc(course.Authors, func(user *entity.User) bool {
		return user.Id == userId
	})
	if !isAuthor {
		return nil, errors.New("пубилковать и снимать с публикации курс могут только авторы")
	}

	course.IsPublished = !course.IsPublished
	course, err = uc.courseRepo.Update(ctx, course)

	return course, err
}

func (c CourseUseCase) DeleteLink(ctx context.Context, id, editorId uuid.UUID) error {
	link, err := c.courseRepo.GetLinkById(ctx, id)
	if err != nil {
		return err
	}

	course := link.Course
	if !course.CanEdit(editorId) {
		return cantEditCourse
	}
	return c.courseRepo.DeleteLink(ctx, id)
}

func (c CourseUseCase) CreateLink(ctx context.Context, link *entity.Link, editorId uuid.UUID) (*entity.Link, error) {
	course, err := c.courseRepo.GetFullById(ctx, link.CourseId)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}
	return link, c.courseRepo.CreateLink(ctx, link)
}
