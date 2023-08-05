package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	Teacher struct {
		TeacherId uuid.UUID        `json:"teacherId"`
		Courses   []*TeacherCourse `json:"courses"`
	}

	TeacherCourse struct {
		CourseId uuid.UUID `json:"courseId"`
		GroupId  uuid.UUID `json:"groupId"`
	}
)

func (r Teacher) toTeacher() *entity.Teacher {
	teacher := &entity.Teacher{UserId: r.TeacherId, DBModel: entity.DBModel{Id: uuid.New()}}
	courses := make([]*entity.TeacherCourse, 0, len(r.Courses))
	for _, coursePost := range r.Courses {
		course := &entity.TeacherCourse{CourseId: coursePost.CourseId, GroupId: coursePost.GroupId}
		courses = append(courses, course)
	}
	teacher.Courses = courses
	return teacher
}
