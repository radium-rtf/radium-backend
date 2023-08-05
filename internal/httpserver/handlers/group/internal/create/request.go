package create

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/otp"
)

type Group struct {
	Name        string      `json:"name"`
	CoursesIds  []uuid.UUID `json:"coursesIds"`
	StudentsIds []uuid.UUID `json:"studentsIds"`
}

func (r Group) toGroup() *entity.Group {
	courses := make([]*entity.Course, 0, len(r.StudentsIds))
	for _, id := range r.StudentsIds {
		courses = append(courses, &entity.Course{DBModel: entity.DBModel{Id: id}})
	}

	students := make([]*entity.User, 0, len(r.StudentsIds))
	for _, id := range r.StudentsIds {
		students = append(students, &entity.User{DBModel: entity.DBModel{Id: id}})
	}

	return &entity.Group{
		DBModel:    entity.DBModel{Id: uuid.New()},
		Name:       r.Name,
		InviteCode: otp.NewOTPGenerator().RandomSecret(10),
		Students:   students,
		Courses:    courses,
	}
}
