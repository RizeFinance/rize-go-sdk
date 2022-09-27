package platform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// AuthService handles all Auth related functionality
type AuthService service

// Response format when fetching an Auth token
type authTokenResponse struct {
	Token string `json:"token"`
}

// Generate an authorization token
func (a *AuthService) getToken() (*authTokenResponse, error) {
	// Check for missing or expired token
	if a.rizeClient.TokenCache.Token == "" || isExpired(a.rizeClient.TokenCache) {
		log.Println("Token is expired or does not exist. Fetching new token...")

		refreshToken, err := a.buildRefreshToken()
		if err != nil {
			return nil, err
		}
		// Store the refresh token (valid for 30 seconds)
		a.rizeClient.TokenCache.Token = refreshToken

		res, err := a.rizeClient.doRequest(http.MethodPost, "auth", nil, nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		response := &authTokenResponse{}
		if err = json.Unmarshal(body, response); err != nil {
			return nil, err
		}

		log.Println(fmt.Sprintf("Token response %+v", response))

		// Validate token exists
		if response.Token == "" {
			return nil, fmt.Errorf("Error fetching auth token")
		}

		//  Save token to client. Auth token is valid for 24hrs
		a.rizeClient.TokenCache.Timestamp = time.Now().Unix()
		a.rizeClient.TokenCache.Token = response.Token
		return response, nil
	}

	log.Println("Token is valid. Using existing auth token...")

	return &authTokenResponse{Token: a.rizeClient.TokenCache.Token}, nil
}

// Generates a JWT refresh token
func (a *AuthService) buildRefreshToken() (string, error) {
	// Encode JWT token with current time and programUID
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"sub": a.rizeClient.cfg.ProgramUID,
	})

	// Sign JWT token with the HMAC key
	signedToken, err := token.SignedString([]byte(a.rizeClient.cfg.HMACKey))
	if err != nil {
		return "", err
	}

	log.Println(fmt.Sprintf("Signed token %s", signedToken))

	return signedToken, nil
}

// Checks to see if the current Auth token should be refreshed
func isExpired(tc *TokenCache) bool {
	currentTime := time.Now().Unix()
	if tc.Timestamp == 0 || ((currentTime - tc.Timestamp) > internal.APITokenMaxAge) {
		return true
	}

	return false
}
