package mapper

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Section struct {
}

func (Section) MultiChoice(section *entity.MultiChoiceSection, verdict entity.Verdict, score uint) *entity.MultiChoiceSectionDto {
	if section == nil {
		return nil
	}
	return &entity.MultiChoiceSectionDto{
		MaxScore: section.MaxScore,
		Verdict:  verdict,
		Score:    score,
		Question: section.Question,
		Variants: section.Variants,
	}
}

func (Section) Choice(section *entity.ChoiceSection, verdict entity.Verdict, score uint) *entity.ChoiceSectionDto {
	if section == nil {
		return nil
	}
	return &entity.ChoiceSectionDto{
		MaxScore: section.MaxScore,
		Question: section.Question,
		Variants: section.Variants,
		Score:    score,
		Verdict:  verdict,
	}
}

func (Section) ShortAnswer(section *entity.ShortAnswerSection, verdict entity.Verdict, score uint) *entity.ShortAnswerSectionDto {
	if section == nil {
		return nil
	}
	return &entity.ShortAnswerSectionDto{
		MaxScore: section.MaxScore,
		Question: section.Question,
		Score:    score,
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

func (p Section) Sections(sections []entity.Section, studentId string, verdicts map[string]map[uuid.UUID]entity.Verdict) []entity.SectionDto {
	dtos := make([]entity.SectionDto, 0, len(sections))
	for _, section := range sections {
		var score uint
		var verdict = entity.VerdictWA
		if userVerdicts, ok := verdicts[studentId]; ok {
			verdict, ok = userVerdicts[section.ID]
			if !ok {
				verdict = entity.VerdictWA
			} else if verdict == entity.VerdictOK {
				if section.ChoiceSection != nil {
					score = section.ChoiceSection.MaxScore
				} else if section.MultiChoiceSection != nil {
					score = section.MultiChoiceSection.MaxScore
				} else {
					score = section.ShortAnswerSection.MaxScore
				}
			} else {
				verdict = entity.VerdictWA
			}
		}
		dto := p.Section(section, verdict, score)
		dtos = append(dtos, dto)
	}
	return dtos
}

func (p Section) Section(section entity.Section, verdict entity.Verdict, score uint) entity.SectionDto {
	return entity.SectionDto{
		ID:                 section.ID,
		PageId:             section.PageId,
		Order:              section.Order,
		TextSection:        p.Text(section.TextSection),
		ChoiceSection:      p.Choice(section.ChoiceSection, verdict, score),
		ShortAnswerSection: p.ShortAnswer(section.ShortAnswerSection, verdict, score),
		MultiChoiceSection: p.MultiChoice(section.MultiChoiceSection, verdict, score),
	}
}
