package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// Handles all Auth related functionality
type authService service

// AuthTokenResponse is the response format received when fetching an Auth token
type AuthTokenResponse struct {
	Token string `json:"token"`
}

// GetToken generates an authorization token if the existing token is expired or not found.
// Otherwise, it will return the existing active token,
func (a *authService) GetToken(ctx context.Context) (*AuthTokenResponse, error) {
	// Check for missing or expired token
	if a.client.TokenCache.Token == "" || isExpired(a.client.TokenCache) {
		log.Println("Token is expired or does not exist. Fetching new token...")

		refreshToken, err := a.buildRefreshToken()
		if err != nil {
			return nil, err
		}
		// Store the refresh token (valid for 30 seconds)
		a.client.TokenCache.Token = refreshToken

		res, err := a.client.doRequest(ctx, http.MethodPost, "auth", nil, nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		response := &AuthTokenResponse{}
		if err = json.Unmarshal(body, response); err != nil {
			return nil, err
		}

		log.Println(fmt.Sprintf("Token response %+v", response))

		// Validate token exists
		if response.Token == "" {
			return nil, fmt.Errorf("Error fetching auth token")
		}

		//  Save token to client. Auth token is valid for 24hrs
		a.client.TokenCache.Timestamp = time.Now().Unix()
		a.client.TokenCache.Token = response.Token
		return response, nil
	}

	log.Println("Token is valid. Using existing auth token...")

	return &AuthTokenResponse{Token: a.client.TokenCache.Token}, nil
}

// Generates a JWT refresh token
func (a *authService) buildRefreshToken() (string, error) {
	// Encode JWT token with current time and programUID
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"sub": a.client.cfg.ProgramUID,
	})

	// Sign JWT token with the HMAC key
	signedToken, err := token.SignedString([]byte(a.client.cfg.HMACKey))
	if err != nil {
		return "", err
	}

	log.Println(fmt.Sprintf("Signed token %s", signedToken))

	return signedToken, nil
}

// Checks to see if the current Auth token should be refreshed
func isExpired(tc *TokenCache) bool {
	currentTime := time.Now().Unix()
	if tc.Timestamp == 0 || ((currentTime - tc.Timestamp) > internal.TokenMaxAge) {
		return true
	}

	return false
}
