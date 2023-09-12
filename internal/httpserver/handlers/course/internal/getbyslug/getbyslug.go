package getbyslug

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetBySlug(ctx context.Context, slug string) (*entity.Course, error)
}

type answersGetter interface {
	GetBySections(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (map[uuid.UUID]*entity.Answer, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Param        slug   path     string  true  "course slug"
// @Success      200   {object} model.Course  "ok"
// @Router       /v1/course/slug/{slug} [get]
func New(getter getter, answersGetter answersGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			slug = chi.URLParam(r, "slug")
		)
		userId, ok := ctx.Value("userId").(uuid.UUID)

		course, err := getter.GetBySlug(ctx, slug)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		if !ok {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, model.NewCourse(course, map[uuid.UUID]*entity.Answer{}))
			return
		}

		sectionsIds := course.SectionsIds()
		answers, err := answersGetter.GetBySections(ctx, sectionsIds, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		c := model.NewCourse(course, answers)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}
