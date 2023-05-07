package entity

import "errors"

type (
	SectionChoicePost struct {
		SlideId  uint            `json:"slide_id"`
		OrderBy  uint            `json:"order_by"`
		Cost     uint            `json:"cost"`
		Question string          `json:"question"`
		Variants map[string]bool `json:"variants"`
	}

	SectionChoice struct {
		Id       uint     `json:"id"`
		SlideId  uint     `json:"slide_id"`
		OrderBy  uint     `json:"order_by"`
		Cost     uint     `json:"cost"`
		Question string   `json:"question"`
		Answer   string   `json:"answer"`
		Variants []string `json:"variants"`
	}

	SectionChoiceDto struct {
		Id       uint     `json:"id"`
		SlideId  uint     `json:"slide_id"`
		OrderBy  uint     `json:"order_by"`
		Cost     uint     `json:"cost"`
		Score    uint     `json:"score"`
		Question string   `json:"question"`
		Variants []string `json:"variants"`
	}

	SectionChoiceAnswerPost struct {
		SectionId uint   `json:"section_id"`
		Answer    string `json:"answer"`
	}

	SectionChoiceAnswer struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		UserId    string               `json:"user_id"`
		Answer    string               `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}

	SectionChoiceAnswerDto struct {
		Id        uint                 `json:"id"`
		SectionId uint                 `json:"section_id"`
		Answer    string               `json:"answer"`
		Verdict   SectionAnswerVerdict `json:"verdict"`
	}
)

func NewSectionChoiceToDto(choice SectionChoice) SectionChoiceDto {
	return SectionChoiceDto{
		Id:       choice.Id,
		SlideId:  choice.SlideId,
		OrderBy:  choice.OrderBy,
		Cost:     choice.Cost,
		Question: choice.Question,
		Variants: choice.Variants,
	}
}

func NewSectionChoicePostToSection(post SectionChoicePost) (SectionChoice, error) {
	var section SectionChoice
	variants := make([]string, 0, len(post.Variants))
	answer := ""
	for variant, isAnswer := range post.Variants {
		variants = append(variants, variant)
		if !isAnswer {
			continue
		}
		if answer != "" {
			return section, errors.New("у такой секции может быть только один ответ")
		}
		answer = variant
	}
	section = SectionChoice{
		SlideId:  post.SlideId,
		OrderBy:  post.OrderBy,
		Cost:     post.Cost,
		Question: post.Question,
		Variants: variants,
		Answer:   answer,
	}
	return section, nil
}

func NewSectionChoiceAnswerToDto(answer SectionChoiceAnswer) SectionChoiceAnswerDto {
	return SectionChoiceAnswerDto{
		Id:        answer.Id,
		SectionId: answer.SectionId,
		Answer:    answer.Answer,
		Verdict:   answer.Verdict,
	}
}
