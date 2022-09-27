package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rizefinance/rize-go-sdk/internal"
	"golang.org/x/exp/slices"
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
	// Cached auth token data
	*TokenCache
	// All available Rize API services
	Auth               *AuthService
	ComplianceWorkflow *ComplianceWorkflowService
}

// TokenCache stores Auth token data
type TokenCache struct {
	Token     string
	Timestamp int64
}

// Default API error format
type apiError struct {
	Errors []struct {
		Code       int       `json:"code"`
		Title      string    `json:"title"`
		Detail     string    `json:"detail"`
		OccurredAt time.Time `json:"occurred_at"`
	} `json:"errors"`
	Status int `json:"status"`
}

func (e *apiError) Error() string {
	return fmt.Sprintf("Error status %d and output:\n%+v\n", e.Status, e.Errors)
}

// NewRizeClient initializes the RizeClient and all services
func NewRizeClient(cfg *RizeConfig) (*RizeClient, error) {
	// Enable debug logging
	internal.EnableLogging(cfg.Debug)

	log.Println("Creating client...")

	// Validate client config
	if err := cfg.validateConfig(); err != nil {
		return nil, err
	}

	r := &RizeClient{}
	r.cfg = cfg
	r.svc.rizeClient = r // Store a reference to the RizeClient rather than creating one for each service
	r.TokenCache = &TokenCache{}

	// Initialize API Services
	r.Auth = (*AuthService)(&r.svc)
	r.ComplianceWorkflow = (*ComplianceWorkflowService)(&r.svc)

	// Generate Auth Token
	_, err := r.Auth.getToken()
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Make the API request and return the response body
func (r *RizeClient) doRequest(path string, method string, data io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("https://%s.rizefs.com/%s/%s", r.cfg.Environment, internal.APIBasePath, path)

	log.Println(fmt.Sprintf("Sending %s request to %s", method, url))

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", r.TokenCache.Token)

	client := &http.Client{
		Timeout: internal.APITimeoutSeconds,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for error response
	if res.StatusCode >= 400 {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var errorOut = &apiError{}
		if err = json.Unmarshal(body, &errorOut); err != nil {
			return nil, err
		}
		// Use apiError type to handle specific error codes from the API server
		return nil, errorOut
	}

	return res, nil
}

// Make sure that we have the proper configuration variables
func (cfg *RizeConfig) validateConfig() error {
	if cfg.ProgramUID == "" {
		return fmt.Errorf("RizeConfig error: ProgramUID is required")
	}

	if cfg.HMACKey == "" {
		return fmt.Errorf("RizeConfig error: HMACKey is required")
	}

	if ok := slices.Contains(internal.Environments, strings.ToLower(cfg.Environment)); !ok {
		log.Println(fmt.Sprintf("Environment %s not recognized. Defaulting to sandbox...", cfg.Environment))
		cfg.Environment = "sandbox"
	}

	return nil
}

// Version outputs the current SDK version
func Version() string {
	return internal.SDKVersion
}
