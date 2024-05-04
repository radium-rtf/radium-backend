package courses

import (
	"github.com/go-chi/render"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

// @Tags account
// @Security ApiKeyAuth
// @Success      200   {object} model.Main "ok"
// @Router       /v2/account/courses [get]
func NewV2(getter getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx    = r.Context()
			userId = ctx.Value("userId").(uuid.UUID)
		)

		user, err := getter.GetFullUser(ctx, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		recommendations, err := getter.GetRecommendations(ctx, userId, 30)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		authorship := append([]*entity2.Course{}, user.Author...)
		authorship = append(authorship, user.Coauthor...)

		my := make([]*entity2.Course, 0, len(user.Courses))
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

		response := model.NewMainCard(user.LastVisitedPage, my, authorship, recommendations, userId)

		render.Status(r, http.StatusOK)
		render.JSON(w, r, response)
	}
}
