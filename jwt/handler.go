package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

const (
	LoginPath = "/oauth/login"
	LogoutPath = "/oauth/logout"
	ValidateTokenPath = "/oauth/validate"
)

// SetRoutes sets endpoints on a router
func SetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(LoginPath, handleLogin).Methods("POST")
	r.HandleFunc(LogoutPath, handleLogout).Methods("POST")
	r.HandleFunc(ValidateTokenPath, handleValidation).Methods("POST")
	return r
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	//"Authorization": "Basic Base64(username:password)"
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 && s[0] != "Basic" {
		w.WriteHeader(401)
	}

	d, err := base64.StdEncoding.DecodeString(s[1])
	handleErr(err, w)

	c := strings.Split(string(d), ":")

	token, err := IssueNew(c[0])
	handleErr(err, w)

	payload, err := json.Marshal(token)
	handleErr(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 && s[0] != "JWT" {
		w.WriteHeader(401)
	}

	err := Revoke(s[1])
	if err.Error() == "not found" {
		w.WriteHeader(404)
	} else {
		handleErr(err, w)
	}

	w.WriteHeader(200)
}

func handleValidation(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	v, err := IsValid(s[1])
	handleErr(err, w)

	if len(s) != 2 && s[0] != "JWT" && !v {
		w.WriteHeader(401)
	}

	w.WriteHeader(200)
}

// handleErr that requires generic handling (usually unexpected errors)
func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
	}
}
