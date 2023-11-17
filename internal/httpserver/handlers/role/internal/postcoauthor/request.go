package postcoauthor

import "github.com/google/uuid"

type Request struct {
	Email    string `json:"email" validate:"email"`
	CourseId uuid.UUID
}
