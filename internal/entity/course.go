package entity

import (
	"errors"
	"mime/multipart"
)

var (
	CourseNotFoundErr = errors.New("курс не найден")
)

type (
	CourseRequest struct {
		Name        string
		Description string
		AuthorId    string
		Logo        multipart.File
		Header      *multipart.FileHeader
		Type        string
	}

	Course struct {
		Id          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorId    string `json:"author_id"`
		Logo        string `json:"logo"`
		Type        string `json:"type"`
	}

	CourseTitle struct {
		Id            uint      `json:"id"`
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		Author        UserDto   `json:"author"`
		Links         []Link    `json:"links"`
		Collaborators []UserDto `json:"collaborators"`
		Logo          string    `json:"logo"`
		Type          string    `json:"type"`
	}

	CourseModules struct {
		Id      uint        `json:"id"`
		Name    string      `json:"name"`
		Modules []ModuleDto `json:"modules"`
		Logo    string      `json:"logo"`
		Type    string      `json:"type"`
	}

	CourseLink struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Link     string `json:"link"`
		CourseId int    `json:"course_id"`
	}

	CourseCollaborator struct {
		Id        string `json:"id"`
		UserEmail string `json:"user_email"`
		CourseId  int    `json:"course_id"`
	}
)

func NewCourseRequest(name, description, courseType, authorId string,
	logo multipart.File, header *multipart.FileHeader) CourseRequest {
	return CourseRequest{Name: name, Description: description, Logo: logo, Header: header,
		Type: courseType, AuthorId: authorId}
}

func NewCourse(id uint, name, description, logo, authorId, courseType string) Course {
	return Course{Id: id, Name: name, Description: description, Logo: logo, Type: courseType, AuthorId: authorId}
}
