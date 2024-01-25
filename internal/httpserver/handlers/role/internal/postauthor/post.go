package postauthor

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/lib/resp"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"net/http"
	"strings"
)

type creator interface {
	AddAuthor(ctx context.Context, email string) error
}

// @Tags role
// @Security ApiKeyAuth
// @Accept json
// @Param request body Email true "почта будущего автора"
// @Success 201
// @Router /v1/role/author [post]
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
		err = creator.AddAuthor(ctx, email)
		if err != nil {
			resp.Error(r, w, err)
			return
		}
	}
}
