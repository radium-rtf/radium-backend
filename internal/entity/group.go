package entity

import "github.com/google/uuid"

type (
	GroupPost struct {
		Name        string      `json:"name"`
		CoursesIds  []uuid.UUID `json:"coursesIds,omitempty"`
		StudentsIds []uuid.UUID `json:"studentsIds,omitempty"`
	}

	Group struct {
		DBModel
		Name       string    `gorm:"not null"`
		InviteCode string    `gorm:"not null"`
		Courses    []*Course `gorm:"many2many:group_course"`
		Students   []*User   `gorm:"many2many:group_student"`
	}

	GroupJoin struct {
		UserId     uuid.UUID
		InviteCode string
	}

	GroupDto struct {
		Id         uuid.UUID    `json:"id"`
		Name       string       `json:"name"`
		InviteCode string       `json:"inviteCode"`
		Courses    []*CourseDto `json:"courses"`
		Students   []*UserDto   `json:"students"`
	}
)
