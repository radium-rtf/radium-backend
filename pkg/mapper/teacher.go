package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Teacher struct {
	g Group
	c Course
}

func (t Teacher) ToCourses(teacher *entity.Teacher) []*entity.TeacherCourseDto { // TODO: мб можно сразу получить
	courseGroups := make(map[uuid.UUID][]*entity.GroupDto)
	courseById := make(map[uuid.UUID]*entity.CourseDto)

	for _, course := range teacher.Courses {
		if _, ok := courseGroups[course.CourseId]; !ok {
			courseGroups[course.CourseId] = make([]*entity.GroupDto, 0)
		}
		if _, ok := courseById[course.CourseId]; !ok {
			courseById[course.CourseId] = t.c.ToDto(course.Course)
		}

		group := t.g.ToDto(course.Group)
		courseGroups[course.CourseId] = append(courseGroups[course.CourseId], group)
	}

	courses := make([]*entity.TeacherCourseDto, 0, len(courseGroups))

	for courseId, groups := range courseGroups {
		courseDto := courseById[courseId]
		teacherCourse := &entity.TeacherCourseDto{Course: courseDto}

		for _, group := range groups {
			teacherCourse.Groups = append(teacherCourse.Groups, group)
		}
		courses = append(courses, teacherCourse)
	}

	return courses
}

func (t Teacher) PostToTeacher(post entity.TeacherPost) *entity.Teacher {
	teacher := &entity.Teacher{UserId: post.TeacherId, DBModel: entity.DBModel{Id: uuid.New()}}
	courses := make([]*entity.TeacherCourse, 0, len(post.Courses))
	for _, coursePost := range post.Courses {
		course := &entity.TeacherCourse{CourseId: coursePost.CourseId, GroupId: coursePost.GroupId}
		courses = append(courses, course)
	}
	teacher.Courses = courses
	return teacher
}
