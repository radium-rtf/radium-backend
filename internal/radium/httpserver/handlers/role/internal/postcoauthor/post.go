package postcoauthor

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
	"strings"
)

type creator interface {
	AddCoauthor(ctx context.Context, email string, courseId, authorId uuid.UUID) error
}

// @Tags role
// @Security ApiKeyAuth
// @Accept json
// @Param request body Request true "почта и курс"
// @Success 201
// @Router /v1/role/coauthor [post]
func New(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Request
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		request.Email = strings.ToLower(request.Email)
		err = creator.AddCoauthor(ctx, request.Email, request.CourseId, userId)
		if err != nil {
			resp.Error(r, w, err)
			return
		}
	}
}
