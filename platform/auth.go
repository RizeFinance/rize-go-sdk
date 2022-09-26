package platform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	token, err := a.buildRefreshToken()
	if err != nil {
		return nil, err
	}
	a.rizeClient.Token = token

	res, err := a.rizeClient.doRequest("auth", "POST", nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &authTokenResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	internal.Logger(fmt.Sprintf("auth.getToken:::Token response %+v", response))

	return response, nil
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

	internal.Logger(fmt.Sprintf("auth.buildRefreshToken::Signed token %s", signedToken))

	return signedToken, nil
}
