package updaterole

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type updater interface {
	UpdateRole(ctx context.Context, user *entity.User) error
}

// @Tags role
// @Security ApiKeyAuth
// @Param        id   path      string  true  "userId"
// @Success      200   {string}  model.User         " "
// @Router       /v1/role/{id} [patch]
func New(updater updater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx     = r.Context()
			request Roles
		)

		userId, err := uuid.Parse(chi.URLParam(r, "id"))

		if err != nil {
			resp.Error(r, w, err)
			return
		}

		user := request.ToUser(userId)

		err = updater.UpdateRole(ctx, user)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		// dto := model.NewGroup(group)
		render.Status(r, http.StatusOK)
		// render.JSON(w, r, dto)
	}
}
