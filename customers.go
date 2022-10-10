package rize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rizefinance/rize-go-sdk/internal"
)

// Handles all Customer related functionality
type customerService service

// Customer data type
type Customer struct {
	UID                   string                     `json:"uid,omitempty"`
	ExternalUID           string                     `json:"external_uid,omitempty"`
	ActivatedAt           time.Time                  `json:"activated_at,omitempty"`
	CreatedAt             time.Time                  `json:"created_at,omitempty"`
	CustomerType          string                     `json:"customer_type,omitempty"`
	Email                 string                     `json:"email,omitempty"`
	Details               *CustomerDetails           `json:"details,omitempty"`
	KYCStatus             string                     `json:"kyc_status,omitempty"`
	KYCStatusReasons      []string                   `json:"kyc_status_reasons,omitempty"`
	LockReason            string                     `json:"lock_reason,omitempty"`
	LockedAt              time.Time                  `json:"locked_at,omitempty"`
	PIIConfirmedAt        time.Time                  `json:"pii_confirmed_at,omitempty"`
	PoolUIDs              []string                   `json:"pool_uids,omitempty"`
	PrimaryCustomerUID    string                     `json:"primary_customer_uid,omitempty"`
	ProfileResponses      []*CustomerProfileResponse `json:"profile_responses,omitempty"`
	ProgramUID            string                     `json:"program_uid,omitempty"`
	SecondaryCustomerUIDs []string                   `json:"secondary_customer_uids,omitempty"`
	Status                string                     `json:"status,omitempty"`
	TotalBalance          string                     `json:"total_balance,omitempty"`
}

// CustomerDetails is an object containing the supplied identifying information for the Customer
type CustomerDetails struct {
	FirstName    string           `json:"first_name,omitempty"`
	MiddleName   string           `json:"middle_name,omitempty"`
	LastName     string           `json:"last_name,omitempty"`
	Suffix       string           `json:"suffix,omitempty"`
	Phone        string           `json:"phone,omitempty"`
	BusinessName string           `json:"business_name,omitempty"`
	DOB          internal.DOB     `json:"dob,omitempty"`
	SSN          string           `json:"ssn,omitempty"`
	SSNLastFour  string           `json:"ssn_last_four,omitempty"`
	Address      *CustomerAddress `json:"address,omitempty"`
}

// CustomerAddress information
type CustomerAddress struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// CustomerProfileResponse contains Profile Response info
type CustomerProfileResponse struct {
	ProfileRequirement    string                       `json:"profile_requirement,omitempty"`
	ProfileRequirementUID string                       `json:"profile_requirement_uid,omitempty"`
	ProfileResponse       *CustomerProfileResponseItem `json:"profile_response,omitempty"`
}

// CustomerProfileResponseItem contains the Customer's response to the Profile Requirement
type CustomerProfileResponseItem struct {
	Num0 string `json:"0,omitempty"`
	Num1 string `json:"1,omitempty"`
	Num2 string `json:"2,omitempty"`
}

// CustomerListParams builds the query parameters used in querying Customers
type CustomerListParams struct {
	UID              string `url:"uid,omitempty" json:"uid,omitempty"`
	Status           string `url:"status,omitempty" json:"status,omitempty"`
	IncludeInitiated bool   `url:"include_initiated,omitempty" json:"include_initiated"`
	KYCStatus        string `url:"kyc_status,omitempty" json:"kyc_status,omitempty"`
	CustomerType     string `url:"customer_type,omitempty" json:"customer_type,omitempty"`
	FirstName        string `url:"first_name,omitempty" json:"first_name,omitempty"`
	LastName         string `url:"last_name,omitempty" json:"last_name,omitempty"`
	Email            string `url:"email,omitempty" json:"email,omitempty"`
	Locked           bool   `url:"locked,omitempty" json:"locked"`
	ProgramUID       string `url:"program_uid,omitempty" json:"program_uid,omitempty"`
	BusinessName     string `url:"business_name,omitempty" json:"business_name,omitempty"`
	ExternalUID      string `url:"external_uid,omitempty" json:"external_uid,omitempty"`
	PoolUID          string `url:"pool_uid,omitempty" json:"pool_uid,omitempty"`
	Limit            int    `url:"limit,omitempty" json:"limit"`
	Offset           int    `url:"offset,omitempty" json:"offset"`
	Sort             string `url:"sort,omitempty" json:"sort,omitempty"`
}

// CustomerCreateParams are the body params used when creating a new Customer
type CustomerCreateParams struct {
	ExternalUID  string `json:"external_uid,omitempty"`
	CustomerType string `json:"customer_type,omitempty"`
	Email        string `json:"email"`
}

// CustomerUpdateParams are the body params used when updating a Customer
type CustomerUpdateParams struct {
	Email   string          `json:"email,omitempty"`
	Details CustomerDetails `json:"details,omitempty"`
}

// CustomerProfileResponseParams are the body params used when updating Customer Profile responses
type CustomerProfileResponseParams struct {
	ProfileRequirementUID string `json:"profile_requirement_uid"`
	ProfileResponse       string `json:"profile_response"`
}

