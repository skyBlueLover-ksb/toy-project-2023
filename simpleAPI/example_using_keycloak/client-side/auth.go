package main

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	myClientID = "myclient"
	//myClientSecret = "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"
	myClientSecret = "QrHp6EXKRMoIjJWjXbzpctAOvMcEZJlG"
)

const (
	CallBackURL         = "http://127.0.0.1:3030/auth/callback"
	UserInfoAPIEndpoint = "http://127.0.0.1:3000/users"

	//ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	//ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConfig *oauth2.Config
var myEndpoint = oauth2.Endpoint{
	AuthURL:   "http://localhost:8080/auth/realms/myrealm/protocol/openid-connect/auth",
	TokenURL:  "http://localhost:8080/auth/realms/myrealm/protocol/openid-connect/token",
	AuthStyle: oauth2.AuthStyleAutoDetect,
}

func init() {
	OAuthConfig = &oauth2.Config{
		ClientID:     myClientID,
		ClientSecret: myClientSecret,
		RedirectURL:  CallBackURL,
		//Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint: myEndpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
