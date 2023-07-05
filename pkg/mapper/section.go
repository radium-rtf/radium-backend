package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Section struct {
}

func (s Section) Sections(sections []*entity.Section, answers map[uuid.UUID]*entity.Answer) []*entity.SectionDto {
	dtos := make([]*entity.SectionDto, 0, len(sections))
	for _, section := range sections {
		var verdict = entity.VerdictEMPTY
		if answer, ok := answers[section.Id]; ok {
			verdict = answer.Verdict
		}
		dto := s.Section(section, verdict)
		dtos = append(dtos, dto)
	}
	return dtos
}

func (s Section) Section(section *entity.Section, verdict entity.Verdict) *entity.SectionDto {
	var sectionType entity.SectionType
	if section.MultiChoiceSection != nil {
		sectionType = entity.MultiChoiceType
	} else if section.ChoiceSection != nil {
		sectionType = entity.ChoiceType
	} else if section.TextSection != nil {
		sectionType = entity.TextType
	} else if section.ShortAnswerSection != nil {
		sectionType = entity.ShortAnswerType
	}
	return &entity.SectionDto{
		Id:       section.Id,
		PageId:   section.PageId,
		Order:    section.Order,
		Content:  section.Content(),
		MaxScore: section.MaxScore(),
		Verdict:  verdict,
		Variants: section.Variants(),
		Type:     sectionType,
	}
}

func (s Section) PostToSection(post *entity.SectionPost) *entity.Section {
	return &entity.Section{
		PageId:             post.PageId,
		Order:              post.Order,
		TextSection:        s.postToText(post.TextSection),
		ChoiceSection:      s.postToChoice(post.ChoiceSection),
		MultiChoiceSection: s.postToMultiChoice(post.MultiChoiceSection),
		ShortAnswerSection: s.postToShortAnswer(post.ShortAnswerSection),
	}
}

func (s Section) postToText(post *entity.TextSectionPost) *entity.TextSection {
	if post == nil {
		return nil
	}
	return &entity.TextSection{
		Content: post.Content,
	}
}

func (s Section) postToChoice(post *entity.ChoiceSectionPost) *entity.ChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.ChoiceSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (s Section) postToMultiChoice(post *entity.MultiChoiceSectionPost) *entity.MultiChoiceSection {
	if post == nil {
		return nil
	}
	return &entity.MultiChoiceSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Variants: post.Variants,
		Question: post.Question,
	}
}

func (s Section) postToShortAnswer(post *entity.ShortAnswerSectionPost) *entity.ShortAnswerSection {
	if post == nil {
		return nil
	}
	return &entity.ShortAnswerSection{
		MaxScore: post.MaxScore,
		Answer:   post.Answer,
		Question: post.Question,
	}
}