// CustomerProfileResponseOrderedListParams are the body params used when updating Customer Profile responses with the `ordered_list` requirement type
type CustomerProfileResponseOrderedListParams struct {
	ProfileRequirementUID string                       `json:"profile_requirement_uid"`
	ProfileResponse       *CustomerProfileResponseItem `json:"profile_response"`
}

// SecondaryCustomerParams are the body params used when creating a new Secondary Customer
type SecondaryCustomerParams struct {
	ExternalUID        string           `json:"external_uid,omitempty"`
	PrimaryCustomerUID string           `json:"primary_customer_uid"`
	Email              string           `json:"email,omitempty"`
	Details            *CustomerDetails `json:"details"`
}

// CustomerResponse is an API response containing a list of Customers
type CustomerResponse struct {
	BaseResponse
	Data []*Customer `json:"data"`
}

// List retrieves a list of Customers filtered by the given parameters
func (c *customerService) List(ctx context.Context, clp *CustomerListParams) (*CustomerResponse, error) {
	// Build CustomerListParams into query string params
	v, err := query.Values(clp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, "customers", v, nil)
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

// Create is used to initialize a new Customer with an email and external_uid
func (c *customerService) Create(ctx context.Context, ccp *CustomerCreateParams) (*Customer, error) {
	if ccp.Email == "" {
		return nil, fmt.Errorf("Email is required")
	}

	bytesMessage, err := json.Marshal(ccp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPost, "customers", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get retrieves overall status about a Customer as well as their total Asset Balances across all accounts
func (c *customerService) Get(ctx context.Context, uid string) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("customers/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update will submit or update a Customer's personally identifiable information (PII) after they are created
func (c *customerService) Update(ctx context.Context, uid string, cu *CustomerUpdateParams) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	bytesMessage, err := json.Marshal(cu)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Delete will archive a Customer
func (c *customerService) Delete(ctx context.Context, uid string, archiveNote string) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	var params = struct {
		ArchiveNote string `json:"archive_note,omitempty"`
	}{}
	params.ArchiveNote = archiveNote

	bytesMessage, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodDelete, fmt.Sprintf("customers/%s", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

// ConfirmPIIData is used to explicitly confirm a Customer's PII data is up-to-date in order to add additional products
func (c *customerService) ConfirmPIIData(ctx context.Context, uid string) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s/identity_confirmation", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Lock will freeze all activities relating to the Customer
func (c *customerService) Lock(ctx context.Context, uid string, lockNote string, lockReason string) (*Customer, error) {
	if uid == "" || lockReason == "" {
		return nil, fmt.Errorf("UID and lockReason are required")
	}

	var params = struct {
		LockNote   string `json:"lock_note,omitempty"`
		LockReason string `json:"lock_reason"`
	}{}
	params.LockNote = lockNote
	params.LockReason = lockReason

	bytesMessage, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s/lock", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Unlock will remove the Customer lock, returning their state to normal
func (c *customerService) Unlock(ctx context.Context, uid string, lockNote string, unlockReason string) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	var params = struct {
		LockNote     string `json:"lock_note,omitempty"`
		UnlockReason string `json:"unlock_reason,omitempty"`
	}{}
	params.LockNote = lockNote
	params.UnlockReason = unlockReason

	bytesMessage, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s/unlock", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateProfileResponses is used to submit a Customer's Profile Responses to Profile Requirements
func (c *customerService) UpdateProfileResponses(ctx context.Context, uid string, cpp []*CustomerProfileResponseParams) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	for _, v := range cpp {
		if v.ProfileRequirementUID == "" || v.ProfileResponse == "" {
			return nil, fmt.Errorf("ProfileRequirementUID and ProfileResponse are required")
		}
	}

	// Wrap profile response params in a `details` json object
	var details = struct {
		Details []*CustomerProfileResponseParams `json:"details"`
	}{
		Details: cpp,
	}
	bytesMessage, err := json.Marshal(&details)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s/update_profile_responses", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateProfileResponsesOrderedList is used to update Customer's Profile Responses with the `ordered_list` requirement type
func (c *customerService) UpdateProfileResponsesOrderedList(ctx context.Context, uid string, cpp []*CustomerProfileResponseOrderedListParams) (*Customer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	for _, v := range cpp {
		if v.ProfileRequirementUID == "" || v.ProfileResponse.Num0 == "" {
			return nil, fmt.Errorf("ProfileRequirementUID and ProfileResponse are required")
		}
	}

	// Wrap profile response params in a `details` json object
	var details = struct {
		Details []*CustomerProfileResponseOrderedListParams `json:"details"`
	}{
		Details: cpp,
	}
	bytesMessage, err := json.Marshal(&details)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPut, fmt.Sprintf("customers/%s/update_profile_responses", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateSecondaryCustomer is used to create a new Secondary Customer
func (c *customerService) CreateSecondaryCustomer(ctx context.Context, scp *SecondaryCustomerParams) (*Customer, error) {
	if scp.PrimaryCustomerUID == "" {
		return nil, fmt.Errorf("PrimaryCustomerUID is required")
	}

	bytesMessage, err := json.Marshal(&scp)
	if err != nil {
		return nil, err
	}

	res, err := c.client.doRequest(ctx, http.MethodPost, "customers/create_secondary", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Customer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
