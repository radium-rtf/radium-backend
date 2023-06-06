//go:build ignore
// +build ignore

package entity

type (
	Collaborator struct {
		CourseId  uint   `json:"courseId"`
		UserEmail string `json:"userEmail"`
	}

	CourseCollaborator struct {
		Id        string `json:"id"`
		UserEmail string `json:"userEmail"`
		CourseId  uint   `json:"courseId"`
	}
)

func (CourseCollaborator) TableName() string {
	return "course_collaborators"
}
