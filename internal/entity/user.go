package entity

import (
	"github.com/google/uuid"
)

type (
	UserDto struct {
		Id     uuid.UUID `json:"id"`
		Email  string    `json:"email"`
		Name   string    `json:"name"`
		Avatar string    `json:"avatar"`
	}

	User struct {
		DBModel
		Avatar           string
		Email            string `gorm:"unique; not null"`
		Name             string `gorm:"not null"`
		Password         string `gorm:"not null"`
		VerificationCode string
		IsVerified       bool
		Courses          []*Course `gorm:"many2many:course_students"`
		Sessions         []Session
	}

	UpdateUserRequest struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	PasswordUpdate struct {
		New     string `json:"new"`
		Current string `json:"current"`
	}
)

func NewUserDto(user *User) UserDto {
	return UserDto{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
	}
}
