package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/iafoosball/auth-service/rand"
	"github.com/iafoosball/auth-service/rs256"
	"time"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Sign(username string) (string, error) {
	jti, err := rand.RuneSequence(10, rand.AlphaUpperNum)
	if err != nil {
		return "", err
	}
	c := Claims{
		username,
		jwt.StandardClaims{
			Id: string(jti),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	rsaKey, err := rs256.ReadPrivateKey("./id_rsa", "test")
	if err != nil {
		return "", err
	}
	token, err := rawToken.SignedString(rsaKey)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return token, nil
}
