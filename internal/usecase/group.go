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
	group   postgres.Group
	teacher postgres.Teacher
	course  postgres.Course
	answer  postgres.Answer
}

func NewGroupUseCase(groupRepo postgres.Group, course postgres.Course, answer postgres.Answer,
	teacher postgres.Teacher) GroupUseCase {
	return GroupUseCase{group: groupRepo, course: course, answer: answer, teacher: teacher}
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

func (uc GroupUseCase) GetWithAnswers(ctx context.Context, teacherId, groupId, courseId uuid.UUID) (*entity.Group, error) {
	courses, err := uc.teacher.GetCoursesByTeacherId(ctx, teacherId)
	if err != nil {
		return nil, err
	}

	contains := slices.ContainsFunc(courses, func(group *entity.TeacherCourseGroup) bool {
		return group.GroupId == groupId && group.CourseId == courseId
	})
	if !contains {
		return nil, errors.New("преподаватель может получать ответы только своих студентов")
	}

	return uc.group.GetWithAnswers(ctx, groupId, courseId)
}

func (uc GroupUseCase) GetReportByCourse(ctx context.Context, userId, courseId, groupId uuid.UUID) (
	*model.Report, error) {
	group, err := uc.group.GetById(ctx, groupId)
	if err != nil {
		return nil, err
	}

	teacherCourses, err := uc.teacher.GetCoursesByTeacherId(ctx, userId)
	if err != nil {
		return nil, err
	}

	contains := slices.ContainsFunc(group.Courses, func(course *entity.Course) bool {
		return course.Id == courseId
	})
	if !contains {
		return nil, errors.New("такой курс не назначен группе")
	}

	contains = slices.ContainsFunc(teacherCourses, func(course *entity.TeacherCourseGroup) bool {
		return course.CourseId == courseId && course.GroupId == groupId
	})
	if !contains {
		return nil, errors.New("только преподаватели имеют доступ к ведомости")
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
