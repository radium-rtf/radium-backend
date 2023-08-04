package answer

import (
	"reflect"

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
	case answer.Answer != nil:
		verdictType = verdict.EMPTY
	case answer.Code != nil:
		verdictType = verdict.WAIT
	default:
		return verdict.Verdict{}, errEmptyAnswer
	}

	return verdict.Verdict{Verdict: verdictType}, nil
}

func (c Checker) multiChoice(answer *entity.MultichoiceSectionAnswer, section *entity.MultiChoiceSection) verdict.Type {
	answerArr := []string(answer.Answer)
	solutionArr := []string(section.Answer)
	verdictType := verdict.OK
	ok := reflect.DeepEqual(answerArr, solutionArr)
	if !ok {
		verdictType = verdict.WA
	}
	return verdictType
}

func (c Checker) choice(answer *entity.ChoiceSectionAnswer, section *entity.ChoiceSection) verdict.Type {
	verdictType := verdict.OK
	ok := answer.Answer == section.Answer
	if !ok {
		verdictType = verdict.WA
	}
	return verdictType
}

func (c Checker) shortAnswer(answer *entity.ShortAnswerSectionAnswer, section *entity.ShortAnswerSection) verdict.Type {
	verdictType := verdict.OK
	ok := answer.Answer == section.Answer
	if !ok {
		verdictType = verdict.WA
	}
	return verdictType
}
