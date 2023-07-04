package v1

import (
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type moduleRoutes struct {
	uc     usecase.ModuleUseCase
	mapper mapper.Module
}

func newModuleRoutes(useCase usecase.ModuleUseCase, signingKey string) *chi.Mux {
	routes := moduleRoutes{uc: useCase}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingKey))
		r.Post("/", handler(routes.postModule).HTTP)
		r.Delete("/{id}", handler(routes.delete).HTTP)
	})
	return r
}

// @Tags module
// @Security ApiKeyAuth
// @Param       request body entity.ModulePost true "moduleRequest"
// @Success      201   {object} entity.ModuleDto       "created"
// @Router       /module [post]
func (r moduleRoutes) postModule(w http.ResponseWriter, request *http.Request) *appError {
	var modulePost entity.ModulePost
	if err := render.DecodeJSON(request.Body, &modulePost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	module := r.mapper.PostToModule(modulePost)
	module, err := r.uc.CreateModule(request.Context(), module)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	m := entity.ModuleDto{}
	dto.Map(&m, module)

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, m)
	return nil
}

// @Tags module
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /module/{id} [delete]
func (r moduleRoutes) delete(w http.ResponseWriter, request *http.Request) *appError {
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
