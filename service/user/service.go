package user

import (
	"database/sql"

	"github.com/kidusshun/hiring_assistant/service/auth"
)

type Service struct {
	store UserStore
}

func NewService(store UserStore) *Service {
	return &Service{
		store: store,
	}
}


func (s *Service) AddUser(user *auth.GoogleUser) (string, error) {
	// check if user exists
	storedUser, err := s.store.GetUserByEmail(user.Email)

	if storedUser == nil {
		if err == sql.ErrNoRows {
			createdUser, err := s.store.CreateUser(user.Email, user.FirstName, user.LastName, user.Picture)

			if err != nil {
				return "", err
			}
			jwtToken, err := auth.GenerateJWT(createdUser.Email)

			if err != nil {
				return "", err
			}
			return jwtToken, nil
		} else {
			return "", err
		}
	} else {
		jwtToken, err := auth.GenerateJWT(storedUser.Email)

		if err != nil {
			return "", err
		}
		return jwtToken, nil
	}
}

func (s *Service) GetMe(email string) (*User, error) {
	user, err := s.store.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}