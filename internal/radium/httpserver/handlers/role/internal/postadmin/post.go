package postadmin

import (
	"context"
	"net/http"
	"strings"

	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
)

type creator interface {
	AddAdmin(ctx context.Context, email string) error
}

// @Tags role
// @Security ApiKeyAuth
// @Accept json
// @Param request body Email true "почта будущего админа"
// @Success 201
// @Router /v1/role/admin [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Email
			ctx     = r.Context()
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		email := strings.ToLower(request.Email)
		err = creator.AddAdmin(ctx, email)
		if err != nil {
			resp.Error(r, w, err)
			return
		}
	}
}
