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
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/rizefinance/rize-go-sdk/internal"
)

var (
	rc  *Client
	ts  *httptest.Server
	err error
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
		Debug:       true,
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

// Validate SDK requests and responses against the latest OpenAPI spec file for the Rize Platform. Also searches
// for any missing req/resp fields between the SDK and OpenAPI spec.
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
