package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, teacher []*entity.TeacherCourseGroup) ([]*entity.TeacherCourseGroup, error)
}

// @Tags teacher
// @Security ApiKeyAuth
// @Param       request body Teacher true " "
// @Success      201   {string}  string        "created"
// @Router       /v1/teacher [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Teacher
			ctx     = r.Context()
		)

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		teacher := request.toCourses()
		teacher, err = creator.Create(ctx, teacher)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, teacher)
	}
}
