package main

import (
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt/handler"
	"github.com/iafoosball/auth-service/rs256"
	"net/http"
)

func main() {
	makeRSAKeysToDisk("test", "./")
	r := mux.NewRouter()
	r = handler.SetRoutes(r)
	http.ListenAndServe(":8001", r)
}

func makeRSAKeysToDisk(password string, pathToDir string) error {
	privateKey, err := rs256.GeneratePrivateKey(1024)
	if err != nil {
		return err
	}
	keyPEM, err := rs256.PrivateKeyToPEM(privateKey, password)
	if err != nil {
		return err
	}
	rs256.WriteKeyToFile(keyPEM, pathToDir + "/id_rsa")
	publicKey, err := rs256.GeneratePublicKeyPEM(privateKey, password)
	if err != nil {
		return err
	}
	rs256.WriteKeyToFile(publicKey, pathToDir + "/id_rsa_pub")
	return err
}










