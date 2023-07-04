package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"
)

type sectionRoutes struct {
	uc     usecase.SectionUseCase
	mapper mapper.Section
}

func newSectionRoutes(useCase usecase.SectionUseCase, signingString string) *chi.Mux {
	routes := sectionRoutes{uc: useCase}
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingString))
		r.Post("/", handler(routes.createSection).HTTP)
		r.Delete("/{id}", handler(routes.delete).HTTP)
	})
	return r
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param request body entity.SectionPost true "Информация для раздела"
// @Success 201 {object} entity.SectionDto "created"
// @Router /section [post]
func (r sectionRoutes) createSection(w http.ResponseWriter, request *http.Request) *appError {
	sectionRequest := &entity.SectionPost{}
	err := json.NewDecoder(request.Body).Decode(&sectionRequest)
	if err != nil {
		return newAppError(errors.New(err.Error()), http.StatusBadRequest)
	}
	section := r.mapper.PostToSection(sectionRequest)
	section, err = r.uc.CreateSection(request.Context(), section)
	if err != nil {
		return newAppError(errors.New(err.Error()), http.StatusBadRequest)
	}

	dto := r.mapper.Section(section, entity.VerdictEMPTY)
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, dto)

	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /section/{id} [delete]
func (r sectionRoutes) delete(w http.ResponseWriter, request *http.Request) *appError {
	destroy, err := newDestroy(request)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	err = r.uc.Delete(request.Context(), destroy)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
