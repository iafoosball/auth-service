package jwt

import (
	"github.com/iafoosball/auth-service/redis"
	"testing"
)

func TestIssueNew(t *testing.T) {
	jwt, err := IssueNew("test")
	if err != nil {
		t.Error(err)
	}

	r, err := redis.GET(jwt.ID)
	if r == nil || err != nil {
		t.Error(r, err)
	}

	r, err = redis.DEL(jwt.ID)
	if r.(int64) == 0 || err != nil {
		t.Error(r.(int64), err)
	}
}

func TestRevoke(t *testing.T) {
	jwt, err := IssueNew("test")
	if err != nil {
		t.Error(err)
	}

	err = Revoke(jwt.Token)
	if err != nil {
		t.Error(err)
	}

	r, err := redis.GET(jwt.ID)
	if r != nil || err != nil {
		t.Error(r, err)
	}
}

func TestIsValid(t *testing.T) {
	jwt, err := IssueNew("test")
	if err != nil {
		t.Error(err)
	}

	ok, err := IsValid(jwt.Token)
	if !ok || err != nil {
		t.Error(ok, err)
	}
}
