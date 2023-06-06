package entity

import "github.com/google/uuid"

type (
	Link struct {
		Id       uuid.UUID `json:"id" gorm:"default:gen_random_uuid()"`
		Name     string    `json:"name" gorm:"type:string"`
		Link     string    `json:"link" gorm:"type:string"`
		CourseId uuid.UUID `json:"courseId"`
	}
)
