package handlers

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"context"
	"log"
	"encoding/base64"
	"crypto/rand"
	"os"
	"time"
	"github.com/joho/godotenv"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = initializeGoogleOauthConfig()

func initializeGoogleOauthConfig() oauth2.Config {
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
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
        Endpoint:     google.Endpoint,
    }
}
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleLogin(c *gin.Context) {

		// log.Println("Google OAuth Config:", googleOauthConfig)

        url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
		// log.Println("Generated Auth URL:", url)

        c.Redirect(http.StatusFound, url)
    }

//func oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
//	oauthState := generateStateOauthCookie(w)

	/*
	AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	validate that it matches the the state query parameter on your redirect callback.
	*/
//	u := googleOauthConfig.AuthCodeURL(oauthState)
//	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
//}

func OauthGoogleCallback(c *gin.Context) {
        code := c.Query("code")
        token, err := googleOauthConfig.Exchange(c, code)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken})
    }


func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}