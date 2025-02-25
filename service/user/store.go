package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*User, error) {
	rows := s.db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	u := new(User)
	u, err := ScanRowToUser(rows)
	if err != nil {
		return nil, err
	}
	return u, nil

}

func (s *Store) GetUserByID(id uuid.UUID) (*User, error) {
	rows := s.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	u := new(User)
	u, err := ScanRowToUser(rows)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil

}

func (s *Store) CreateUser(email, firstName, lastName, profilePictureURL, accessToken string) (*User, error) {
	row := s.db.QueryRow("INSERT INTO users (email, access_token, first_name, last_name, profile_picture_url) VALUES ($1, $2, $3, $4, $5) RETURNING id, email,access_token, first_name, last_name, profile_picture_url, created_at, updated_at", email,accessToken, firstName, lastName, profilePictureURL)
	
	createdUser, err := ScanRowToUser(row)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func ScanRowToUser(rows *sql.Row) (*User, error) {

	user := new(User)
	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.AccessToken,
		&user.FirstName,
		&user.LastName,
		&user.ProfilePictureURL,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}