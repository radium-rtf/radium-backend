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
		Id               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
		Avatar           string
		Email            string
		Name             string
		Password         string
		VerificationCode string
		IsVerified       bool
		Courses          []*Course `gorm:"many2many:course_students"`
		Sessions         []Session
	}

	UpdateUserRequest struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	UserName struct {
		Name string `json:"name"`
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
