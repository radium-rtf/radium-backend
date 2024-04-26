package model

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	TeacherCourse struct {
		Course *Course  `json:"course"`
		Groups []*Group `json:"groups"`
	}
)

func NewTeacherCourses(teacher []*entity2.Teacher) []*TeacherCourse {
	courseGroups := make(map[uuid.UUID][]*Group)
	courseById := make(map[uuid.UUID]*Course)

	for _, course := range teacher {
		if _, ok := courseGroups[course.CourseId]; !ok {
			courseGroups[course.CourseId] = make([]*Group, 0)
		}
		if _, ok := courseById[course.CourseId]; !ok {
			courseById[course.CourseId] = NewCourse(course.Course, map[uuid.UUID][]*entity2.Answer{}, uuid.UUID{})
		}

		group := NewGroup(course.Group)
		courseGroups[course.CourseId] = append(courseGroups[course.CourseId], group)
	}

	courses := make([]*TeacherCourse, 0, len(courseGroups))

	for courseId, groups := range courseGroups {
		courseDto := courseById[courseId]
		teacherCourse := &TeacherCourse{Course: courseDto}

		teacherCourse.Groups = append(teacherCourse.Groups, groups...)
		courses = append(courses, teacherCourse)
	}

	return courses
}
