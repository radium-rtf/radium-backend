package invite

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"net/http"
)

type connector interface {
	Join(ctx context.Context, userId uuid.UUID, courseId string) error
}

// @Tags group
// @Security ApiKeyAuth
// @Param        inviteCode   path      string  true  "inviteCode"
// @Success      200   {string}  string        " "
// @Router       /v1/group/invite/{inviteCode} [patch]
func New(connector connector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()

		inviteCode := chi.URLParam(r, "inviteCode")
		userId := r.Context().Value("userId").(uuid.UUID)

		err := connector.Join(ctx, userId, inviteCode)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
