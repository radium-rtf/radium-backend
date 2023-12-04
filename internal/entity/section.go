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
	FileType        = SectionType("file")
)

type (
	SectionType string

	Section struct {
		bun.BaseModel `bun:"table:sections"`
		DBModel

		PageId uuid.UUID
		Page   Page `bun:"rel:belongs-to,join:page_id=id"`

		Order    float64
		MaxScore uint
		Type     SectionType

		Content  string
		Variants pq.StringArray

		Answer  string
		Answers pq.StringArray

		Keys      pq.StringArray
		FileTypes pq.StringArray

		MaxAttempts sql.NullInt16
	}
)

func NewSection(maxAttempts sql.NullInt16, pageId uuid.UUID, order float64,
	maxScore uint, content, answer string,
	variants, answers, keys, types []string, sectionType SectionType) (*Section, error) {
	if sectionType == TextType {
		maxScore = 0
	}
	section := &Section{
		DBModel:     DBModel{Id: uuid.New()},
		Order:       order,
		MaxScore:    maxScore,
		PageId:      pageId,
		Type:        sectionType,
		Content:     content,
		Variants:    variants,
		Answer:      answer,
		Answers:     answers,
		Keys:        keys,
		FileTypes:   types,
		MaxAttempts: maxAttempts,
	}

	switch sectionType {
	case PermutationType:
		variants := slices.Clone(answers)
		rand.Shuffle(len(variants), func(i, j int) {
			variants[i], variants[j] = variants[j], variants[i]
		})
		section.Variants = variants
	case ChoiceType:
	case ShortAnswerType:
	case MultiChoiceType:
	case FileType:
	case TextType:
	case CodeType:
	case AnswerType:
	default:
		return nil, errors.New("не удалось создать секцию")
	}

	return section, nil
}

func (s Section) GetVariants() []string {
	variants := []string(s.Variants)
	rand.Shuffle(len(variants), func(i, j int) {
		variants[i], variants[j] = variants[j], variants[i]
	})
	return variants
}

func (s Section) GetMaxScore() uint {
	if s.Type == TextType {
		return 0
	}
	return s.MaxScore
}
