package rize_test

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
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

var (
	rc  *rize.Client
	ts  *httptest.Server
	err error
	rl  []*http.Request

	errDetails = &rize.ErrorDetails{
		Code:       http.StatusNotFound,
		Title:      "Path/Method not found",
		OccurredAt: time.Now(),
	}
	errors = append([]*rize.ErrorDetails{}, errDetails)
)

// TestMain is the test runner init
func TestMain(m *testing.M) {
	// Create mock test server
	ts = httptest.NewServer(http.HandlerFunc(mockHandler))
	// Create a request list to log all mock requests fed into the handler
	rl = []*http.Request{}

	// Create new Rize client for tests
	config := rize.Config{
		ProgramUID:  "program_uid",
		HMACKey:     "hmac_key",
		Environment: "sandbox",
		BaseURL:     ts.URL,
		Debug:       false,
	}
	rc, err = rize.NewClient(&config)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Run all package tests
	code := m.Run()

	// Check for any schema paths not covered by the SDK
	if err := validatePaths(); err != nil {
		log.Fatal(err.Error())
	}

	os.Exit(code)

	defer ts.Close()
}

// Mock HTTP request handler for all test cases
func mockHandler(w http.ResponseWriter, r *http.Request) {
	// Consolidate requests by schema path
	path := strings.TrimPrefix(r.URL.Path+"/", "/"+internal.BasePath+"/")
	path = path[:strings.Index(path, "/")]

	// Log requests for OpenAPI schema comparison
	rl = append(rl, r)

	// Handle requests by path
	switch path {
	case "adjustments":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/adjustments" {
				adj := append([]*rize.Adjustment{}, adjustment)
				resp, _ := json.Marshal(&rize.AdjustmentListResponse{Data: adj})
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
			adjTypes := append([]*rize.AdjustmentType{}, adjustmentType)
			resp, _ := json.Marshal(&rize.AdjustmentTypeListResponse{Data: adjTypes})
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
			art := append([]*rize.CardArtwork{}, artwork)
			resp, _ := json.Marshal(&rize.CardArtworkListResponse{Data: art})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(artwork)
		w.Write(resp)
	case "compliance_workflows":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/compliance_workflows" {
				work := append([]*rize.Workflow{}, workflow)
				resp, _ := json.Marshal(&rize.WorkflowListResponse{Data: work})
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
				customers := append([]*rize.Customer{}, customer)
				resp, _ := json.Marshal(&rize.CustomerListResponse{Data: customers})
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
			acct := append([]*rize.CustodialAccount{}, custodialAccount)
			resp, _ := json.Marshal(&rize.CustodialAccountListResponse{Data: acct})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(custodialAccount)
		w.Write(resp)
	case "custodial_line_items":
		if r.URL.Path == "/"+internal.BasePath+"/custodial_line_items" {
			item := append([]*rize.CustodialLineItem{}, custodialLineItem)
			resp, _ := json.Marshal(&rize.CustodialLineItemListResponse{Data: item})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(custodialLineItem)
		w.Write(resp)
	case "custodial_partners":
		if r.URL.Path == "/"+internal.BasePath+"/custodial_partners" {
			acct := append([]*rize.CustodialPartner{}, custodialPartner)
			resp, _ := json.Marshal(&rize.CustodialPartnerListResponse{Data: acct})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(custodialPartner)
		w.Write(resp)
	case "customer_products":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/customer_products" {
				prod := append([]*rize.CustomerProduct{}, customerProduct)
				resp, _ := json.Marshal(&rize.CustomerProductListResponse{Data: prod})
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
				cards := append([]*rize.DebitCard{}, debitCard)
				resp, _ := json.Marshal(&rize.DebitCardListResponse{Data: cards})
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
				doc := append([]*rize.Document{}, document)
				resp, _ := json.Marshal(&rize.DocumentListResponse{Data: doc})
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
	case "evaluations":
		if r.URL.Path == "/"+internal.BasePath+"/evaluations" {
			eval := append([]*rize.Evaluation{}, evaluation)
			resp, _ := json.Marshal(&rize.EvaluationListResponse{Data: eval})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(evaluation)
		w.Write(resp)
	case "kyc_documents":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/kyc_documents" {
				doc := append([]*rize.KYCDocument{}, kycDocument)
				resp, _ := json.Marshal(&rize.KYCDocumentListResponse{Data: doc})
				w.Write(resp)
				return
			} else if r.URL.Path == "/"+internal.BasePath+"/kyc_documents/{uid/view}" {
				// Document file
				resp, _ := json.Marshal([]byte{})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(kycDocument)
			w.Write(resp)
		}
	case "pinwheel_jobs":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/pinwheel_jobs" {
				job := append([]*rize.PinwheelJob{}, pinwheelJob)
				resp, _ := json.Marshal(&rize.PinwheelJobListResponse{Data: job})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(pinwheelJob)
			w.Write(resp)
		}
	case "pools":
		if r.URL.Path == "/"+internal.BasePath+"/pools" {
			p := append([]*rize.Pool{}, pool)
			resp, _ := json.Marshal(&rize.PoolListResponse{Data: p})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(pool)
		w.Write(resp)
	case "products":
		if r.URL.Path == "/"+internal.BasePath+"/products" {
			p := append([]*rize.Product{}, product)
			resp, _ := json.Marshal(&rize.ProductListResponse{Data: p})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(product)
		w.Write(resp)
	case "sandbox":
		resp, _ := json.Marshal(&rize.SandboxResponse{Success: "true"})
		w.Write(resp)
	case "synthetic_accounts":
		switch r.Method {
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/synthetic_accounts" {
				synth := append([]*rize.SyntheticAccount{}, syntheticAccount)
				resp, _ := json.Marshal(&rize.SyntheticAccountListResponse{Data: synth})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(syntheticAccount)
			w.Write(resp)
		}
	case "synthetic_account_types":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/synthetic_account_types" {
				synth := append([]*rize.SyntheticAccountType{}, syntheticAccountType)
				resp, _ := json.Marshal(&rize.SyntheticAccountTypeListResponse{Data: synth})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(syntheticAccountType)
			w.Write(resp)
		}
	case "synthetic_line_items":
		if r.URL.Path == "/"+internal.BasePath+"/synthetic_line_items" {
			synth := append([]*rize.SyntheticLineItem{}, syntheticLineItem)
			resp, _ := json.Marshal(&rize.SyntheticLineItemListResponse{Data: synth})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(syntheticLineItem)
		w.Write(resp)
	case "transactions":
		if r.URL.Path == "/"+internal.BasePath+"/transactions" {
			t := append([]*rize.Transaction{}, transaction)
			resp, _ := json.Marshal(&rize.TransactionListResponse{Data: t})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(transaction)
		w.Write(resp)
	case "transaction_events":
		if r.URL.Path == "/"+internal.BasePath+"/transaction_events" {
			t := append([]*rize.TransactionEvent{}, transactionEvent)
			resp, _ := json.Marshal(&rize.TransactionEventListResponse{Data: t})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(transactionEvent)
		w.Write(resp)
	case "transfers":
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/"+internal.BasePath+"/transfers" {
				t := append([]*rize.Transfer{}, transfer)
				resp, _ := json.Marshal(&rize.TransferListResponse{Data: t})
				w.Write(resp)
				return
			}
			fallthrough
		default:
			resp, _ := json.Marshal(transfer)
			w.Write(resp)
		}
	default:
		errDetails.Detail = fmt.Sprintf("Error in path %s, method %s", path, r.Method)
		resp, _ := json.Marshal(&rize.Error{Errors: errors, Status: http.StatusNotFound})
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
			if err != nil {
				return err
			}
		} else {
			bytesParams, err = json.Marshal(&bodyParams)
			if err != nil {
				return err
			}
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

// Check SDK path/method combinations against the OpenAPI schema to find anything that's missing
func validatePaths() error {
	// Convert test requests to OpenAPI path format
	paths := []string{}
	for _, r := range rl {
		r.URL.Scheme = "https"
		r.URL.Host = "sandbox.newline53.com"
		route, err := internal.FindRoute(r)
		if err != nil {
			return err
		}
		paths = append(paths, route)
	}

	schemaPaths := internal.BuildSchemaPathsList()
	diff := internal.Difference(schemaPaths, paths)
	if len(diff) > 0 {
		return fmt.Errorf("SDK is missing the following path + method combinations that are present in the OpenAPI schema:\n%s", diff)
	}

	return nil
}
