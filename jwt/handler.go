package jwt

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	AuthJWTTokenPath = "/auth/login"
)

// SetRoutes sets endpoints on a router
func SetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(AuthJWTTokenPath, handleJWT).Methods("POST")
	return r
}

func handleJWT(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	usernames, ok := vals["username"]
	if !ok || len(usernames) != 1 {
		w.WriteHeader(404)
	}
	username := usernames[0]

	token, err := IssueNew(username, "test")
	handleErr(err, w)

	payload, err := json.Marshal(token)
	handleErr(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// handleErr that requires generic handling (usually unexpected errors)
func handleErr(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
	}
}
