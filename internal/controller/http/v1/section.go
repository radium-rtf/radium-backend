package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase"
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

			r.Post("/question", handler(routes.postQuestion).HTTP)
			r.Post("/question/answer", handler(routes.postQuestionAnswer).HTTP)

			r.Post("/choice", handler(routes.postChoice).HTTP)
			r.Post("/choice/answer", handler(routes.postChoiceAnswer).HTTP)

			r.Post("/multichoice", handler(routes.postMultiChoice).HTTP)
			r.Post("/multichoice/answer", handler(routes.postMultiChoiceAnswer).HTTP)
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

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionQuestionPost true "создание вопросительной секции"
// @Success      201   {object} entity.SectionQuestionDto "ok"
// @Router      /section/question [post]
func (r sectionRoutes) postQuestion(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionQuestionPost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	section, err := r.uc.CreateQuestion(request.Context(), sectionPost)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionQuestionAnswerPost true "ответ на вопрос"
// @Success      201   {object} entity.SectionQuestionAnswerDto "ok"
// @Router      /section/question/answer [post]
func (r sectionRoutes) postQuestionAnswer(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionQuestionAnswerPost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	userId := request.Context().Value("userId").(string)
	section, err := r.uc.CreateQuestionAnswer(request.Context(), sectionPost, userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionChoicePost true "создание секции c вариантом ответа"
// @Success      201   {object} entity.SectionChoiceDto "ok"
// @Router      /section/choice [post]
func (r sectionRoutes) postChoice(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionChoicePost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	section, err := r.uc.CreateChoice(request.Context(), sectionPost)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionChoiceAnswerPost true "ответ на задание с вариантом ответа"
// @Success      201   {object} entity.SectionChoiceAnswerDto "ok"
// @Router      /section/choice/answer [post]
func (r sectionRoutes) postChoiceAnswer(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionChoiceAnswerPost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	userId := request.Context().Value("userId").(string)
	section, err := r.uc.CreateChoiceAnswer(request.Context(), sectionPost, userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionMultiChoicePost true "создание секции c несколькими вариантами ответа"
// @Success      201   {object} entity.SectionMultiChoiceDto "ok"
// @Router      /section/multichoice [post]
func (r sectionRoutes) postMultiChoice(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionMultiChoicePost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	section, err := r.uc.CreateMultiChoice(request.Context(), sectionPost)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}

// @Tags section
// @Security ApiKeyAuth
// @Param       request body entity.SectionMultiChoiceAnswerPost true "ответ на задание с несколькими вариантами ответа"
// @Success      201   {object} entity.SectionMultiChoiceAnswerDto "ok"
// @Router      /section/multichoice/answer [post]
func (r sectionRoutes) postMultiChoiceAnswer(w http.ResponseWriter, request *http.Request) *appError {
	var sectionPost entity.SectionMultiChoiceAnswerPost
	if err := render.DecodeJSON(request.Body, &sectionPost); err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	userId := request.Context().Value("userId").(string)
	section, err := r.uc.CreateMultiChoiceAnswer(request.Context(), sectionPost, userId)
	if err != nil {
		return newAppError(err, http.StatusBadRequest)
	}

	render.Status(request, http.StatusCreated)
	render.JSON(w, request, section)
	return nil
}
