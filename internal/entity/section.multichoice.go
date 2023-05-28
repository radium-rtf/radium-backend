package entity

type (
	SectionMultiChoicePost struct {
		SlideId  uint            `json:"slide_id"`
		OrderBy  uint            `json:"order_by"`
		Cost     uint            `json:"cost"`
		Question string          `json:"question"`
		Variants map[string]bool `json:"variants"`
	}

	SectionMultiChoice struct {
		Id       uint     `json:"id"`
		SlideId  uint     `json:"slide_id"`
		OrderBy  uint     `json:"order_by"`
		Cost     uint     `json:"cost"`
		Question string   `json:"question"`
		Answer   []string `json:"answer"`
		Variants []string `json:"variants"`
	}

	SectionMultiChoiceDto struct {
		Id       uint     `json:"id"`
		SlideId  uint     `json:"slide_id"`
		OrderBy  uint     `json:"order_by"`
		Cost     uint     `json:"cost"`
		Score    uint     `json:"score"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
	}

	SectionMultiChoiceAnswerPost struct {
		SectionId uint     `json:"section_id"`
		Answer    []string `json:"answer"`
	}

	SectionMultiChoiceAnswer struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		UserId    string               `json:"user_id"`
		Answer    []string             `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}

	SectionMultiChoiceAnswerDto struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		Answer    []string             `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}
)

func NewSectionMultiChoiceToDto(choice SectionMultiChoice) SectionMultiChoiceDto {
	return SectionMultiChoiceDto{
		Id:       choice.Id,
		SlideId:  choice.SlideId,
		OrderBy:  choice.OrderBy,
		Cost:     choice.Cost,
		Question: choice.Question,
		Variants: choice.Variants,
	}
}

func NewSectionMultiChoicePostToSection(post SectionMultiChoicePost) (SectionMultiChoice, error) {
	var section SectionMultiChoice
	variants := make([]string, 0, len(post.Variants))
	answer := make([]string, 0)
	for variant, isAnswer := range post.Variants {
		variants = append(variants, variant)
		if !isAnswer {
			continue
		}
		answer = append(answer, variant)
	}
	section = SectionMultiChoice{
		SlideId:  post.SlideId,
		OrderBy:  post.OrderBy,
		Cost:     post.Cost,
		Question: post.Question,
		Variants: variants,
		Answer:   answer,
	}
	return section, nil
}

func NewSectionMultiChoiceAnswerToDto(answer SectionMultiChoiceAnswer) SectionMultiChoiceAnswerDto {
	return SectionMultiChoiceAnswerDto{
		Id:        answer.Id,
		SectionId: answer.SectionId,
		Answer:    answer.Answer,
		Verdict:   answer.Verdict,
	}
}
