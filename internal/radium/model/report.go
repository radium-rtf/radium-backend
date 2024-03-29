package model

import (
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	Report struct {
		Header ReportHeader `json:"header"`
		Rows   []ReportRow  `json:"rows"`
	}

	ReportRow struct {
		User   *User `json:"user"`
		Score  uint  `json:"score"`
		Values []int `json:"values"`
	}

	ReportHeader struct {
		MaxScore uint                `json:"maxScore"`
		Values   []ReportHeaderValue `json:"values"`
	}

	ReportHeaderValue struct {
		Name     string `json:"name"`
		IsModule bool   `json:"isModule"`
		MaxScore uint   `json:"maxScore"`
	}
)

func NewGroupReport(answersCollection *entity2.UsersAnswersCollection, course *entity2.Course) *Report {
	students := make(map[uuid.UUID]bool)
	for _, s := range course.Students {
		students[s.Id] = true
	}
	c := NewCourse(course, map[uuid.UUID][]*entity2.Answer{}, uuid.UUID{})
	headerValues := make([]ReportHeaderValue, 0, len(c.Modules)*3)
	reportRows := make([]ReportRow, 0, len(c.Modules))

	for _, module := range c.Modules {
		headerValue := ReportHeaderValue{MaxScore: module.MaxScore, IsModule: true, Name: module.Name}
		headerValues = append(headerValues, headerValue)
		for _, page := range module.Pages {
			headerValue := ReportHeaderValue{MaxScore: page.MaxScore, Name: page.Name}
			headerValues = append(headerValues, headerValue)
		}
	}

	for _, user := range answersCollection.Users {
		if !students[user.Id] {
			continue
		}
		answers, ok := answersCollection.AnswersByUserId[user.Id]
		if !ok {
			answers = &entity2.AnswersCollection{AnswerBySectionId: make(map[uuid.UUID][]*entity2.Answer)}
		}

		row := ReportRow{User: NewUser(user), Values: []int{}}
		c := NewCourse(course, answers.AnswerBySectionId, user.Id)
		row.Score = c.Score

		for _, m := range c.Modules {
			row.Values = append(row.Values, int(m.Score))
			for _, page := range m.Pages {
				row.Values = append(row.Values, int(page.Score))
			}
		}

		reportRows = append(reportRows, row)
	}

	header := ReportHeader{MaxScore: c.MaxScore, Values: headerValues}
	statement := &Report{Header: header, Rows: reportRows}
	return statement
}
