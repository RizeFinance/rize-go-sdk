package rize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rizefinance/rize-go-sdk/internal"
)

var (
	rc  *Client
	ts  *httptest.Server
	err error
)

// Complete Customer{} response data
var customer = &Customer{
	UID:                "y9reyPMNEWuuYSC1",
	ExternalUID:        "partner-generated-id",
	ActivatedAt:        time.Now(),
	CreatedAt:          time.Now(),
	CustomerType:       "primary",
	Email:              "olive.oyl@rizemoney.com",
	KYCStatus:          "manual_review",
	KYCStatusReasons:   []string{"Approved"},
	LockReason:         "other",
	LockedAt:           time.Now(),
	PIIConfirmedAt:     time.Now(),
	PoolUIDs:           []string{"HiuQZJNjCd79LLYq", "NoPJB9g9ZQTh5qMv"},
	PrimaryCustomerUID: "EhrQZJNjCd79LLYq",
	ProfileResponses: []*CustomerProfileResponse{{
		ProfileRequirement: "Please provide your approximate annual income in USD.",
		ProfileResponse: &CustomerProfileResponseItem{
			Num0: "string",
			Num1: "string",
			Num2: "string",
		},
		ProfileRequirementUID: "ptRLF7nQvy8VoqM1",
	}},
	ProgramUID:            "kaxHFJnWvJxRJZxr",
	SecondaryCustomerUIDs: []string{"464QyebpxbBNrGkX"},
	Status:                "initiated",
	TotalBalance:          "12345.67",
	Details: &CustomerDetails{
		FirstName:    "Olive",
		MiddleName:   "Olivia",
		LastName:     "Oyl",
		Suffix:       "Jr.",
		Phone:        "5555551212",
		BusinessName: "Oliver's Olive Emporium",
		DOB:          internal.DOB(time.Now()),
		SSN:          "111-11-1111",
		Address: &CustomerAddress{
			Street1:    "123 Abc St.",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
		SSNLastFour: "3333",
	},
}

func init() {
	// Create mock test server
	ts = httptest.NewServer(http.HandlerFunc(mockHandler))

	// Create new Rize client for tests
	config := Config{
		ProgramUID:  "program_uid",
		HMACKey:     "hmac_key",
		Environment: "sandbox",
		BaseURL:     ts.URL,
		Debug:       true,
	}
	rc, err = NewClient(&config)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Mock HTTP request handler
func mockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v1/auth":
		bytesMessage, err := json.Marshal(&AuthTokenResponse{Token: "auth-header.payload.signature"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Write(bytesMessage)
	case "/api/v1/customers":
		switch r.Method {
		case http.MethodGet:
			customers := append([]*Customer{}, customer)
			bytesMessage, err := json.Marshal(&CustomerResponse{Data: customers})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			w.Write(bytesMessage)
		case http.MethodPost:
			bytesMessage, err := json.Marshal(customer)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			w.Write(bytesMessage)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func TestList(t *testing.T) {
	params := CustomerListParams{
		UID:              "uKxmLxUEiSj5h4M3",
		Status:           "identity_verified",
		IncludeInitiated: true,
		KYCStatus:        "denied",
		CustomerType:     "primary",
		FirstName:        "Olive",
		LastName:         "Oyl",
		Email:            "olive.oyl@popeyes.com",
		Locked:           false,
		ProgramUID:       "pQtTCSXz57fuefzp",
		BusinessName:     "Business inc",
		ExternalUID:      "client-generated-id",
		PoolUID:          "wTSMX1GubP21ev2h",
		Limit:            100,
		Offset:           0,
		Sort:             "first_name_asc",
	}

	resp, err := rc.Customers.List(context.Background(), &params)
	if err != nil {
		t.Fatal("Error fetching customers\n", err)
	}

	if err := validateSchema(http.MethodGet, "/customers", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreate(t *testing.T) {
	params := CustomerCreateParams{
		ExternalUID:  "client-generated-id",
		CustomerType: "primary",
		Email:        "olive.oyl@popeyes.com",
	}

	resp, err := rc.Customers.Create(context.Background(), &params)
	if err != nil {
		t.Fatal("Error creating customer\n", err)
	}

	if err := validateSchema(http.MethodPost, "/customers", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func validateSchema(method string, path string, status int, queryParams interface{}, bodyParams interface{}, resp interface{}) error {
	var (
		v   url.Values
		b   io.Reader
		err error
	)

	// Handle query params
	if queryParams != nil {
		v, err = query.Values(queryParams)
		if err != nil {
			return err
		}
	}

	// Handle body params
	if bodyParams != nil {
		bytesMessage, err := json.Marshal(bodyParams)
		if err != nil {
			return err
		}
		b = bytes.NewBuffer(bytesMessage)
	}

	// Validate request schema
	input, err := internal.ValidateRequest(method, path, v, b)
	if err != nil {
		return err
	}

	// Validate response schema
	bytesResp, err := json.Marshal(&resp)
	if err != nil {
		return err
	}
	if err := internal.ValidateResponse(status, bytesResp, input); err != nil {
		return err
	}

	// Generate list of request keys (query string or body param) from OpenAPI schema request
	schemaReq, _, err := internal.GetRequestKeys(method, path)
	if err != nil {
		return err
	}

	// Generate list of response keys from OpenAPI schema response
	schemaResp := internal.RecurseResponseKeys(method, path, status)

	// Generate list of keys from SDK request parameter json
	var sdkReq = []string{}
	if queryParams != nil || bodyParams != nil {
		var bytesParams []byte
		if queryParams != nil {
			bytesParams, err = json.Marshal(&queryParams)
		} else {
			bytesParams, err = json.Marshal(&bodyParams)
		}
		p := make(map[string]interface{})
		if err := json.Unmarshal(bytesParams, &p); err != nil {
			return err
		}
		sdkReq = internal.JSONKeys(p)
	}

	// Generate list of keys from SDK response json
	k := make(map[string]interface{})
	if err := json.Unmarshal(bytesResp, &k); err != nil {
		return err
	}
	sdkResp := internal.JSONKeys(k)

	// Compare request keys from OpenAPI spec with keys from SDK struct
	reqDiff := internal.Difference(schemaReq, sdkReq)
	if len(reqDiff) > 0 {
		return fmt.Errorf("Request is missing the following keys that are present in the OpenAPI schema:\n%s", reqDiff)
	}

	// Compare response keys from OpenAPI spec with keys from SDK struct
	respDiff := internal.Difference(schemaResp, sdkResp)
	if len(respDiff) > 0 {
		return fmt.Errorf("Response is missing the following keys that are present in the OpenAPI schema:\n%s", respDiff)
	}

	return nil
}
