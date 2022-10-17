package rize

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/rizefinance/rize-go-sdk/internal"
)

var (
	rc  *Client
	ts  *httptest.Server
	err error

	errDetails = &ErrorDetails{
		Code:       http.StatusNotFound,
		Title:      "Path/Method not found",
		OccurredAt: time.Now(),
	}
	errors = append([]*ErrorDetails{}, errDetails)
)

// TestMain is the test runner init
func TestMain(m *testing.M) {
	// Create mock test server
	ts = httptest.NewServer(http.HandlerFunc(mockHandler))

	// Create new Rize client for tests
	config := Config{
		ProgramUID:  "program_uid",
		HMACKey:     "hmac_key",
		Environment: "sandbox",
		BaseURL:     ts.URL,
		Debug:       false,
	}
	rc, err = NewClient(&config)
	if err != nil {
		log.Fatal(err.Error())
	}

	os.Exit(m.Run())

	defer ts.Close()
}

// Mock HTTP request handler for all test cases
func mockHandler(w http.ResponseWriter, r *http.Request) {
	// Consolidate requests by schema path
	path := strings.TrimPrefix(r.URL.Path+"/", "/"+internal.BasePath+"/")
	path = path[:strings.Index(path, "/")]
	switch path {
	case "adjustments":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/adjustments" {
				adj := append([]*Adjustment{}, adjustment)
				resp, _ := json.Marshal(&AdjustmentResponse{Data: adj})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(adjustment)
			w.Write(resp)
		}
	case "adjustment_types":
		if r.URL.Path == "/"+internal.BasePath+"/adjustment_types" {
			adjTypes := append([]*AdjustmentType{}, adjustmentType)
			resp, _ := json.Marshal(&AdjustmentTypeResponse{Data: adjTypes})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(adjustmentType)
		w.Write(resp)
	case "assets":
		// virtual_card_image
		resp, _ := json.Marshal([]byte{})
		w.Write(resp)
	case "auth":
		resp, _ := json.Marshal(tokenResponse)
		w.Write(resp)
	case "card_artworks":
		if r.URL.Path == "/"+internal.BasePath+"/card_artworks" {
			art := append([]*CardArtwork{}, artwork)
			resp, _ := json.Marshal(&CardArtworkResponse{Data: art})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(artwork)
		w.Write(resp)
	case "compliance_workflows":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/compliance_workflows" {
				work := append([]*Workflow{}, workflow)
				resp, _ := json.Marshal(&WorkflowResponse{Data: work})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(workflow)
			w.Write(resp)
		}
	case "customers":
		switch r.Method {
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/customers" {
				customers := append([]*Customer{}, customer)
				resp, _ := json.Marshal(&CustomerResponse{Data: customers})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(customer)
			w.Write(resp)
		}
	case "custodial_accounts":
		if r.URL.Path == "/"+internal.BasePath+"/custodial_accounts" {
			acct := append([]*CustodialAccount{}, custodialAccount)
			resp, _ := json.Marshal(&CustodialAccountResponse{Data: acct})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(custodialAccount)
		w.Write(resp)
	case "custodial_partners":
		if r.URL.Path == "/"+internal.BasePath+"/custodial_partners" {
			acct := append([]*CustodialPartner{}, custodialPartner)
			resp, _ := json.Marshal(&CustodialPartnerResponse{Data: acct})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(custodialPartner)
		w.Write(resp)
	case "customer_products":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/customer_products" {
				prod := append([]*CustomerProduct{}, customerProduct)
				resp, _ := json.Marshal(&CustomerProductResponse{Data: prod})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(customerProduct)
			w.Write(resp)
		}
	case "debit_cards":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/debit_cards" {
				cards := append([]*DebitCard{}, debitCard)
				resp, _ := json.Marshal(&DebitCardResponse{Data: cards})
				w.Write(resp)
				return
			} else if r.URL.Path == "/"+internal.BasePath+"/debit_cards/{uid}/pin_change_token" {
				resp, _ := json.Marshal(PINToken)
				w.Write(resp)
				return
			} else if r.URL.Path == "/"+internal.BasePath+"/debit_cards/{uid}/access_token" {
				resp, _ := json.Marshal(accessToken)
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(debitCard)
			w.Write(resp)
		}
	case "documents":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/documents" {
				doc := append([]*Document{}, document)
				resp, _ := json.Marshal(&DocumentResponse{Data: doc})
				w.Write(resp)
				return
			} else if r.URL.Path == "/"+internal.BasePath+"/documents/{uid/view}" {
				// Document file
				resp, _ := json.Marshal([]byte{})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(document)
			w.Write(resp)
		}
	default:
		errDetails.Detail = fmt.Sprintf("Error in path %s, method %s", path, r.Method)
		resp, _ := json.Marshal(&Error{Errors: errors, Status: http.StatusNotFound})
		w.WriteHeader(http.StatusNotFound)
		w.Write(resp)
	}
}

// Validate SDK requests and responses against the latest OpenAPI spec file for the Rize Platform. Also searches
// for any missing req/resp fields between the SDK and OpenAPI spec.
func validateSchema(method string, path string, status int, queryParams interface{}, bodyParams interface{}, resp interface{}) error {
	var (
		v                   url.Values
		b                   io.Reader
		bytesResp           []byte
		schemaResp, sdkResp []string
		err                 error
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
	if resp != nil {
		bytesResp, err = json.Marshal(&resp)
		if err != nil {
			return err
		}
		if err := internal.ValidateResponse(status, bytesResp, input); err != nil {
			return err
		}
	}

	// Generate list of request keys (query string or body param) from OpenAPI schema request
	schemaReq, err := internal.GetRequestKeys(method, path, status)
	if err != nil {
		return err
	}

	// Skip response validation for requests that do not generate a response
	if resp != nil {
		// Generate list of response keys from OpenAPI schema response
		schemaResp, err = internal.RecurseResponseKeys(method, path, status)
		if err != nil {
			return err
		}
	}

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

	if resp != nil {
		// Generate list of keys from SDK response json
		k := make(map[string]interface{})
		if err := json.Unmarshal(bytesResp, &k); err != nil {
			return err
		}
		sdkResp = internal.JSONKeys(k)
	}

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
