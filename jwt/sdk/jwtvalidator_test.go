package sdk

import (
	"encoding/json"
	"github.com/iafoosball/auth-service/jwt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestJWTValidator_ValidateToken(t *testing.T) {
	// ----------------- CONFIG
	hp := strings.Split(os.Getenv("SERVICE_ADDR"), ":")
	h := hp[0]
	p := hp[1]
	basePath := "http://"+h+":"+p

	// ----------------- LOGIN
	client := http.DefaultClient
	req, err := http.NewRequest("POST", basePath+"/oauth/login", nil)
	if err != nil {
		t.Error(err)
	}
	req.SetBasicAuth("test", "test1234")

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "200 OK" {
		t.Errorf("Expected 200 but was %v", resp.Status)
	}

	d := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var j jwt.JWT
	err = d.Decode(&j)
	if err != nil {
		t.Error(err)
	}

	// ----------------- VERIFY TOKEN
	pInt, err := strconv.Atoi(p)
	if err != nil {
		t.Error(err)
	}
	v := JWTValidator{
		Protocol: "http",
		Hostname: h,
		Port: pInt,
	}
	if ok, _ := v.ValidateToken(j.Token); !ok {
		t.Error("Validation failed on new token")
	}

	// ----------------- LOGOUT
	req, err = http.NewRequest("POST", basePath+"/oauth/logout", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Authorization", "JWT "+j.Token)

	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "200 OK" {
		t.Errorf("Expected 200 but was %v", resp.Status)
	}
}
