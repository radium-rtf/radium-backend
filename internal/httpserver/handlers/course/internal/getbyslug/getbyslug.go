package getbyslug

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetBySlug(ctx context.Context, slug string) (*entity.Course, error)
}

// @Tags course
// @Param        slug   path     string  true  "course slug"
// @Success      200   {object} model.Course  "ok"
// @Router       /v1/course/slug/{slug} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			slug = chi.URLParam(r, "slug")
		)

		course, err := getter.GetBySlug(ctx, slug)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourse(course)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
