package main

import (
	//"github.com/gorilla/mux"
	//"github.com/iafoosball/auth-service/jwt"
	//"github.com/iafoosball/auth-service/rs256"
	//"github.com/iafoosball/auth-service/social"
	//"net/http"
	"fmt"
	"github.com/iafoosball/auth-service/redis"
)

const port = "8001"
const redirectUrl = "http://localhost:8001"

func main() {
	//rs256.MakeRSAKeysToDisk("test", "./rs256/")
	//r := mux.NewRouter()
	//r = jwt.SetRoutes(r)
	//r = social.SetRoutes(r)
	//http.ListenAndServe(":" + port, r)

	// SET, GET, DEL
	r, err := redis.Perform("SET", "lol", "lol")
	fmt.Println(r)
	fmt.Println(err)
}
