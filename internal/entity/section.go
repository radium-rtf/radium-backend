package entity

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/uptrace/bun"
	"math/rand"
	"slices"
)

const (
	ChoiceType      = SectionType("choice")
	MultiChoiceType = SectionType("multiChoice")
	TextType        = SectionType("text")
	ShortAnswerType = SectionType("shortAnswer")
	AnswerType      = SectionType("answer")
	CodeType        = SectionType("code")
	PermutationType = SectionType("permutation")
	MappingType     = SectionType("mapping")
)

type (
	SectionType string

	Section struct {
		bun.BaseModel `bun:"table:sections"`
		DBModel

		PageId   uuid.UUID
		Order    float64
		MaxScore uint
		Type     SectionType

		Content  string
		Variants pq.StringArray

		Answer  string
		Answers pq.StringArray

		Keys pq.StringArray

		MaxAttempts sql.NullInt16
	}
)

func NewSection(maxAttempts sql.NullInt16, pageId uuid.UUID, order float64, maxScore uint, content string, answer string, variants []string, answers []string, sectionType SectionType, keys []string) (*Section, error) {
	if sectionType == TextType {
		maxScore = 0
	}
	section := &Section{
		DBModel:     DBModel{Id: uuid.New()},
		Order:       order,
		MaxScore:    maxScore,
		PageId:      pageId,
		Type:        sectionType,
		MaxAttempts: maxAttempts,
	}

	section.Content = content
	switch sectionType {
	case MappingType:
		section.Keys = keys
		section.Answers = answers
		section.Variants = variants
	case PermutationType:
		variants := slices.Clone(answers)
		rand.Shuffle(len(variants), func(i, j int) {
			variants[i], variants[j] = variants[j], variants[i]
		})
		section.Answers = answers
		section.Variants = variants
	case ChoiceType:
		section.Answer = answer
		section.Variants = variants
	case ShortAnswerType:
		section.Answer = answer
	case MultiChoiceType:
		section.Answers = answers
		section.Variants = variants
	case TextType:
	case CodeType:
	case AnswerType:
	default:
		return nil, errors.New("не удалось создать секцию")
	}

	return section, nil
}

func (s Section) GetVariants() []string {
	if s.Type != MappingType {
		variants := []string(s.Variants)
		rand.Shuffle(len(variants), func(i, j int) {
			variants[i], variants[j] = variants[j], variants[i]
		})
		return variants
	}
	return s.Variants
}

func (s Section) GetMaxScore() uint {
	if s.Type == TextType {
		return 0
	}
	return s.MaxScore
}
