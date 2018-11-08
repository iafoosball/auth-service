package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/iafoosball/auth-service/rand"
	"github.com/iafoosball/auth-service/redis"
	"github.com/iafoosball/auth-service/rs256"
	"time"
)

type Claims struct {
	Username string `json:"usr,omitempty"`
	jwt.StandardClaims
}

type JWT struct {
	Token string `json:"access_token,omitempty"`
	ID    string `json:"-"`
	TTL   int64  `json:"-"`
}

// newSigned creates new JWT token signed with RSA private key, protected with password.
// Private key must be PEM encoded and password protected using AES256 CBC algorithm.
func newSigned(username string, password string) (JWT, error) {
	jti, err := rand.RuneSequence(10, rand.AlphaUpperNum)
	now := time.Now().Add(24 * time.Hour).Unix()
	if err != nil {
		return JWT{}, err
	}
	c := Claims{
		username,
		jwt.StandardClaims{
			Id:        string(jti),
			ExpiresAt: now,
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	rsaKey, err := rs256.ReadPrivateKey(password)
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}
	tokenString, err := rawToken.SignedString(rsaKey)
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}
	token := JWT{
		Token: tokenString,
		ID:    string(jti),
		TTL:   now,
	}
	return token, nil
}

// IssueNew JWT token from auth-service and returns JSON payload containing the token
func IssueNew(username string, password string) (JWT, error) {
	token, err := newSigned(username, password)
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}

	_, err = redis.Perform("SET", token.ID, username)
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}

	return token, err
}

func Revoke() {

}

func Validate() {

}
