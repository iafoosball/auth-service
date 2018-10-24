package router

import (
	"github.com/gorilla/mux"
	"github.com/iafoosball/auth-service/handler"
)

const (
	AuthPath = "/auth"
)

func SetRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc(AuthPath + "/token", handler.JWT).Methods("POST")
	return r
}

