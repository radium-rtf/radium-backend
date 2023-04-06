package entity

import "errors"

var (
	GroupNotFoundErr        = errors.New("группа не найдена")
	GroupStudentNotFoundErr = errors.New("студент не найден в группе")
	GroupTeacherNotFoundErr = errors.New("данный пользователь не я вляется преподавателем в группе")
)

type GroupName struct {
	Name string
}

type GroupJoin struct {
	UserId, GroupId string
}

type Group struct {
	Id, Name string
}

type GroupTeacher struct {
	Id, UserId, GroupId string
}

type GroupStudent struct {
	Id, GroupId, UserId string
}

type GroupStudents struct {
	Group    Group
	Students []UserDto
}
