package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Section struct {
}

func (Section) MultiChoice(section *entity.MultiChoiceSection, verdict entity.Verdict) *entity.MultiChoiceSectionDto {
	if section == nil {
		return nil
	}
	return &entity.MultiChoiceSectionDto{
		MaxScore: section.MaxScore,
		Verdict:  verdict,
		Question: section.Question,
		Variants: section.Variants,
	}
}

func (Section) Choice(section *entity.ChoiceSection, verdict entity.Verdict) *entity.ChoiceSectionDto {
	if section == nil {
		return nil
	}
	return &entity.ChoiceSectionDto{
		MaxScore: section.MaxScore,
		Question: section.Question,
		Variants: section.Variants,
		Verdict:  verdict,
	}
}

func (Section) ShortAnswer(section *entity.ShortAnswerSection, verdict entity.Verdict) *entity.ShortAnswerSectionDto {
	if section == nil {
		return nil
	}
	return &entity.ShortAnswerSectionDto{
		MaxScore: section.MaxScore,
		Question: section.Question,
		Verdict:  verdict,
	}
}

func (Section) Text(section *entity.TextSection) *entity.TextSectionDto {
	if section == nil {
		return nil
	}
	return &entity.TextSectionDto{
		Content: section.Content,
	}
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
	return &entity.SectionDto{
		Id:                 section.Id,
		PageId:             section.PageId,
		Order:              section.Order,
		TextSection:        s.Text(section.TextSection),
		ChoiceSection:      s.Choice(section.ChoiceSection, verdict),
		ShortAnswerSection: s.ShortAnswer(section.ShortAnswerSection, verdict),
		MultiChoiceSection: s.MultiChoice(section.MultiChoiceSection, verdict),
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
