package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

const (
	LoginPath         = "/oauth/login"
	LogoutPath        = "/oauth/logout"
	ValidateTokenPath = "/oauth/verify"
)

// SetRoutes sets endpoints on a router
func SetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(LoginPath, handleLogin).Methods("POST")
	r.HandleFunc(LogoutPath, handleLogout).Methods("POST")
	r.HandleFunc(ValidateTokenPath, handleVerify).Methods("POST")
	return r
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	//"Authorization": "Basic Base64(username:password)"
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 && s[0] != "Basic" {
		log.Println("jwt/handler.go: Incorrect auth scheme")
		w.WriteHeader(401)
	} else {
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
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 && s[0] != "JWT" {
		log.Println("jwt/handler.go: Incorrect auth scheme")
		w.WriteHeader(401)
	}

	err := Revoke(s[1])

	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
		} else {
			handleErr(err, w)
		}
	} else {
		w.WriteHeader(200)
	}
}

func handleVerify(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 && s[0] != "JWT" {
		log.Println("jwt/handler.go: Incorrect auth scheme")
		w.WriteHeader(401)
	}

	ok, err := IsValid(s[1])
	handleErr(err, w)

	if !ok {
		log.Println("jwt/handler.go: JWT verification fail")
		w.WriteHeader(401)
	} else {
		w.WriteHeader(200)
	}
}

// handleErr that requires generic handling (usually unexpected errors)
func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println("jwt/handler.go: Error " + err.Error())
		w.WriteHeader(500)
	}
}
