package social

import (
	"net/http"
	"testing"
)

// TODO: test callbacks
func TestSetRoutes(t *testing.T) {
	testRedirect("facebook", t)
	testRedirect("google", t)
}

func testRedirect(provider string, t *testing.T) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:8001/oauth/"+provider, nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 302 || resp.Body != nil {
		t.Error(resp, err)
	}
}
