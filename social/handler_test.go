package social

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

// ----------------- CONFIG
var hp = strings.Split(os.Getenv("SERVICE_ADDR"), ":")
var h = hp[0]
var p = hp[1]
var basePath = "http://"+h+":"+p

// TODO: test callbacks
func TestSetRoutes(t *testing.T) {
	testRedirect("facebook", t)
	testRedirect("google", t)
}

func testRedirect(provider string, t *testing.T) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", basePath+"/oauth/"+provider, nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 302 || resp.Body != nil {
		t.Error(resp, err)
	}
}
