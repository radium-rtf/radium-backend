package v1

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type pageRoutes struct {
	uc     usecase.PageUseCase
	mapper mapper.Page
}

func newPageRoutes(useCase usecase.PageUseCase, signingString string) *chi.Mux {
	routes := pageRoutes{uc: useCase}

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(authToken(signingString))
		r.Get("/{id}", handler(routes.getById).HTTP)
	})

	r.Group(func(r chi.Router) {
		r.Use(authRequired(signingString))
		r.Post("/", handler(routes.post).HTTP)
		r.Delete("/{id}", handler(routes.delete).HTTP)
	})

	return r
}

// @Tags page
// @Security ApiKeyAuth
// @Param       request body entity.PagePost true "создание"
// @Success      201   {object} entity.PageDto "ok"
// @Router      /page [post]
func (r pageRoutes) post(w http.ResponseWriter, request *http.Request) *appError {
	var post entity.PagePost
	if err := render.DecodeJSON(request.Body, &post); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}
	page := r.mapper.NewPostToPage(&post)
	pageDto, err := r.uc.CreatePage(request.Context(), page)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, pageDto)
	return nil
}

// @Tags page
// @Security ApiKeyAuth
// @Param        id   path     string  true  "page id"
// @Success 200 {object} entity.PageDto "ok"
// @Router /page/{id} [get]
func (r pageRoutes) getById(w http.ResponseWriter, request *http.Request) *appError {
	id, err := uuid.Parse(chi.URLParam(request, "id"))
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	userId, ok := request.Context().Value("userId").(uuid.UUID)
	var pageDto *entity.PageDto
	if !ok {
		pageDto, err = r.uc.GetByID(request.Context(), id, nil)
	} else {
		pageDto, err = r.uc.GetByID(request.Context(), id, &userId)
	}

	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusOK)
	render.JSON(w, request, pageDto)

	return nil
}

// // @Tags slide
// // @Security ApiKeyAuth
// // @Param        courseId   path      integer  true  "course id"
// // @Param        moduleNameEng   path     string  true  "moduleName"
// // @Success      201   {object} entity.ModuleSlides "ok"
// // @Router      /slide/{courseId}/{moduleNameEng} [get]
// func (r slideRoutes) getSlides(w http.ResponseWriter, request *http.Request) *appError {
// 	courseId, err := strconv.Atoi(chi.URLParam(request, "courseId"))
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	moduleNameEng := chi.URLParam(request, "moduleNameEng")
// 	slideRequest := entity.SlidesRequest{ModuleNameEng: moduleNameEng, CourseId: uint(courseId)}
// 	slide, err := r.uc.GetSlides(request.Context(), slideRequest)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}
// 	render.Status(request, http.StatusOK)
// 	render.JSON(w, request, slide)
// 	return nil
// }

// @Tags page
// @Security ApiKeyAuth
// @Accept json
// @Param        id   path      string  true  "id"
// @Param   is_soft     query     boolean    false  "по умолчанию soft"
// @Success 200
// @Router /page/{id} [delete]
func (r pageRoutes) delete(w http.ResponseWriter, request *http.Request) *appError {
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
