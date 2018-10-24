package handler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/crypto"
	"github.com/iafoosball/auth-service/model"
	"net/http"
	"time"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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

	//fmt.Print(validation.JWTisValid(token))
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

	token, err := rawToken.SignedString(crypto.ReadPrivateKey("./crypto/privateKey"))
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return token, nil
}


