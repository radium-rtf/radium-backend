package answer

import (
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/answer/verdict"
)

type Checker struct {
}

func (c Checker) Check(section *entity.Section, answer *entity.Answer) (verdict.Verdict, error) {
	/*
		var verdictType verdict.Type

		isExists := func(v1, v2 any) bool {
			return !reflect.ValueOf(v1).IsNil() && !reflect.ValueOf(v2).IsNil()
		}

		switch {
		case isExists(answer.MultiChoice, section.MultiChoiceSection):
			verdictType = c.multiChoice(answer.MultiChoice, section.MultiChoiceSection)

		case isExists(answer.Choice, section.ChoiceSection):
			verdictType = c.choice(answer.Choice, section.ChoiceSection)

		case isExists(answer.ShortAnswer, section.ShortAnswerSection):
			verdictType = c.shortAnswer(answer.ShortAnswer, section.ShortAnswerSection)

		case isExists(answer.Permutation, section.PermutationSection):
			verdictType = c.permutation(answer.Permutation, section.PermutationSection)

		case isExists(answer.Answer, section.AnswerSection) ||
			isExists(answer.Code, section.CodeSection):
			verdictType = verdict.WAIT

		default:
			return verdict.Verdict{}, errChecker
		}

		return verdict.Verdict{Verdict: verdictType}, nil

	*/
	panic("not implemented")
}

/*
func (c Checker) multiChoice(answer *entity.MultichoiceSectionAnswer, section *entity.MultiChoiceSection) verdict.Type {
	answerArr := []string(answer.Answer)
	solutionArr := []string(section.Answer)
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

func (c Checker) permutation(permutation *entity.PermutationSectionAnswer, section *entity.PermutationSection) verdict.Type {
	ok := slices.Equal(permutation.Answer, section.Answer)
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
*/
