package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetUpConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "121074428867-46fsbq8dhran0hnd5hant4pptkvvlvoi.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8000/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
