package social

import (
	"github.com/danilopolani/gocialite/structs"
	"golang.org/x/oauth2"
	"gopkg.in/danilopolani/gocialite.v0"
	"os"
)

const baseURL = "http://localhost:8001"

var gocial = gocialite.NewDispatcher()

// These credentials are old and revoked, please don't try to use them. Thanks!
// Google ClientID and ClientSecret for Oauth2 has to be supplied by the app environment
var providerSecrets = map[string]map[string]string{
	"facebook": {
		"clientID":     "111235093157542",
		"clientSecret": "2be6078a12fe8b1eecc89c6dea8b949c",
		"baseURL":      baseURL + "/auth/facebook/callback",
	},
	"google": {
		"clientID":     getEnv("GOOGLE_OAUTH2_CLIENT_ID", ""),
		"clientSecret": getEnv("GOOGLE_OAUTH2_CLIENT_SECRET", ""),
		"baseURL":      baseURL + "/auth/google/callback",
	},
}

var providerScopes = map[string][]string{
	"facebook": {},
	"google":   {"email", "profile", "openid"},
}

func RedirectURL(provider string) (string, error) {
	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]

	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["baseURL"],
		)

	return authURL, err
}

func ParseOauthResponse(state string, code string) (*structs.User, *oauth2.Token, error) {
	user, token, err := gocial.Handle(state, code)
	return user, token, err
}

// getEnv returns environmental variable called name, or fallback if empty
func getEnv(name string, fallback string) string {
	v, ok := os.LookupEnv(name)
	if ok {
		return v
	}
	return fallback
}