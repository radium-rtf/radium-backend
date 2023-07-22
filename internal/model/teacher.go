package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	TeacherCourse struct {
		Course *Course  `json:"course"`
		Groups []*Group `json:"groups"`
	}
)

func NewTeacherCourses(teacher *entity.Teacher) []*TeacherCourse {
	courseGroups := make(map[uuid.UUID][]*Group)
	courseById := make(map[uuid.UUID]*Course)

	for _, course := range teacher.Courses {
		if _, ok := courseGroups[course.CourseId]; !ok {
			courseGroups[course.CourseId] = make([]*Group, 0)
		}
		if _, ok := courseById[course.CourseId]; !ok {
			courseById[course.CourseId] = NewCourse(course.Course)
		}

		group := NewGroup(course.Group)
		courseGroups[course.CourseId] = append(courseGroups[course.CourseId], group)
	}

	courses := make([]*TeacherCourse, 0, len(courseGroups))

	for courseId, groups := range courseGroups {
		courseDto := courseById[courseId]
		teacherCourse := &TeacherCourse{Course: courseDto}

		for _, group := range groups {
			teacherCourse.Groups = append(teacherCourse.Groups, group)
		}
		courses = append(courses, teacherCourse)
	}

	return courses
}
