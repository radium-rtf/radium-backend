package entity

import (
	"mime/multipart"
)

type (
	CourseRequest struct {
		Name        string
		Description string
		Logo        multipart.File
		Header      *multipart.FileHeader
		Type        string
		Chat        string
	}

	Course struct {
		Id          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Logo        string `json:"logo"`
		Chat        string `json:"chat"`
		Type        string `json:"type"`
	}
)

func NewCourseRequest(name, description, courseType string, chat string, logo multipart.File, header *multipart.FileHeader) CourseRequest {
	return CourseRequest{Name: name, Description: description, Logo: logo, Header: header, Type: courseType, Chat: chat}
}

func NewCourse(id uint, name, description, logo, chat, courseType string) Course {
	return Course{Id: id, Name: name, Description: description, Logo: logo, Type: courseType, Chat: chat}
}
