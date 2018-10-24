package validation

//import (
//	"fmt"
//	jwt "github.com/dgrijalva/jwt-go"
//	crypto"github.com/iafoosball/auth-service/crypto"
//)

//func JWTisValid(t string) bool {
//	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
//			fmt.Print("Wrong ALG")
//			return nil, nil
//		}
//		return crypto.GetPublicKey(), nil
//	})
//
//	if err == nil && token.Valid {
//		return true
//	}
//	fmt.Print(err)
//	return false
//}