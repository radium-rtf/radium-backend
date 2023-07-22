package verify

import (
	"context"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type verify interface {
	VerifyEmail(ctx context.Context, verificationCode string) (bool, error)
}

func New(verify verify) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request  Request
			response Response
			ctx      = r.Context()
		)

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
		}

		success, err := verify.VerifyEmail(ctx, request.VerificationCode)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err.Error())
		}

		response.Success = success
		render.Status(r, http.StatusOK)
		render.JSON(w, r, response)
	}
}
