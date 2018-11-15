package social

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"log"
	"net/http"
)

const (
	SocialLoginPath         = "/oauth/{provider}"
	SocialLoginCallbackPath = "/oauth/{provider}/callback"
)

// SetRoutes sets endpoints on a router
func SetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(SocialLoginPath, redirectHandler).Methods("GET")
	r.HandleFunc(SocialLoginCallbackPath, callbackHandler).Methods("GET")
	return r
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	authURL, err := RedirectURL(provider)
	handleErr(err, w)

	http.Redirect(w, r, authURL, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	state := vals["state"]
	code := vals["code"]

	user, _, err := ParseOauthResponse(state[0], code[0])
	handleErr(err, w)

	// The following needs to be done here:
	// <get email from provider
	// <Get users/{email} (need to use pagination), return username
	// <if exists, put username in JWT
	// <otherwise, get credentials input from user
	// <create user in users-service
	// When created:

	// return JWT with valid username (currently user has to be registered with same username as in google)
	token, err := jwt.IssueNew(user.Email)
	handleErr(err, w)

	payload, err := json.Marshal(token)
	handleErr(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// handleErr that requires generic handling (usually unexpected errors)
func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println("social/handler.go: Error " + err.Error())
		w.WriteHeader(500)
	}
}
