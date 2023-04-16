package entity

import "errors"

var (
	GroupNotFoundErr        = errors.New("группа не найдена")
	GroupStudentNotFoundErr = errors.New("студент не найден в группе")
	GroupTeacherNotFoundErr = errors.New("данный пользователь не я вляется преподавателем в группе")
)

type (
	GroupName struct {
		Name string
	}

	GroupJoin struct {
		UserId, GroupId string
	}

	Group struct {
		Id, Name string
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
