package main

import (
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/router"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	router.SetRoutes(r)
	http.ListenAndServe(":8001", r)
}
