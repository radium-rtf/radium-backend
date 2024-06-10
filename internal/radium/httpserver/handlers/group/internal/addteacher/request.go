package addteacher

import "github.com/google/uuid"

type AddTeacher struct {
	UserId uuid.UUID `json:"userId"`
}
