package postteacher

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/lib/decode"
	"net/http"
)

type creator interface {
	AddTeacher(ctx context.Context, email string) error
}

// @Tags role
// @Security ApiKeyAuth
// @Accept json
// @Param request body Email true "почта будущего преподавателя"
// @Success 201
// @Router /v1/role/teacher [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Email
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		email := request.Email
		err = creator.AddTeacher(ctx, email)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		render.Status(r, http.StatusCreated)
	}
}
