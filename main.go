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

	//SET <OK, nil>
	r, err := redis.SET("lodl", "bet", 1000)
	fmt.Println(r)
	fmt.Println(err)
	//returns []uint8, <nil if not found>
	//r, _ := redis.GET("lodl")
	//if r != nil {
	//	fmt.Println(string(r.([]byte)))
	//}
	// 0 is not found, 1 if deleted
	//r, _ = redis.DEL("lodl")
	//fmt.Println(r)
	//r, _ = redis.GET("lol")
	//fmt.Println(r)
	}
