package resumes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kidusshun/hiring_assistant/service/auth"
	"github.com/kidusshun/hiring_assistant/utils"
)


type Handler struct {
	ResumesService ResumeService
}


func NewHandler(service ResumeService) *Handler {
	return &Handler{
		ResumesService: service,
	}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.With(auth.CheckBearerToken).Post("/resumes", h.StoreResumes)
	router.With(auth.CheckBearerToken).Get("/resumes/{jobPostingID}", h.getResumes)
}


func(h *Handler) StoreResumes(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail").(string)

	var payload CreateResumesPayload;
	err := json.NewDecoder(r.Body).Decode(&payload)

	
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	err = utils.Validate.Struct(payload)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing required fields in payload"))
		return
	}


	resumes, err := h.ResumesService.StoreResumeService(userEmail, payload)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resumes)
}


func(h *Handler) getResumes(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail").(string)
	jobPostingID := chi.URLParam(r, "jobPostingID")

	log.Println(jobPostingID)

	if jobPostingID == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing jobPostingID"))
		return
	}

	jobPostingUUID, err := uuid.Parse(jobPostingID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid jobPostingID format"))
		return
	}

	resumes, err := h.ResumesService.GetResumesService(userEmail, jobPostingUUID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resumes)
}