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


func main() {
	defer fmt.Println("auth-service exited")
	initRSA(rs256.PathToRSAPub)
	r := mux.NewRouter()
	r = jwt.SetRoutes(r)
	r = social.SetRoutes(r)
	http.ListenAndServe(getEnv("SERVICE_ADDR", "localhost:8001"), r)
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

// getEnv returns environmental variable called name, or fallback if empty
func getEnv(name string, fallback string) string {
	v, ok := os.LookupEnv(name)
	if ok {
		return v
	}
	return fallback
}