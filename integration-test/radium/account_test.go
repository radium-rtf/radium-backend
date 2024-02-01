package radium

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gavv/httpexpect/v2"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"net/http"
	"testing"
)

func TestAccount(t *testing.T) {
	e := newHttpExpectWithAuth(t)
	getAccount(e)
	updateAccount(t, e)
	getAccountCourses(e)
}

func getAccount(e *httpexpect.Expect) *model.User {
	var user model.User
	e.GET("/v1/account").
		Expect().
		Status(http.StatusOK).
		JSON().
		Decode(&user)
	return &user
}

func updateAccount(t *testing.T, e *httpexpect.Expect) {
	type update struct {
		Avatar string `json:"avatar"`
		Name   string `json:"name"`
	}

	json := update{Avatar: gofakeit.URL(), Name: gofakeit.LetterN(10)}
	var user model.User
	e.PATCH("/v1/account").
		WithJSON(json).
		Expect().
		Status(http.StatusOK).
		JSON().
		Decode(&user)
	if t.Failed() {
		return
	}

	user = *getAccount(e)
	if t.Failed() {
		return
	}

	if user.Name != json.Name || user.Avatar != json.Avatar {
		t.Fail()
	}
}

func getAccountCourses(e *httpexpect.Expect) {
	e.GET("/v1/account/courses").
		Expect().
		Status(http.StatusOK)
}
