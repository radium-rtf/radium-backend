package getbyslug

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetBySlug(ctx context.Context, slug string) (*entity.Course, error)
	GetBySlugAndUser(ctx context.Context, slug string, userId uuid.UUID) (*entity.Course, error)
}

// @Tags course
// @Security ApiKeyAuth
// @Param        slug   path     string  true  "slug"
// @Success      200   {object} model.Course  "ok"
// @Router       /v1/course/slug/{slug} [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			slug = chi.URLParam(r, "slug")
		)
		userId, ok := ctx.Value("userId").(uuid.UUID)

		if !ok {
			responseIfNotAuthorized(getter, slug, w, r)
			return
		}

		course, err := getter.GetBySlugAndUser(ctx, slug, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		c := model.NewCourseWithUserGroups(course, map[uuid.UUID][]*entity.Answer{}, userId)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, c)
	}
}

func responseIfNotAuthorized(getter getter, slug string, w http.ResponseWriter, r *http.Request) {
	course, err := getter.GetBySlug(r.Context(), slug)
	if err != nil {
		resp.Error(r, w, err)
		return
	}

	c := model.NewCourse(course, map[uuid.UUID][]*entity.Answer{}, uuid.UUID{})
	render.Status(r, http.StatusOK)
	render.JSON(w, r, c)
}
