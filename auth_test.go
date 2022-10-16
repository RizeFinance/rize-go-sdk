package rize

import (
	"context"
	"net/http"
	"testing"
)

// Complete AuthTokenResponse{} data
var tokenResponse = &AuthTokenResponse{
	Token: "auth-header.payload.signature",
}

func TestGetToken(t *testing.T) {
	resp, err := rc.Auth.GetToken(context.Background())
	if err != nil {
		t.Fatal("Error fetching Auth token\n", err)
	}

	if err := validateSchema(http.MethodPost, "/auth", http.StatusCreated, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
