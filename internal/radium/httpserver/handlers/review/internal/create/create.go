package create

import (
	"context"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/pkg/decode"
	"github.com/radium-rtf/radium-backend/pkg/resp"
	"net/http"
)

type creator interface {
	Create(ctx context.Context, review *entity.Review) (*entity.Review, error)
}

// @Tags review
// @Security ApiKeyAuth
// @Accept json
// @Param request body Review true " "
// @Success 201 {object} model.Review "created"
// @Router /v1/review [post]
func NewReview(creator creator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request Review
			ctx     = r.Context()
			userId  = ctx.Value("userId").(uuid.UUID)
		)

		err := decode.Json(r.Body, &request)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		review := request.toReview(userId)
		review, err = creator.Create(ctx, review)
		if err != nil {
			resp.Error(r, w, err)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, model.NewReview(review))
	}
}
