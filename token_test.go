package vkads

import (
	"testing"
)

func TestAuthCodeFlow_AgencyClientCredentialsGrantToken(t *testing.T) {
	token, err := acf.ClientCredentialsGrantToken(true)
	if err != nil {
		t.Error(err)
		return
	}

	if token.AccessToken != "" {
		t.Error("empty token")
		return
	}
}

func TestAuthCodeFlow_DeleteTokens(t *testing.T) {
	if err := acf.DeleteTokens("", 0); err != nil {
		t.Error(err)
		return
	}
}
