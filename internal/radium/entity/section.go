package entity

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/http"
	"github.com/uptrace/bun"
	"math/rand"
	"slices"
	"strings"
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
	MediaType       = SectionType("media")
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

		Url  sql.NullString
		File *File `bun:"rel:belongs-to,join:url=url"`

		Keys      pq.StringArray
		FileTypes pq.StringArray

		MaxAttempts sql.NullInt16

		UsersAnswers []*Answer `bun:"rel:has-many,join:id=section_id"`
	}
)

func NewSection(maxAttempts sql.NullInt16, pageId uuid.UUID, order float64, maxScore uint, content, answer string, variants, answers, keys, types []string, sectionType SectionType, url string) (*Section, error) {
	if sectionType == TextType || sectionType == MediaType {
		maxScore = 0
		maxAttempts = sql.NullInt16{}
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
		Url:         sql.NullString{String: url, Valid: len(url) != 0},
	}

	shuffledAnswers := slices.Clone(answers)
	rand.Shuffle(len(shuffledAnswers), func(i, j int) {
		shuffledAnswers[i], shuffledAnswers[j] = shuffledAnswers[j], shuffledAnswers[i]
	})
	switch sectionType {
	case PermutationType:
		section.Variants = shuffledAnswers
	case MappingType:
		section.Variants = shuffledAnswers
	case ChoiceType:
	case ShortAnswerType:
	case MultiChoiceType:
	case FileType:
	case TextType:
	case CodeType:
	case AnswerType:
	case MediaType:
		file, err := http.NewFileFromUrl(section.Url.String)
		if err != nil {
			return nil, err
		}
		Type := file.Type
		isVideo := strings.HasPrefix(Type, "video")
		isImage := strings.HasPrefix(Type, "image")
		if Type != "iframe" && !isVideo && !isImage {
			return nil, errors.New("ссылка может быть только iframe или с content-type image или video")
		}

		section.File = &File{
			Url:  file.Url,
			Type: file.Type,
			Name: file.Name,
			Size: file.Size,
		}

		section.Url = sql.NullString{String: section.File.Url, Valid: true}

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
