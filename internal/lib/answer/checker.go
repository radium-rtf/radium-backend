package answer

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type Checker struct {
}

func (c Checker) Check(section *entity.Section, answer *entity.Answer) (verdict.Verdict, error) {
	var verdictType verdict.Type

	switch {
	case answer.MultiChoice != nil:
		verdictType = c.multiChoice(answer.MultiChoice, section.MultiChoiceSection)
	case answer.Choice != nil:
		verdictType = c.choice(answer.Choice, section.ChoiceSection)
	case answer.ShortAnswer != nil:
		verdictType = c.shortAnswer(answer.ShortAnswer, section.ShortAnswerSection)
	case answer.Answer != nil || answer.Code != nil:
		verdictType = verdict.WAIT
	default:
		return verdict.Verdict{}, errEmptyAnswer
	}

	return verdict.Verdict{Verdict: verdictType}, nil
}

func (c Checker) multiChoice(answer *entity.MultichoiceSectionAnswer, section *entity.MultiChoiceSection) verdict.Type {
	answerArr := []string(answer.Answer)
	solutionArr := []string(section.Answer)
	if len(answerArr) != len(solutionArr) {
		return verdict.WA
	}

	solutionMap := c.toMap(solutionArr)
	answerMap := c.toMap(answerArr)
	if len(answerArr) != len(solutionArr) {
		return verdict.WA
	}

	for ans, count := range answerMap {
		if sCount, ok := solutionMap[ans]; ok && count == sCount {
			continue
		}
		return verdict.WA
	}

	return verdict.OK
}

func (c Checker) choice(answer *entity.ChoiceSectionAnswer, section *entity.ChoiceSection) verdict.Type {
	ok := answer.Answer == section.Answer
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}

func (c Checker) shortAnswer(answer *entity.ShortAnswerSectionAnswer, section *entity.ShortAnswerSection) verdict.Type {
	ok := answer.Answer == section.Answer // TODO: (не)учитывать caps lock
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}

func (Checker) toMap(arr []string) map[string]int {
	m := make(map[string]int, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
		m[v] += 1
	}
	return m
}
