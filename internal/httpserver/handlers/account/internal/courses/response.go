package courses

import "github.com/radium-rtf/radium-backend/internal/model"

type Courses struct {
	My              []*model.Course `json:"my"`
	Recommendations []*model.Course `json:"recommendations"`
}
