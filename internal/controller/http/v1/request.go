package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"net/http"
	"strconv"
)

func newDestroy(r *http.Request) (*entity.Destroy, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "parse id")
	}
	isSoft, err := strconv.ParseBool(r.URL.Query().Get("is_soft"))
	if err != nil {
		isSoft = true
	}
	destroy := &entity.Destroy{Id: id, IsSoft: isSoft}
	return destroy, nil
}
