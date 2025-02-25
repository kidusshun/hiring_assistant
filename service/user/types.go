package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/kidusshun/hiring_assistant/service/auth"
)

type UserStore interface {
	CreateUser(email, firstName, lastName, profilePictureURL, accessToken string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type UserService interface {
	AddUser(user *auth.GoogleUser, accessToken string) (string, error)
	GetMe(email string) (*User, error)
}



type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ProfilePictureURL string `json:"profile_picture_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginPayload struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}