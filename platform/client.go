package platform

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rizefinance/rize-go-sdk/internal"
)

// Service type to store the client reference
type service struct {
	rizeClient *RizeClient
}

// RizeConfig stores Rize configuration values
type RizeConfig struct {
	ProgramUID  string
	HMACKey     string
	Environment string
	Debug       bool
}

// RizeClient is the top-level client containing all APIs
type RizeClient struct {
	// All configuration values
	cfg *RizeConfig
	// Stores a reference to the RizeClient for child services
	svc service
	// Auth token
	token string
	// All available Rize API services
	Auth *AuthService
}

// NewRizeClient initializes the RizeClient and all services
func NewRizeClient(cfg *RizeConfig) *RizeClient {
	internal.Logger("client.NewRizeClient::Creating client...")

	r := &RizeClient{}
	r.cfg = cfg
	// Store a reference to the RizeClient rather than creating one for each service
	r.svc.rizeClient = r

	// Initialize API Services
	r.Auth = (*AuthService)(&r.svc)

	// Generate Auth Token
	r.Auth.getToken()

	return r
}

// Make the API request and return the response body
func (r *RizeClient) doRequest(path string, method string, data io.Reader) ([]byte, error) {
	url := fmt.Sprintf("https://%s.rizefs.com/api/v1/%s", r.cfg.Environment, path)

	internal.Logger(fmt.Sprintf("client.doRequest::Sending %s request to %s", method, url))

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", r.token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Version outputs the current SDK version
func Version() string {
	return internal.SDK_VERSION
}
