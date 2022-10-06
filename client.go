package rize

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/rizefinance/rize-go-sdk/internal"
	"golang.org/x/exp/slices"
)

// Service type to store the client reference
type service struct {
	client *Client
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

// Config stores Rize configuration values
type Config struct {
	// Program within the target environment
	ProgramUID string
	// HMAC key for the target environment
	HMACKey string
	// Rize infrastructure target environment. Defaults to `sandbox``
	Environment string
	// Provide your own HTTPClient configuration (optional)
	HTTPClient *http.Client
	// Change the API base URL for local/unit testing
	BaseURL string
	// Enable debug logging
	Debug bool
}

// Client is the top-level client containing all APIs
type Client struct {
	// All configuration values
	cfg *Config
	// Allows additional configuration options like proxy, timeouts, etc
	httpClient *http.Client
	// Set custom `user-agent` header string
	userAgent string
	// Cached Auth token data
	*TokenCache
	// All available Rize API services
	Adjustments         *adjustmentService
	Auth                *authService
	CardArtworks        *cardArtworkService
	ComplianceWorkflows *complianceWorkflowService
	CustodialAccounts   *custodialAccountService
	CustodialPartners   *custodialPartnerService
	CustomerProducts    *customerProductService
	Customers           *customerService
	DebitCards          *debitCardService
	Documents           *documentService
	Evaluations         *evaluationService
	PinwheelJobs        *pinwheelJobService
	KYCDocuments        *kycDocumentService
	Pools               *poolService
	Products            *productService
	Sandbox             *sandboxService
	SyntheticAccounts   *syntheticAccountService
	Transactions        *transactionService
	Transfers           *transferService
}

// TokenCache stores Auth token data
type TokenCache struct {
	Token     string
	Timestamp int64
}

// Error is the default API error format
type Error struct {
	Errors []struct {
		Code       int       `json:"code"`
		Title      string    `json:"title"`
		Detail     string    `json:"detail"`
		OccurredAt time.Time `json:"occurred_at"`
	} `json:"errors"`
	Status int `json:"status"`
}

// Format error output
func (e *Error) Error() string {
	return fmt.Sprintf("Error status %d and output:\n%+v\n", e.Status, e.Errors)
}

// NewClient initializes the Client and all services
func NewClient(cfg *Config) (*Client, error) {
	// Enable debug logging
	internal.EnableLogging(cfg.Debug)

	log.Println("Creating client...")

	// Validate client config
	if err := cfg.validateConfig(); err != nil {
		return nil, err
	}

	rc := &Client{}
	rc.cfg = cfg
	rc.httpClient = cfg.HTTPClient
	rc.userAgent = fmt.Sprintf("%s/%s (Go: %s)", "rize-go-sdk", internal.SDKVersion, runtime.Version())
	rc.TokenCache = &TokenCache{}

	// Initialize API Services
	rc.Adjustments = &adjustmentService{client: rc}
	rc.Auth = &authService{client: rc}
	rc.CardArtworks = &cardArtworkService{client: rc}
	rc.ComplianceWorkflows = &complianceWorkflowService{client: rc}
	rc.CustodialAccounts = &custodialAccountService{client: rc}
	rc.CustodialPartners = &custodialPartnerService{client: rc}
	rc.CustomerProducts = &customerProductService{client: rc}
	rc.Customers = &customerService{client: rc}
	rc.DebitCards = &debitCardService{client: rc}
	rc.Documents = &documentService{client: rc}
	rc.Evaluations = &evaluationService{client: rc}
	rc.KYCDocuments = &kycDocumentService{client: rc}
	rc.PinwheelJobs = &pinwheelJobService{client: rc}
	rc.Pools = &poolService{client: rc}
	rc.Products = &productService{client: rc}
	rc.Sandbox = &sandboxService{client: rc}
	rc.SyntheticAccounts = &syntheticAccountService{client: rc}
	rc.Transactions = &transactionService{client: rc}
	rc.Transfers = &transferService{client: rc}

	// Generate Auth Token
	_, err := rc.Auth.GetToken(context.Background())
	if err != nil {
		return nil, err
	}

	return rc, nil
}

// Make the API request and return an http.Response. Checks for valid auth token.
func (rc *Client) doRequest(ctx context.Context, method string, path string, query url.Values, data io.Reader) (*http.Response, error) {
	// Check for valid auth token and refresh if necessary
	if path != "auth" {
		if _, err := rc.Auth.GetToken(ctx); err != nil {
			return nil, err
		}
	}

	url := fmt.Sprintf("%s/%s/%s", rc.cfg.BaseURL, internal.BasePath, path)

	log.Println(fmt.Sprintf("Sending %s request to %s", method, url))

	req, err := http.NewRequestWithContext(ctx, method, url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", rc.userAgent)
	req.Header.Add("Authorization", rc.TokenCache.Token)
	req.URL.RawQuery = query.Encode()

	res, err := rc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for error response
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var errorOut = &Error{}
		if err = json.Unmarshal(body, &errorOut); err != nil {
			return nil, err
		}
		// Use RizeError type to handle specific error codes from the API server
		return nil, errorOut
	}

	return res, nil
}

// Make sure that we have the proper configuration variables
func (cfg *Config) validateConfig() error {
	if cfg.ProgramUID == "" {
		return fmt.Errorf("Config error: ProgramUID is required")
	}

	if cfg.HMACKey == "" {
		return fmt.Errorf("Config error: HMACKey is required")
	}

	if ok := slices.Contains(internal.Environments, strings.ToLower(cfg.Environment)); !ok {
		log.Println(fmt.Sprintf("Environment %s not recognized. Defaulting to sandbox...", cfg.Environment))
		cfg.Environment = "sandbox"
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: internal.TimeoutSeconds,
		}
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = fmt.Sprintf("https://%s.rizefs.com", cfg.Environment)
	}

	return nil
}

// Version outputs the current SDK version
func Version() string {
	return internal.SDKVersion
}
