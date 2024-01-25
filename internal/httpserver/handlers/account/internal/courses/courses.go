package courses

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/model"
	"net/http"
)

type getter interface {
	GetFullUser(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetRecommendations(ctx context.Context, userId uuid.UUID, limit int) ([]*entity.Course, error)
}

// @Tags account
// @Security ApiKeyAuth
// @Success      200   {object} Courses "ok"
// @Router       /v1/account/courses [get]
func New(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		user, err := getter.GetFullUser(ctx, userId)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		recommendations, err := getter.GetRecommendations(ctx, userId, 30)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
			return
		}

		authorship := append([]*entity.Course{}, user.Author...)
		authorship = append(authorship, user.Coauthor...)

		my := make([]*entity.Course, 0, len(user.Courses))
		authorshipSet := make(map[uuid.UUID]bool, len(authorship))
		for _, c := range authorship {
			authorshipSet[c.Id] = true
		}

		for _, c := range user.Courses {
			if authorshipSet[c.Id] {
				continue
			}
			my = append(my, c)
		}

		response := Courses{
			Authorship:      model.NewCourses(authorship, userId),
			My:              model.NewCourses(my, userId),
			Recommendations: model.NewCourses(recommendations, userId),
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, response)
	}
}
