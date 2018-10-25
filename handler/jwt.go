package handler

import (
	"encoding/json"
	"fmt"
	"github.com/iafoosball/auth-service/jwt"
	"net/http"
)

type Token struct {
	Token string `json:"access_token,omitempty"`
}

func JWT(w http.ResponseWriter, r *http.Request) {
	// here we get username from Basic Auth request coming from Kong
	vals := r.URL.Query()

	usernames, ok := vals["username"]
	if !ok || len(usernames) != 1 {
		w.WriteHeader(404)
	}
	username := usernames[0]

	token, err := jwt.NewSigned(username, "./id_rsa", "test")
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}

	payload, err := json.Marshal(Token{
		Token: "JWT " + token,
	})

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}

	fmt.Print((jwt.ValidateSignature(token, "./id_rsa_pub","test")))
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}


