package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

// ----------------- CONFIG
var hp = strings.Split(os.Getenv("SERVICE_ADDR"), ":")
var h = hp[0]
var p = hp[1]
var basePath = "http://" + h + ":" + p

func TestSetRoutes(t *testing.T) {
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
		fmt.Printf("Expected 200 but was %v", resp.Status)
		t.Fail()
	}
	// ----------------- VERIFY OBTAINED JWT FROM BODY
	d := json.NewDecoder(resp.Body)
	var jt JWT
	err = d.Decode(&jt)
	if err != nil {
		t.Error(err)
	}
	st := jt.Token
	req, err = http.NewRequest("POST", basePath+"/oauth/verify", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Authorization", "JWT "+st)

	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "200 OK" {
		fmt.Printf("Expected 200 but was %v", resp.Status)
		t.Fail()
	}
	// ----------------- LOGOUT
	req, err = http.NewRequest("POST", basePath+"/oauth/logout", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Authorization", "JWT "+st)

	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "200 OK" {
		fmt.Printf("Expected 200 but was %v", resp.Status)
		t.Fail()
	}
	// ----------------- VERIFY REVOKED JWT
	req, err = http.NewRequest("POST", basePath+"/oauth/verify", nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Authorization", "JWT "+st)

	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "401 Unauthorized" {
		fmt.Printf("Expected 401 Not Found but was %v", resp.Status)
		t.Fail()
	}
}
