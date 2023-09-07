package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"slices"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type GroupUseCase struct {
	group  postgres.Group
	course postgres.Course
	answer postgres.Answer
}

func NewGroupUseCase(groupRepo postgres.Group, course postgres.Course, answer postgres.Answer) GroupUseCase {
	return GroupUseCase{group: groupRepo, course: course, answer: answer}
}

func (uc GroupUseCase) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	return uc.group.Create(ctx, group)
}

func (uc GroupUseCase) Join(ctx context.Context, studentId uuid.UUID, code string) error {
	return uc.group.JoinStudent(ctx, studentId, code)
}

func (uc GroupUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	return uc.group.GetById(ctx, id)
}

func (uc GroupUseCase) Get(ctx context.Context) ([]*entity.Group, error) {
	return uc.group.Get(ctx)
}

func (uc GroupUseCase) GetReportByCourse(ctx context.Context, userId, courseId, groupId uuid.UUID) (
	*model.Report, error) {
	group, err := uc.group.GetById(ctx, groupId)
	if err != nil {
		return nil, err
	}

	contains := slices.ContainsFunc(group.Students, func(user *entity.User) bool {
		return userId == user.Id
	})
	if !contains {
		return nil, errors.New("у вас нет доступа к просмотру этой ведомости")
	}

	contains = slices.ContainsFunc(group.Courses, func(course *entity.Course) bool {
		return course.Id == courseId
	})
	if !contains {
		return nil, errors.New("такой курс не назначен группе")
	}

	return uc.newReport(ctx, courseId, group)
}

func (uc GroupUseCase) newReport(ctx context.Context, courseId uuid.UUID, group *entity.Group) (
	*model.Report, error) {
	course, err := uc.course.GetFullById(ctx, courseId)
	if err != nil {
		return nil, err
	}

	sectionsIds := course.SectionsIds()
	usersIds := make([]uuid.UUID, 0, len(group.Students))
	for _, user := range group.Students {
		usersIds = append(usersIds, user.Id)
	}

	collection, err := uc.answer.GetByUsers(ctx, usersIds, sectionsIds)
	if err != nil {
		return nil, err
	}

	report := model.NewGroupReport(collection, course)
	return report, nil
}
