package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"net/http"
)

type sectionRoutes struct {
	uc usecase.SectionUseCase
}

func newSectionRoutes(h chi.Router, useCase usecase.SectionUseCase, signingString string) {
	routes := sectionRoutes{uc: useCase}

	h.Route("/section", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingString))
			r.Post("/text", handler(routes.postText).HTTP)
		})
	})
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionTextPost true "создвние текствой секции"
// @Success      201   {object} entity.Slide "ok"
// @Router      /section/text [post]
func (r sectionRoutes) postText(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionTextPost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	section, err := r.uc.CreateText(request.Context(), sectionPost)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}
