package radium

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
)

type Courses struct {
	Empty *model.Course
	Full  *model.Course
}

func TestCourse(t *testing.T) {
	e := newHttpExpectWithAuth(t)
	courses := createCourses(t, e)
	if t.Failed() {
		return
	}

	course := courses.Full
	getCourse(t, e, course)
}

func createCourses(t *testing.T, e *httpexpect.Expect) Courses {
	type Course struct {
		Name             string       `json:"name" validate:"max=128"`
		ShortDescription string       `json:"shortDescription" validate:"max=400"`
		Description      string       `json:"description" validate:"max=3000"`
		Logo             string       `json:"logo" validate:"url"`
		Banner           string       `json:"banner" validate:"url"`
		Links            []model.Link `json:"links" validate:"dive"`
	}

	createCourse := func(json Course) *model.Course {
		var course model.Course
		e.POST("/v1/course").
			WithJSON(json).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Decode(&course)
		return &course
	}
	var courses Courses
	var json Course
	courses.Empty = createCourse(json)

	json = Course{
		Name:             gofakeit.LetterN(10),
		ShortDescription: gofakeit.LetterN(10),
		Description:      gofakeit.LetterN(20),
		Logo:             gofakeit.URL(),
		Banner:           gofakeit.URL(),
		Links:            []model.Link{{Link: gofakeit.URL(), Name: gofakeit.LetterN(10)}},
	}
	courses.Full = createCourse(json)

	if !FieldsFirstStructEqualSecond(&json, courses.Full.Access) || courses.Full.Access == "" {
		t.Fail()
		return courses
	}

	json.Links[0].Id = courses.Full.Links[0].Id
	if !FieldsFirstStructEqualSecond(json.Links[0], courses.Full.Links[0]) {
		t.Fail()
		return courses
	}

	return courses
}

func getCourse(t *testing.T, e *httpexpect.Expect, course *model.Course) {
	type test struct {
		Url  string
		Name string
	}
	tests := []test{
		{Url: "/v1/course/" + course.Id.String(), Name: "get by id"},
		{Url: "/v1/course/slug/" + course.Slug, Name: "get by slug"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			var received model.Course
			e.GET(tt.Url).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				Decode(&received)

			if !reflect.DeepEqual(&received, course) {
				t.Fail()
			}
		})
	}
}
