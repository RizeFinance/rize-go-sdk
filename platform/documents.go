package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Document operations
type documentService service

// Document data type
type Document struct {
	UID                  string    `json:"uid,omitempty"`
	DocumentType         string    `json:"document_type,omitempty"`
	ScopeType            string    `json:"scope_type,omitempty"`
	Name                 string    `json:"name,omitempty"`
	PeriodStartedAt      time.Time `json:"period_started_at,omitempty"`
	PeriodEndedAt        time.Time `json:"period_ended_at,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	CustomerUids         []string  `json:"customer_uids,omitempty"`
	CustodialAccountUids []string  `json:"custodial_account_uids,omitempty"`
	SyntheticAccountUids []string  `json:"synthetic_account_uids,omitempty"`
}

// DocumentListParams builds the query parameters used in querying Documents
type DocumentListParams struct {
	DocumentType        string `json:"document_type,omitempty"`
	Month               int    `json:"month,omitempty"`
	Year                int    `json:"year,omitempty"`
	CustodialAccountUID string `json:"custodial_account_uid,omitempty"`
	CustomerUID         string `json:"customer_uid,omitempty"`
	SyntheticAccountUID string `json:"synthetic_account_uid,omitempty"`
	Limit               int    `json:"limit,omitempty"`
	Offset              int    `json:"offset,omitempty"`
}

// DocumentResponse is an API response containing a list of Documents
type DocumentResponse struct {
	BaseResponse
	Data []*Document `json:"data"`
}

// List retrieves a list of Documents filtered by the given parameters
func (d *documentService) List(dlp *DocumentListParams) (*DocumentResponse, error) {
	// Build DocumentListParams into query string params
	v, err := query.Values(dlp)
	if err != nil {
		return nil, err
	}

	res, err := d.rizeClient.doRequest(http.MethodGet, "documents", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &DocumentResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Document
func (d *documentService) Get(uid string) (*Document, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := d.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("documents/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Document{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// View is used to retrieve a Document and return it in either PDF or HTML format
func (d *documentService) View(uid string) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	// TODO: Does this require a different Accept header type (application/pdf)?
	res, err := d.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("documents/%s/view", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}
