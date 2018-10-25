package jwt

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/iafoosball/auth-service/rs256"
)

type Validator struct {
	PathToPublicKey string
	Password        string
}

// ValidateSignature allows to verify RSA signed JWT using public SHA256 key saved on local disk.
// Public key at location pathToPub must be PEM encoded and password protected using AES256 CBC algorithm.
func (v Validator) ValidateSignature(token string, pathToPub string, password string) bool {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Wrong ALG in JWT")
		}
		pubKey, err := rs256.ReadPublicKey(v.PathToPublicKey, v.Password)
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
