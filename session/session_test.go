package session

import (
	"log"
	"os"
	"testing"
)

func TestRippleAccessToken(t *testing.T) {
	if os.Getenv("RIPPLE_TESTER_CLIENT_ID_PROD") == "" ||
		os.Getenv("RIPPLE_TESTER_CLIENT_SECRET_PROD") == "" {
		return
	}

	s := New(
		WithClientId(os.Getenv("RIPPLE_TESTER_CLIENT_ID_PROD")),
		WithClientSecret(os.Getenv("RIPPLE_TESTER_CLIENT_SECRET_PROD")),
	)

	token, err := s.AccessToken()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(token)
}
