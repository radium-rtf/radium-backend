package addstudent

import "github.com/google/uuid"

type AddStudent struct {
	UserId uuid.UUID `json:"userId"`
}
