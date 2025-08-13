package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	_, err := GetAPIKey(http.Header{
		"a": []string{"1"},
	})

	if err != ErrNoAuthHeaderIncluded {
		t.Fail()
	}

	_, err = GetAPIKey(http.Header{
		"Authorization": []string{"1"},
	})

	if err.Error() != "malformed authorization header" {
		t.Fail()
	}

	key, _ := GetAPIKey(http.Header{
		"Authorization": []string{"ApiKey ###"},
	})

	if key != "###" {
		t.Fail()
	}
}
