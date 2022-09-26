package platform

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	// Auth token
	Token string
	// All available Rize API services
	Auth               *AuthService
	ComplianceWorkflow *ComplianceWorkflowService
}

// NewRizeClient initializes the RizeClient and all services
func NewRizeClient(cfg *RizeConfig) (*RizeClient, error) {
	// Enable debug logging
	if cfg.Debug {
		os.Setenv("debug", fmt.Sprintf("%t", cfg.Debug))
	}

	internal.Logger("client.NewRizeClient::Creating client...")

	// Validate client config
	if err := cfg.validateConfig(); err != nil {
		return nil, err
	}

	r := &RizeClient{}
	r.cfg = cfg
	r.svc.rizeClient = r // Store a reference to the RizeClient rather than creating one for each service

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
	req.Header.Add("Authorization", r.Token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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

	if ok := slices.Contains(internal.ENVIRONMENTS, strings.ToLower(cfg.Environment)); !ok {
		internal.Logger(fmt.Sprintf("Environment %s not recognized. Defaulting to sandbox...", cfg.Environment))
		cfg.Environment = "sandbox"
	}

	return nil
}

// Version outputs the current SDK version
func Version() string {
	return internal.SDK_VERSION
}
