package entity

import "github.com/google/uuid"

type (
	Teacher struct {
		DBModel
		UserId uuid.UUID
		Group  []Group `gorm:"many2many:group_teacher"`
	}

	TeacherCourseDto struct {
		DBModel
		Course Course
		Groups []Group
	}

	TeacherDto struct {
		DBModel
		UserId uuid.UUID
		Course []TeacherCourseDto
	}
)
