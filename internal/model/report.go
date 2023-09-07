package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
)

type (
	ReportRow struct {
		User   User  `json:"user"`
		Score  uint  `json:"score"`
		Values []int `json:"values"`
	}

	ReportHeader struct {
		MaxScore uint      `json:"maxScore"`
		Modules  []*Module `json:"modules"`
	}

	Report struct {
		Header ReportHeader `json:"header"`
		Rows   []ReportRow  `json:"rows"`
	}
)

func NewGroupReport(answersCollection *entity.UsersAnswersCollection, course *entity.Course) *Report {
	c := NewCourse(course, map[uuid.UUID]*entity.Answer{})
	for _, module := range c.Modules {
		for _, page := range module.Pages {
			page.Sections = []*Section{}
		}
	}

	header := ReportHeader{MaxScore: c.MaxScore, Modules: c.Modules}
	statement := &Report{Header: header}

	for _, user := range answersCollection.Users {
		answers, ok := answersCollection.AnswersByUserId[user.Id]
		if !ok {
			answers = &entity.AnswersCollection{AnswerBySectionId: make(map[uuid.UUID]*entity.Answer)}
		}
		row := ReportRow{User: NewUser(user)}
		c := NewCourse(course, answers.AnswerBySectionId)
		row.Score = c.Score
		for _, m := range c.Modules {
			row.Values = append(row.Values, int(m.Score))
			for _, page := range m.Pages {
				row.Values = append(row.Values, int(page.Score))
			}
		}
		statement.Rows = append(statement.Rows, row)
	}

	return statement
}
