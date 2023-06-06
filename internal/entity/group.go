package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrGroupNotFound        = errors.New("группа не найдена")
	ErrGroupStudentNotFound = errors.New("студент не найден в группе")
	ErrGroupTeacherNotFound = errors.New("данный пользователь не является преподавателем в группе")
)

type (
	GroupName struct {
		Name string
	}

	GroupJoin struct {
		UserId, GroupId uuid.UUID
	}

	Group struct {
		Id       uuid.UUID `gorm:"default:gen_random_uuid()"`
		Name     string
		Students []User `gorm:"many2many:group_student"`
	}

	GroupTeacher struct {
		Id, UserId, GroupId string
	}

	GroupStudent struct {
		Id, GroupId, UserId string
	}

	GroupStudents struct {
		Group    Group
		Students []UserDto
	}
)
