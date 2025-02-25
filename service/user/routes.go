package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kidusshun/hiring_assistant/service/auth"
	"github.com/kidusshun/hiring_assistant/utils"
)

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{
		service:service,
	}
}



func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Post("/auth/google", h.googleAuth)
	router.With(auth.CheckBearerToken).Get("/user/me", h.getMe)
}


func (h*Handler) googleAuth(w http.ResponseWriter, r *http.Request) {
	var request LoginPayload
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := auth.VerifyGoogleToken(request.AccessToken)

	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	
	jwtToken, err := h.service.AddUser(user, request.AccessToken)

	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := LoginResponse{
		Token: jwtToken,
	}
	log.Println(response)
	utils.WriteJSON(w, http.StatusOK, response)
}


func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("userEmail").(string)

	user, err := h.service.GetMe(userEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteError(w, http.StatusNotFound, err)
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)

}