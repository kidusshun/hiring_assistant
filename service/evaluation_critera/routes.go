package evaluationcritera

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kidusshun/hiring_assistant/service/auth"
	"github.com/kidusshun/hiring_assistant/utils"
)


type Handler struct {
	service EvaluationCriteriaService
}


func NewHandler(service EvaluationCriteriaService) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) RegisterRoutes(router chi.Router) {
	router.With(auth.CheckBearerToken).Post("/evaluation_critera", h.createEvaluationCriteria)
	router.With(auth.CheckBearerToken).Get("/evaluation_critera", h.getEvaluationCriteria)
}


func (h *Handler) createEvaluationCriteria(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email").(string)

	var request []CreateCriteriaPayload
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	evalutionCriterias, err := h.service.AddEvaluationCriteria(email, request)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, evalutionCriterias)
}


func (h *Handler) getEvaluationCriteria(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("email").(string)

	var payload GetEvaluationCriteriaPayload

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	evaluationCriterias, err := h.service.GetEvaluationCriteria(userEmail, payload)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, evaluationCriterias)

}