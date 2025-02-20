package jobposting

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kidusshun/hiring_assistant/service/auth"
	"github.com/kidusshun/hiring_assistant/utils"
)


type Handler struct {
	service JobPostingService
}


func NewHandler(service JobPostingService) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) RegisterRoutes(router	chi.Router) {
	router.With(auth.CheckBearerToken).Post("/job_posting", h.createJobPosting)
	router.With(auth.CheckBearerToken).Get("/job_posting", h.getJobPostings)
}


func (h *Handler) createJobPosting(w http.ResponseWriter, r *http.Request) {
	// get the user email from the context
	userEmail := r.Context().Value("userEmail").(string)

	// get the job posting details from the request body
	var request CreateJobPostingPayload
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	createdPosting, err := h.service.CreateJobPosting(userEmail, request)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdPosting)
	
}
func (h *Handler) getJobPostings(w http.ResponseWriter, r *http.Request) {
	// get the user email from the context
	userEmail, ok := r.Context().Value("userEmail").(string)
	if !ok {
		log.Println("could not get user email from context")
		utils.WriteError(w, http.StatusInternalServerError, errors.New("could not get user email from context"))
		return
	}

	// get the limit and offset from the query params
	limit := r.URL.Query().Get("limit")
	limitInt, err := ConvertStrToInt(limit)
	if err != nil {
		limitInt = 10
	}

	offset := r.URL.Query().Get("offset")
	offsetInt, err := ConvertStrToInt(offset)
	if err != nil {
		offsetInt = 0
	}
	

	jobPostings, err := h.service.GetJobPostings(userEmail, limitInt, offsetInt)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, jobPostings)
}