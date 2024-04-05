package usecase

import (
	"context"
	"errors"
	"slices"

	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/validator"

	"github.com/google/uuid"
)

type CourseUseCase struct {
	courseRepo postgres2.Course
	user       postgres2.User
}

func NewCourseUseCase(courseRepo postgres2.Course, user postgres2.User) CourseUseCase {
	return CourseUseCase{courseRepo: courseRepo, user: user}
}

func (uc CourseUseCase) Create(ctx context.Context, course *entity2.Course, creatorId uuid.UUID) (*entity2.Course, error) {
	creator, err := uc.user.GetById(ctx, creatorId)
	if err != nil {
		return nil, err
	}

	if !creator.Roles.IsAuthor {
		return nil, errors.New("только автор может созавать курсы")
	}

	return uc.courseRepo.Create(ctx, course)
}

func (uc CourseUseCase) GetCourses(ctx context.Context) ([]*entity2.Course, error) {
	return uc.courseRepo.Get(ctx)
}

func (uc CourseUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity2.Course, error) {
	return uc.courseRepo.GetFullById(ctx, id)
}

func (uc CourseUseCase) GetByIdAndUser(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity2.Course, error) {
	return uc.courseRepo.GetFullByIdAndUser(ctx, id, userId)
}

func (c CourseUseCase) GetBySlugAndUser(ctx context.Context, slug string, userId uuid.UUID) (*entity2.Course, error) {
	return c.courseRepo.GetFullBySlugAndUser(ctx, slug, userId)
}

func (uc CourseUseCase) GetBySlug(ctx context.Context, slug string) (*entity2.Course, error) {
	return uc.courseRepo.GetFullBySlug(ctx, slug)
}

func (uc CourseUseCase) Join(ctx context.Context, userId uuid.UUID, courseId uuid.UUID) (*entity2.Course, error) {
	err := uc.courseRepo.Join(ctx, userId, courseId)
	if err != nil {
		return &entity2.Course{}, err
	}
	return uc.courseRepo.GetFullById(ctx, courseId)
}

func (uc CourseUseCase) Delete(ctx context.Context, id, deleterId uuid.UUID, isSoft bool) error {
	course, err := uc.courseRepo.GetFullById(ctx, id)
	if err != nil {
		return err
	}
	canDelete := slices.ContainsFunc(course.Authors, func(user *entity2.User) bool {
		return user.Id == deleterId
	})
	if !canDelete {
		return errors.New("только автор курса может его удалить")
	}
	return uc.courseRepo.Delete(ctx, id, isSoft)
}

func (uc CourseUseCase) Update(ctx context.Context, update *entity2.Course, editorId uuid.UUID) (*entity2.Course, error) {
	course, err := uc.courseRepo.GetFullById(ctx, update.Id)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}

	return uc.courseRepo.Update(ctx, update)
}

func (uc CourseUseCase) Publish(ctx context.Context, id uuid.UUID, userId uuid.UUID) (*entity2.Course, error) {
	course, err := uc.courseRepo.GetFullById(ctx, id)
	if err != nil {
		return nil, err
	}
	if course.Access == "" && validator.Struct(course) != nil {
		return nil, errors.New("access является обязательным полем")
	}
	isAuthor := slices.ContainsFunc(course.Authors, func(user *entity2.User) bool {
		return user.Id == userId
	})
	if !isAuthor {
		return nil, errors.New("пубилковать и снимать с публикации курс могут только авторы")
	}

	return uc.courseRepo.UpdateAccess(ctx, course.Id, string(course.Access))
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

func (c CourseUseCase) CreateLink(ctx context.Context, link *entity2.Link, editorId uuid.UUID) (*entity2.Link, error) {
	course, err := c.courseRepo.GetFullById(ctx, link.CourseId)
	if err != nil {
		return nil, err
	}
	if !course.CanEdit(editorId) {
		return nil, cantEditCourse
	}
	return link, c.courseRepo.CreateLink(ctx, link)
}
