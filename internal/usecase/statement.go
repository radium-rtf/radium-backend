package usecase

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type StatementUseCase struct {
	course CourseUseCase
	answer AnswerUseCase
	group  GroupUseCase
}

func NewStatementUseCase(course postgres.Course, answer postgres.Answer, group postgres.Group,
	section postgres.Section) StatementUseCase {
	return StatementUseCase{
		course: NewCourseUseCase(course),
		answer: NewAnswerUseCase(section, answer),
		group:  NewGroupUseCase(group),
	}
}

func (uc StatementUseCase) Get(groupId uuid.UUID, courseId uuid.UUID) {
	panic("")
}
