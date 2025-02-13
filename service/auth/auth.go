package auth

import (
	"github.com/kidusshun/hiring_assistant/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     config.GoogleClient.GoogleClientID,
	ClientSecret: config.GoogleClient.GoogleClientSecret,
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}