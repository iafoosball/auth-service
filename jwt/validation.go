package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iafoosball/auth-service/rs256"
)

// ValidateSignature allows to verify RSA signed JWT using public SHA256 key saved on local disk.
// Public key must be PEM encoded and password protected using AES256 CBC algorithm.
func ValidateSignature(token string, pathToFile string, password string) bool {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Wrong ALG in JWT")
		}
		pubKey, err := rs256.ReadPublicKey(pathToFile, password)
		if err != nil {
			return nil, err
		}
		return pubKey, nil
	})
	if err == nil && parsed.Valid {
		return true
	}
	return false
}