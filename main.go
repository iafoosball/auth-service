package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"github.com/iafoosball/auth-service/rs256"
	"github.com/iafoosball/auth-service/social"
	"net/http"
	"os"
)

func main() {
	// run this once on clean build to generate RSA keys
	if err := rs256.MakeRSAKeysToDisk("test"); err != nil {
		fmt.Println(err)
	}
	defer fmt.Println("Exited....")
	r := mux.NewRouter()
	r = jwt.SetRoutes(r)
	r = social.SetRoutes(r)
	var addr string
	if addr = os.Getenv("AUTH_ADDR"); addr == "" {
		addr = "localhost:8001"
	}
	http.ListenAndServe(addr, r)
}
