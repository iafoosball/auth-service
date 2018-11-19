package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"github.com/iafoosball/auth-service/rs256"
	"github.com/iafoosball/auth-service/social"
	"log"
	"net/http"
	"os"
)

func SERVICE_ADDR() string {
	var addr string
	if addr = os.Getenv("SERVICE_ADDR"); addr == "" {
		addr = "localhost:8001"
	}
	return addr
}

func main() {
	defer fmt.Println("auth-service exited")
	initRSA(rs256.PathToRSAPub)
	r := mux.NewRouter()
	r = jwt.SetRoutes(r)
	r = social.SetRoutes(r)
	http.ListenAndServe(SERVICE_ADDR(), r)
}

func initRSA(pathToKey string) {
	if _, err := os.Stat(pathToKey); err != nil {
		if os.IsNotExist(err) {
			if err := rs256.MakeRSAKeysToDisk("test"); err != nil {
				log.Fatal("Failed creating RSA keys: " + err.Error())
			}
		} else {
			log.Fatal("Failed creating RSA keys: " + err.Error())
		}
	}
}
