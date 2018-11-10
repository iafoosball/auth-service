package main

import (
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/jwt"
	"github.com/iafoosball/auth-service/social"
	"net/http"
)

func main() {
	//rs256.MakeRSAKeysToDisk("test")
	r := mux.NewRouter()
	r = jwt.SetRoutes(r)
	r = social.SetRoutes(r)
	http.ListenAndServe(":8001", r)
}
