package auth

type GoogleUser struct {
	FirstName string	`json:"given_name"`
	LastName  string  	`json:"family_name"`
	Picture   string 	`json:"picture"`
	Email     string 	`json:"email"`
}