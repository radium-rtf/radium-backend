package entity

type (
	SectionQuestionPost struct {
		SlideId       uint   `json:"slide_id"`
		CaseSensitive bool   `json:"case_sensitive"`
		OrderBy       uint   `json:"order_by"`
		Cost          uint   `json:"cost"`
		Question      string `json:"question"`
		Answer        string `json:"answer"`
	}

	SectionQuestion struct {
		Id            uint   `json:"id"`
		SlideId       uint   `json:"slide_id"`
		CaseSensitive bool   `json:"case_sensitive"`
		OrderBy       uint   `json:"order_by"`
		Cost          uint   `json:"cost"`
		Question      string `json:"question"`
		Answer        string `json:"answer"`
	}

	SectionQuestionAnswerPost struct {
		SectionId uint   `json:"section_id"`
		Answer    string `json:"answer"`
	}

	SectionQuestionAnswer struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		UserId    string               `json:"user_id"`
		Answer    string               `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}

	SectionQuestionAnswerDto struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		Answer    string               `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}

	SectionQuestionDto struct {
		Id            uint   `json:"id"`
		SlideId       uint   `json:"slide_id"`
		CaseSensitive bool   `json:"case_sensitive"`
		OrderBy       uint   `json:"order_by"`
		Cost          uint   `json:"cost"`
		Score         uint   `json:"score"`
		Question      string `json:"question"`
	}
)

func NewSectionQuestionAnswerToDto(answer SectionQuestionAnswer) SectionQuestionAnswerDto {
	return SectionQuestionAnswerDto{
		Id:        answer.Id,
		SectionId: answer.SectionId,
		Answer:    answer.Answer,
		Verdict:   answer.Verdict,
	}
}
