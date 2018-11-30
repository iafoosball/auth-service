package sdk

import (
	"errors"
	"github.com/iafoosball/auth-service/jwt"
	"net/http"
)

// JWTValidator constructor configures the connection with auth-service in order to validate incoming tokens in
// HTTP request headers. All fields are required
type JWTValidator struct {
	Protocol string
	Hostname string
	Port     int
}

// ValidateToken against remote auth-service.
func (v JWTValidator) ValidateToken(token string) (bool, error) {
	url := v.Protocol + "://" + v.Hostname + ":" + string(v.Port) + jwt.ValidateTokenPath

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "JWT "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code == 200 {
		return true, nil
	} else {
		return false, errors.New("Http Status Code Error: " + string(code))
	}
}
