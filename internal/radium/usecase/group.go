package usecase

import (
	"context"
	"errors"
	"slices"

	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type GroupUseCase struct {
	group   postgres2.Group
	teacher postgres2.Teacher
	course  postgres2.Course
	answer  postgres2.Answer
}

func NewGroupUseCase(groupRepo postgres2.Group, course postgres2.Course, answer postgres2.Answer,
	teacher postgres2.Teacher) GroupUseCase {
	return GroupUseCase{group: groupRepo, course: course, answer: answer, teacher: teacher}
}

func (uc GroupUseCase) UpdateGroup(ctx context.Context, group *entity2.Group) (*entity2.Group, error) {
	return uc.group.Update(ctx, group)
}

func (uc GroupUseCase) Create(ctx context.Context, group *entity2.Group) (*entity2.Group, error) {
	return uc.group.Create(ctx, group)
}

func (uc GroupUseCase) Join(ctx context.Context, studentId uuid.UUID, code string) error {
	return uc.group.JoinStudent(ctx, studentId, code)
}

func (uc GroupUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity2.Group, error) {
	return uc.group.GetById(ctx, id)
}

func (uc GroupUseCase) AddCourse(ctx context.Context, groupId uuid.UUID, courseId uuid.UUID) (*entity2.Group, error) {
	return uc.group.AddCourse(ctx, groupId, courseId)
}

func (uc GroupUseCase) Get(ctx context.Context) ([]*entity2.Group, error) {
	return uc.group.Get(ctx)
}

func (uc GroupUseCase) GetWithAnswers(ctx context.Context, teacherId, groupId, courseId uuid.UUID) (*entity2.Group, error) {
	courses, err := uc.teacher.GetCoursesByTeacherId(ctx, teacherId)
	if err != nil {
		return nil, err
	}

	contains := slices.ContainsFunc(courses, func(group *entity2.Teacher) bool {
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

	containsCourse := slices.ContainsFunc(group.Courses, func(course *entity2.Course) bool {
		return course.Id == courseId
	})
	if !containsCourse {
		return nil, errors.New("такой курс не назначен группе")
	}

	isTeacher := slices.ContainsFunc(teacherCourses, func(course *entity2.Teacher) bool {
		return course.CourseId == courseId && course.GroupId == groupId
	})
	canViewReport := isTeacher || slices.ContainsFunc(group.Students, func(user *entity2.User) bool {
		return user.Id == userId
	})

	if !canViewReport {
		return nil, errors.New("только преподаватели и студенты имеют доступ к ведомости")
	}

	return uc.newReport(ctx, courseId, group)
}

func (uc GroupUseCase) newReport(ctx context.Context, courseId uuid.UUID, group *entity2.Group) (
	*model.Report, error) {
	course, err := uc.course.GetFullWithStudents(ctx, courseId)
	if err != nil {
		return nil, err
	}

	// todo: убрать !!!!
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
