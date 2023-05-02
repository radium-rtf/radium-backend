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

	SectionMultipleChoice struct {
		SlideId  uint
		OrderBy  uint
		Question string
		Markdown string
		Answers  SectionMultiAnswers
	}

	SectionChoice struct {
		SlideId  uint
		OrderBy  uint
		Question string
		Markdown string
	}

	SectionQuestion struct {
		SlideId  uint
		OrderBy  uint
		Question string
		Answer   string
	}

	SectionMultiAnswers map[string]bool
)
