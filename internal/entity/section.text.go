package entity

type (
	SectionTextPost struct {
		SlideId  uint   `json:"slide_id"`
		OrderBy  uint   `json:"order_by"`
		Markdown string `json:"markdown"`
	}

	SectionText struct {
		Id       uint   `json:"id"`
		SlideId  uint   `json:"slide_id"`
		OrderBy  uint   `json:"order_by"`
		Markdown string `json:"markdown"`
	}
)
