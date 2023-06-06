package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type moduleRoutes struct {
	uc usecase.ModuleUseCase
}

func newModuleRoutes(h chi.Router, useCase usecase.ModuleUseCase, signingKey string) {
	routes := moduleRoutes{uc: useCase}
	h.Route("/module", func(r chi.Router) {
		// r.Get("/{courseId}", handler(routes.getModules).HTTP)

		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingKey))
			r.Post("/", handler(routes.postModule).HTTP)
		})
	})
}

// @Tags module
// @Security ApiKeyAuth
// @Param       request body entity.ModuleRequest true "moduleRequest"
// @Success      201   {object} entity.ModuleDto       "created"
// @Router       /module [post]
func (r moduleRoutes) postModule(w http.ResponseWriter, request *http.Request) *appError {
	var module entity.ModuleRequest
	if err := render.DecodeJSON(request.Body, &module); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	moduleDto, err := r.uc.CreateModule(request.Context(), module)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	render.Status(request, http.StatusCreated)
	render.JSON(w, request, moduleDto)
	return nil
}

// // @Tags module
// // @Security ApiKeyAuth
// // @Param        courseId   path      integer  true  "course id"
// // @Success      200   {object} entity.Module "ok"
// // @Router       /module/{courseId} [get]
// func (r moduleRoutes) getModules(w http.ResponseWriter, request *http.Request) *appError {
// 	courseId := chi.URLParam(request, "courseId")
// 	id, err := strconv.Atoi(courseId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	courseModules, err := r.uc.GetCourseModules(request.Context(), id)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusOK)
// 	render.JSON(w, request, courseModules)
// 	return nil
// }
