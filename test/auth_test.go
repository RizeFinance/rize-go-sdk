package rize_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete AuthTokenResponse{} data
var tokenResponse = &rize.AuthTokenResponse{
	Token: "auth-header.payload.signature",
}

func TestAuthService_GetToken(t *testing.T) {
	resp, err := rc.Auth.GetToken(context.Background())
	if err != nil {
		t.Fatal("Error fetching Auth token\n", err)
	}

	if err := validateSchema(http.MethodPost, "/auth", http.StatusCreated, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
