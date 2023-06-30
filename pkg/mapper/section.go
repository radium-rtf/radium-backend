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

func (p Section) Sections(sections []entity.Section, answers map[uuid.UUID]*entity.Answer) []*entity.SectionDto {
	dtos := make([]*entity.SectionDto, 0, len(sections))
	for _, section := range sections {
		var verdict = entity.VerdictEMPTY
		if answer, ok := answers[section.ID]; ok {
			verdict = answer.Verdict
		}
		dto := p.Section(section, verdict)
		dtos = append(dtos, dto)
	}
	return dtos
}

func (p Section) Section(section entity.Section, verdict entity.Verdict) *entity.SectionDto {
	return &entity.SectionDto{
		ID:                 section.ID,
		PageId:             section.PageId,
		Order:              section.Order,
		TextSection:        p.Text(section.TextSection),
		ChoiceSection:      p.Choice(section.ChoiceSection, verdict),
		ShortAnswerSection: p.ShortAnswer(section.ShortAnswerSection, verdict),
		MultiChoiceSection: p.MultiChoice(section.MultiChoiceSection, verdict),
	}
}
