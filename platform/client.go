package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rizefinance/rize-go-sdk/internal"
	"golang.org/x/exp/slices"
)

// Service type to store the client reference
type service struct {
	rizeClient *RizeClient
}

// BaseResponse is the default 'List' endpoint response.
// It is intended to be included in a response type specific to a service, which
// includes a Data array specific to that service type
type BaseResponse struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
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
	// Stores a reference to the RizeClient for child services to use internally
	svc service
	// Cached auth token data
	*TokenCache
	// All available Rize API services
	Auth               *authService
	ComplianceWorkflow *complianceWorkflowService
	Customer           *customerService
}

// TokenCache stores Auth token data
type TokenCache struct {
	Token     string
	Timestamp int64
}

// APIError is the default API error format
type APIError struct {
	Errors []struct {
		Code       int       `json:"code"`
		Title      string    `json:"title"`
		Detail     string    `json:"detail"`
		OccurredAt time.Time `json:"occurred_at"`
	} `json:"errors"`
	Status int `json:"status"`
}

func (e *APIError) Error() string {
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
	r.Auth = (*authService)(&r.svc)
	r.ComplianceWorkflow = (*complianceWorkflowService)(&r.svc)
	r.Customer = (*customerService)(&r.svc)

	// Generate Auth Token
	_, err := r.Auth.getToken()
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Make the API request and return an http.Response. Checks for valid auth token.
func (r *RizeClient) doRequest(method string, path string, query url.Values, data io.Reader) (*http.Response, error) {
	// Check for valid auth token and refresh if necessary
	if path != "auth" {
		if _, err := r.Auth.getToken(); err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("https://%s.rizefs.com/%s/%s", r.cfg.Environment, internal.APIBasePath, path)

	log.Println(fmt.Sprintf("Sending %s request to %s", method, url))

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
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
	if res.StatusCode >= http.StatusBadRequest {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var errorOut = &APIError{}
		if err = json.Unmarshal(body, &errorOut); err != nil {
			return nil, err
		}
		// Use APIError type to handle specific error codes from the API server
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
