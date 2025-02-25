package resumes

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
	ResumesService ResumeService
}


func NewHandler(service ResumeService) *Handler {
	return &Handler{
		ResumesService: service,
	}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.With(auth.CheckBearerToken).Post("/resumes", h.StoreResumes)
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
