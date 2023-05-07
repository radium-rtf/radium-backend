package entity

const (
	SectionAnswerOK = SectionAnswerVerdict("OK")
	SectionAnswerWA = SectionAnswerVerdict("WA")
)

type (
	SectionAnswerVerdict string

	SectionMultipleChoice struct {
		SlideId  uint
		OrderBy  uint
		Question string
		Markdown string
	}
)
