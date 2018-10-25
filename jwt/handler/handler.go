package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"net/http"
)

const (
	AuthJWTTokenPath = "/auth/token"
	PathToRSAPrivateKey = "./id_rsa"
)

type Token struct {
	Token string `json:"access_token,omitempty"`
}

// SetRoutes lol
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

	token, err := jwt.NewSigned(username, PathToRSAPrivateKey, "test")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}
	payload, err := json.Marshal(Token{
		Token: "handleJWT " + token,
	})
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}


