package answer

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
	"slices"
)

type Checker struct {
}

func (c Checker) Check(section *entity.Section, answer *entity.Answer) (verdict.Verdict, error) {
	if section.Type != answer.Type {
		return verdict.Verdict{}, ErrInvalidSectionId
	}

	var verdictType verdict.Type

	switch {
	case answer.Type == entity.MultiChoiceType:
		verdictType = c.multiChoice(answer, section)

	case answer.Type == entity.ChoiceType:
		verdictType = c.choice(answer, section)

	case answer.Type == entity.ShortAnswerType:
		verdictType = c.shortAnswer(answer, section)

	case answer.Type == entity.PermutationType:
		verdictType = c.permutation(answer, section)

	case answer.Type == entity.MappingType:
		verdictType = c.mapping(answer, section)

	case answer.Type == entity.AnswerType ||
		answer.Type == entity.CodeType:
		verdictType = verdict.WAIT

	default:
		return verdict.Verdict{}, ErrNotImpl
	}

	return verdict.Verdict{Verdict: verdictType}, nil
}

func (c Checker) multiChoice(answer *entity.Answer, section *entity.Section) verdict.Type {
	answerArr := []string(answer.Answers)
	solutionArr := []string(section.Answers)
	if len(answerArr) != len(solutionArr) {
		return verdict.WA
	}

	solutionCounter := c.toCounter(solutionArr)
	answerCounter := c.toCounter(answerArr)
	if len(answerCounter) != len(solutionCounter) {
		return verdict.WA
	}

	for ans, count := range answerCounter {
		if sCount, ok := solutionCounter[ans]; ok && count == sCount {
			continue
		}
		return verdict.WA
	}

	return verdict.OK
}

func (c Checker) choice(answer *entity.Answer, section *entity.Section) verdict.Type {
	ok := answer.Answer == section.Answer
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}

func (c Checker) shortAnswer(answer *entity.Answer, section *entity.Section) verdict.Type {
	ok := answer.Answer == section.Answer // TODO: (не)учитывать caps lock
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}

func (c Checker) permutation(permutation *entity.Answer, section *entity.Section) verdict.Type {
	ok := slices.Equal(permutation.Answers, section.Answers)
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}

func (Checker) toCounter(arr []string) map[string]int {
	m := make(map[string]int, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
		m[v] += 1
	}
	return m
}

func (c Checker) mapping(answer *entity.Answer, section *entity.Section) verdict.Type {
	ok := slices.Equal(answer.Answers, section.Answers)
	if !ok {
		return verdict.WA
	}
	return verdict.OK
}
