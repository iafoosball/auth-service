package main

import (
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"github.com/iafoosball/auth-service/social"
	"net/http"
	"os"
)

func main() {
	// run this once on clean build to generate RSA keys
	//rs256.MakeRSAKeysToDisk("test")

	r := mux.NewRouter()
	r = jwt.SetRoutes(r)
	r = social.SetRoutes(r)
	http.ListenAndServe("localhost:" + os.Getenv("AUTH_PORT"), r)
}
