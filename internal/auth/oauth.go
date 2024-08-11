package auth

import (
	"context"
	"encoding/json"
	"exam-craft/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Google OAuth configuration
var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  config.RedirectURL,
	ClientID:     config.GoogleClientID,
	ClientSecret: config.GoogleSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

// GoogleUserInfo holds the user's information obtained from Google
type GoogleUserInfo struct {
	Email string `json:"email"`
}

// GetGoogleUserInfo fetches user info from Google using the provided code
func GetGoogleUserInfo(code string) (*GoogleUserInfo, error) {
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
