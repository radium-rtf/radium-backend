package entity

type (
	Collaborator struct {
		CourseId  uint   `json:"course_id"`
		UserEmail string `json:"user_email"`
	}

	CourseCollaborator struct {
		Id        string `json:"id"`
		UserEmail string `json:"user_email"`
		CourseId  uint   `json:"course_id"`
	}
)
