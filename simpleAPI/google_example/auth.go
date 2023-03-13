package main

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	myGoogleClientID = "206115972812-mm4hv7h8vaareea0kte844925pgbun90.apps.googleusercontent.com"
	//myGoogleClientSecret = "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"
	myGoogleClientSecret = "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"
)

const (
	CallBackURL         = "http://127.0.0.1:8080/auth/callback"
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConfig *oauth2.Config

func init() {
	OAuthConfig = &oauth2.Config{
		ClientID:     myGoogleClientID,
		ClientSecret: myGoogleClientSecret,
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConfig.AuthCodeURL(state)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
