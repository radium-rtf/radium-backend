package v1

import (
	"encoding/json"
	"errors"
	"github.com/radium-rtf/radium-backend/pkg/mapper"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
)

type sectionRoutes struct {
	uc     usecase.SectionUseCase
	mapper mapper.Section
}

func newSectionRoutes(h chi.Router, useCase usecase.SectionUseCase, signingString string) {
	routes := sectionRoutes{uc: useCase}

	h.Route("/section", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(authRequired(signingString))
			r.Post("/", handler(routes.createSection).HTTP)
			// r.Post("/question/answer", handler(routes.postQuestionAnswer).HTTP)
			// r.Post("/choice/answer", handler(routes.postChoiceAnswer).HTTP)
			// r.Post("/multichoice/answer", handler(routes.postMultiChoiceAnswer).HTTP)
		})
	})
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

// // @Tags section
// // @Security ApiKeyAuth
// // @Param       request body entity.SectionQuestionAnswerPost true "ответ на вопрос"
// // @Success      201   {object} entity.SectionQuestionAnswerDto "ok"
// // @Router      /section/question/answer [post]
// func (r sectionRoutes) postQuestionAnswer(w http.ResponseWriter, request *http.Request) *appError {
// 	var sectionPost entity.SectionQuestionAnswerPost
// 	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	userId := request.Context().Value("userId").(string)
// 	section, err := r.uc.CreateQuestionAnswer(request.Context(), sectionPost, userId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	render.Status(request, http.StatusCreated)
// 	render.JSON(w, request, section)
// 	return nil
// }

// // @Tags section
// // @Security ApiKeyAuth
// // @Param       request body entity.SectionChoiceAnswerPost true "ответ на задание с вариантом ответа"
// // @Success      201   {object} entity.SectionChoiceAnswerDto "ok"
// // @Router      /section/choice/answer [post]
// func (r sectionRoutes) postChoiceAnswer(w http.ResponseWriter, request *http.Request) *appError {
// 	var sectionPost entity.SectionChoiceAnswerPost
// 	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	userId := request.Context().Value("userId").(string)
// 	section, err := r.uc.CreateChoiceAnswer(request.Context(), sectionPost, userId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	render.Status(request, http.StatusCreated)
// 	render.JSON(w, request, section)
// 	return nil
// }

// // @Tags section
// // @Security ApiKeyAuth
// // @Param       request body entity.SectionMultiChoiceAnswerPost true "ответ на задание с несколькими вариантами ответа"
// // @Success      201   {object} entity.SectionMultiChoiceAnswerDto "ok"
// // @Router      /section/multichoice/answer [post]
// func (r sectionRoutes) postMultiChoiceAnswer(w http.ResponseWriter, request *http.Request) *appError {
// 	var sectionPost entity.SectionMultiChoiceAnswerPost
// 	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	userId := request.Context().Value("userId").(string)
// 	section, err := r.uc.CreateMultiChoiceAnswer(request.Context(), sectionPost, userId)
// 	if err != nil {
// 		return newAppError(err, http.StatusBadRequest)
// 	}

// 	render.Status(request, http.StatusCreated)
// 	render.JSON(w, request, section)
// 	return nil
// }
