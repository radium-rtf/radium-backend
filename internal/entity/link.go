package entity

type (
	LinkRequest struct {
		CourseId uint   `json:"course_id"`
		Name     string `json:"name"`
		Link     string `json:"link"`
	}

	Link struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Link     string `json:"link"`
		CourseId uint
	}

	LinkDto struct {
		Name string `json:"name"`
		Link string `json:"link"`
	}
)
