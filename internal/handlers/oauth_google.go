package handlers

import (
	"golang.org/x/oauth2"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"encoding/base64"
	"crypto/rand"
	// "os"
	"time"
	
	"encoding/json"
	"makerspace-api/models"
	"makerspace-api/config"
)
/*
   OAUTH Flow
1. Intialiszation
1.1  USER on the front end presses login button. And triggers /auth/login url.
1.2  Backend Generated URL and sends as json to Front End for user redirection to Googles login page

2. UserData fetching 
2.1 Frontend user is redirected to Callback URL /api/sessions/oauth/google 
2.2 Front end has to pass to the  Backend the authorization code from callback url 
2.3 Backend exchanges auth code for an access token
2.3.1 IMPLEMENT SECUIRTY MESUARES AGAINST CSRF ATTACKS!!! 
2.4 Backend fetch User data such as: 
	// GoogleUserData
	ID	    string `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"name"`
	FirstName   string `json:"given_name"`
	LastName    string `json:"family_name"`
	Picture     string `json:"picture"`

	from google API with the access token. 

2.5 Also backend have to fetch refresh token etc
2.6 Backend stores userData in the database  for role based access. 
2.7 Backend sends back UserData with assigned role and an access token

*/

// Sets credentials for google oauth
var googleOauthConfig = config.OauthInit()


const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleLogin(c *gin.Context) {


	oauthState := generateStateOauthCookie(c)
	fmt.Printf("State cookie: %s", oauthState)

        url := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)

	c.JSON(http.StatusOK, gin.H{"google_oauth_url": url})
	// c.Redirect(http.StatusTemporaryRedirect, url)
    }

//func oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie

	/*
	AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	validate that it matches the the state query parameter on your redirect callback.
	*/
//	u := googleOauthConfig.AuthCodeURL(oauthState)
//	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
//}

func OauthGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	fmt.Printf("State cookie %s",state)
	/*
	AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
	validate that it matches the the state query parameter on your redirect callback.
	*/
	

	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	// Use the token.AccessToken to fetch user info from Google's API
	userInfo, err := getUserDataFromGoogle(token.AccessToken) // token.AccessToken
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	// Store user info in your backend database 
	// ...

	c.JSON(http.StatusOK, gin.H{
		"user_info": userInfo,
		"access_token": token.AccessToken,})

}

func getUserDataFromGoogle(accessToken string) (models.GoogleUserInfo, error) {
    var userInfo models.GoogleUserInfo

    url := oauthGoogleUrlAPI + accessToken

    response, err := http.Get(url)
    if err != nil {
        return userInfo, fmt.Errorf("failed getting user info: %s", err.Error())
    }
    defer response.Body.Close()

    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return userInfo, fmt.Errorf("failed read response: %s", err.Error())
    }

    // Unmarshal the JSON response into the userInfo struct
    err = json.Unmarshal(contents, &userInfo)
    if err != nil {
        return userInfo, fmt.Errorf("failed to unmarshal JSON: %s", err.Error())
    }

    return userInfo, nil
}

// Rework this as state cookies needs to be sent together with the google login URL, Google will then send back it as state query
// In the callback URL
// Verify it at the callback route
func generateStateOauthCookie(c *gin.Context) string {
    expiration := time.Now().Add(20 * time.Minute)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)

    cookie := http.Cookie{
        Name:     "oauthstate",
        Value:    state,
        Expires:  expiration,
        HttpOnly: true,
    }

    http.SetCookie(c.Writer, &cookie)

    return state
}

