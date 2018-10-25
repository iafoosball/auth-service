package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/rs256"
	"github.com/iafoosball/auth-service/model"
	"github.com/iafoosball/auth-service/jwt"
	"net/http"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWT struct {
	Token string `json:"token,omitempty"`
}

func JWT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// this does not work, plus kong does not support path parameters
	token, err := signedJWT(params["username"])
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}

	payload, err := json.Marshal(model.JWT{
		Token: token,
	})

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(500)
	}

	fmt.Print(jwt.JWTisValid(token+"lol", "test"))
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func signedJWT(username string) (string, error) {
	claims := JWTClaims{
		username,
		jwt.StandardClaims{
			Id: "priv",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//seq, err := rand.RuneSequence(30, rand.AlphaNum)
	//if err != nil {
	//	return "", err
	//}
	rsaKey, err := rs256.ReadPrivateKey("./id_rsa", "test")
	if err != nil {
		return "", err
	}
	token, err := rawToken.SignedString(rsaKey)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return token, nil
}


