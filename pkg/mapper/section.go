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
		var (
			verdict    = entity.VerdictEMPTY
			score      = uint(0)
			answerStr  = ""
			answersArr []string
		)

		answer, ok := answers[section.Id]
		if ok {
			verdict = answer.Verdict
			score = answer.Score(section)
			answerStr = answer.AnswerStr()
			answersArr = answer.Answers()
		}

		dto := s.Section(section, verdict, score, answerStr, answersArr)
		dtos = append(dtos, dto)
	}
	return dtos
}

func (s Section) Section(section *entity.Section, verdict entity.Verdict,
	score uint, answer string, answers []string) *entity.SectionDto {
	var sectionType entity.SectionType
	if section.MultiChoiceSection != nil {
		sectionType = entity.MultiChoiceType
	} else if section.ChoiceSection != nil {
		sectionType = entity.ChoiceType
	} else if section.TextSection != nil {
		sectionType = entity.TextType
	} else if section.ShortAnswerSection != nil {
		sectionType = entity.ShortAnswerType
	} else if section.AnswerSection != nil {
		sectionType = entity.AnswerType
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
		Score:    score,
		Answers:  answers,
		Answer:   answer,
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
		AnswerSection:      s.postToAnswer(post.AnswerSection),
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

func (s Section) postToAnswer(post *entity.AnswerSectionPost) *entity.AnswerSection {
	if post == nil {
		return nil
	}
	return &entity.AnswerSection{
		Question: post.Question,
		MaxScore: post.MaxScore,
	}
}
