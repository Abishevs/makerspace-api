package config

import (
	"log"
	"github.com/joho/godotenv"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func OauthInit() oauth2.Config {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
    redirectURL := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

    return oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  redirectURL,
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint:     google.Endpoint,
    }
}
