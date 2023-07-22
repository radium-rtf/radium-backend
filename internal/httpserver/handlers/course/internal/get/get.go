package get

import (
	"context"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetCourses(ctx context.Context) ([]*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Success      200   {object} model.Course        "ok"
// @Router       /v1/course [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := getter.GetCourses(r.Context())
		if err != nil {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourses(courses)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
