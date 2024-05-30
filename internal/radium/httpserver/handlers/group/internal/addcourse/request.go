package addcourse

import "github.com/google/uuid"

type AddCourse struct {
	CourseId uuid.UUID `json:"courseId"`
}
