package social

import (
	"github.com/danilopolani/gocialite/structs"
	"golang.org/x/oauth2"
	"gopkg.in/danilopolani/gocialite.v0"
)

const port = "8001"
const redirectURL = "http://localhost:8001"

var gocial = gocialite.NewDispatcher()

var providerSecrets = map[string]map[string]string{
	"facebook": {
		"clientID":     "111235093157542",
		"clientSecret": "2be6078a12fe8b1eecc89c6dea8b949c",
		"redirectURL":  redirectURL + "/auth/facebook/callback",
	},
	"google": {
		"clientID":     "659698836120-dosqs9rtc1p8eqcnl2qdmjpu1ujef9l9.apps.googleusercontent.com",
		"clientSecret": "6LaVZkuC01sQm9dq4a_laVfo",
		"redirectURL":  redirectURL + "/auth/google/callback",
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
			providerData["redirectURL"],
		)

	return authURL, err
}

func ParseOauthResponse(state string, code string) (*structs.User, *oauth2.Token, error) {
	user, token, err := gocial.Handle(state, code)
	return user, token, err
}
