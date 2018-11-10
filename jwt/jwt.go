package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/iafoosball/auth-service/rand"
	"github.com/iafoosball/auth-service/redis"
	"github.com/iafoosball/auth-service/rs256"
	"time"
)

// these should definitely be in some sort of secret datastore
const (
	PrivKeyPassphrase = "test"
	PubKeyPassphrase = "test"
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

// IssueNew JWT token from auth-service and returns JSON payload containing the token
func IssueNew(username string) (JWT, error) {
	token, err := newSigned(username)
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}

	r, err := redis.SET(token.ID, username, token.TTL)
	if r == nil {
		panic("token with that ID was already registered")
	}
	if err != nil {
		fmt.Print(err)
		return JWT{}, err
	}

	return token, err
}
// Revoke token from redis db.
func Revoke(token string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return err
	}

	tid := claims["ID"]
	r, err := redis.DEL(tid.(string))
	if r.(int) == 0 {
		return errors.New("not found")
	}
	return nil
}

// IsValid verifies whether token is properly signed and issued by auth-service.
func IsValid(token string) (bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, keyFunc)
	if err != nil {
		return false, err
	}

	tid := claims["ID"]
	r, err := redis.DEL(tid.(string))
	if err != nil {
		return false, err
	}

	if r != nil {
		return true, err
	}

	return false, err
}

// newSigned creates new JWT token signed with RSA private key, protected with password.
// Private key must be PEM encoded and password protected using AES256 CBC algorithm.
func newSigned(username string) (JWT, error) {
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
	rsaKey, err := rs256.ReadPrivateKey(PrivKeyPassphrase)
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

// keyFunc for validating JWT signature against RSA public key
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, errors.New("Wrong ALG in JWT")
	}
	pubKey, err := rs256.ReadPublicKey(PubKeyPassphrase)
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}