package platform

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Customer related functionality
type customerService service

// Customer data type
type Customer struct {
	UID                   string                 `json:"uid"`
	ExternalUID           string                 `json:"external_uid"`
	ActivatedAt           time.Time              `json:"activated_at"`
	CreatedAt             time.Time              `json:"created_at"`
	CustomerType          string                 `json:"customer_type"`
	Email                 string                 `json:"email"`
	Details               map[string]interface{} `json:"details"`
	KYCStatus             string                 `json:"kyc_status"`
	KYCStatusReasons      []string               `json:"kyc_status_reasons"`
	LockReason            string                 `json:"lock_reason"`
	LockedAt              time.Time              `json:"locked_at"`
	PoolUIDs              []string               `json:"pool_uids"`
	PrimaryCustomerUID    string                 `json:"primary_customer_uid"`
	ProfileResponses      []interface{}          `json:"profile_responses"`
	ProgramUID            string                 `json:"program_uid"`
	SecondaryCustomerUIDs []string               `json:"secondary_customer_uids"`
	Status                string                 `json:"status"`
	TotalBalance          string                 `json:"total_balance"`
}

// CustomerResponse is an API response containing a list of customers
type CustomerResponse struct {
	BaseResponse
	Data []Customer `json:"data"`
}

// CustomerListParams builds the query parameters used in querying customers
type CustomerListParams struct {
	Status           string `url:"status,omitempty"`
	IncludeInitiated bool   `url:"include_initiated,omitempty"`
	KYCStatus        string `url:"kyc_status,omitempty"`
	CustomerType     string `url:"customer_type,omitempty"`
	FirstName        string `url:"first_name,omitempty"`
	LastName         string `url:"last_name,omitempty"`
	Email            string `url:"email,omitempty"`
	Locked           bool   `url:"locked,omitempty"`
	ProgramUID       string `url:"program_uid,omitempty"`
	BusinessName     string `url:"business_name,omitempty"`
	ExternalUID      string `url:"external_uid,omitempty"`
	PoolUID          string `url:"pool_uid,omitempty"`
	Limit            int    `url:"limit,omitempty"`
	Offset           int    `url:"offset,omitempty"`
	Sort             string `url:"sort,omitempty"`
}

// List retrieves a list of Customers filtered by the given parameters
func (c *customerService) List(clp *CustomerListParams) (*CustomerResponse, error) {
	// Build CustomerListParams into query string params
	v, err := query.Values(clp)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodGet, "customers", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustomerResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
