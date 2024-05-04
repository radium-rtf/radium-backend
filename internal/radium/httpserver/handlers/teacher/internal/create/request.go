package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	Teacher struct {
		UserId  uuid.UUID        `json:"userId"`
		Courses []*TeacherCourse `json:"courses"`
	}

	TeacherCourse struct {
		CourseId uuid.UUID `json:"courseId"`
		GroupId  uuid.UUID `json:"groupId"`
	}
)

func (r Teacher) toCourses() []*entity.Teacher {
	userId := r.UserId

	courses := make([]*entity.Teacher, 0, len(r.Courses))
	for _, coursePost := range r.Courses {
		course := &entity.Teacher{
			UserId:   userId,
			CourseId: coursePost.CourseId,
			GroupId:  coursePost.GroupId,
		}
		courses = append(courses, course)
	}
	return courses
}
