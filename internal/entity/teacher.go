package entity

import "github.com/google/uuid"

type (
	Teacher struct {
		DBModel
		UserId  uuid.UUID `gorm:"uniqueIndex; type:uuid"`
		Courses []*TeacherCourse
	}

	TeacherCourse struct {
		TeacherId uuid.UUID `gorm:"type:uuid; primaryKey"`
		CourseId  uuid.UUID `gorm:"type:uuid; primaryKey"`
		GroupId   uuid.UUID `gorm:"type:uuid; primaryKey"`
		Course    *Course
		Group     *Group
	}
)
