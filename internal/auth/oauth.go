package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"exam-craft/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google OAuth configuration
var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  config.RedirectURL,
	ClientID:     config.GoogleClientID,
	ClientSecret: config.GoogleSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

// GoogleUserInfo holds the user's information obtained from Google
type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
}

// GoogleLogin redirects the user to Google's OAuth 2.0 authorization endpoint.
func GoogleLogin(c *gin.Context) {
	url := googleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles the callback from Google, exchanging the authorization code for an access token,
// and returning the user info. The handler will then use this information to handle user login or registration.
func GoogleCallback(c *gin.Context) (*GoogleUserInfo, error) {
	code := c.Query("code")
	if code == "" {
		return nil, errors.New("no code in callback request")
	}

	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
