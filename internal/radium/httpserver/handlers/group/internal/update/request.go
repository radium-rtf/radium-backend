package update

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type Group struct {
	Name        string      `json:"name" validate:"required,min=1,max=40"`
	CoursesIds  []uuid.UUID `json:"coursesIds"`
	StudentsIds []uuid.UUID `json:"studentsIds"`
}

func (r Group) toGroup(groupId uuid.UUID) *entity2.Group {
	courses := make([]*entity2.Course, 0, len(r.StudentsIds))
	for _, id := range r.CoursesIds {
		courses = append(courses, &entity2.Course{DBModel: entity2.DBModel{Id: id}})
	}

	students := make([]*entity2.User, 0, len(r.StudentsIds))
	for _, id := range r.StudentsIds {
		students = append(students, &entity2.User{DBModel: entity2.DBModel{Id: id}})
	}

	return &entity2.Group{
		DBModel:  entity2.DBModel{Id: groupId},
		Name:     r.Name,
		Students: students,
		Courses:  courses,
	}
}
