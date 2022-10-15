package rize

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Evaluation operations
type evaluationService service

// Evaluation data type
type Evaluation struct {
	UID       string              `json:"uid,omitempty"`
	Outcome   string              `json:"outcome,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	Flags     *EvaluationFlag     `json:"flags,omitempty"`
	PIIMatch  *EvaluationPIIMatch `json:"pii_match,omitempty"`
}

// EvaluationFlag provides a mapping of categories to outcomes for those categories
type EvaluationFlag struct {
	DocumentQualityCheck bool `json:"Document Quality Check,omitempty"`
	FraudCheck           bool `json:"Fraud Check,omitempty"`
	FinancialCheck       bool `json:"Financial Check,omitempty"`
	WatchListCheck       bool `json:"Watch List Check,omitempty"`
}

// EvaluationPIIMatch provides a mapping of KYC categories to results returned from various services
type EvaluationPIIMatch struct {
	DOBMatch     bool `json:"DOB Match,omitempty"`
	SSNMatch     bool `json:"SSN Match,omitempty"`
	NameMatch    bool `json:"Name Match,omitempty"`
	EmailMatch   bool `json:"Email Match,omitempty"`
	PhoneMatch   bool `json:"Phone Match,omitempty"`
	AddressMatch bool `json:"Address Match,omitempty"`
}

// EvaluationListParams builds the query parameters used in querying Evaluations
type EvaluationListParams struct {
	CustomerUID string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	Latest      bool   `url:"latest,omitempty" json:"latest,omitempty"`
}

// EvaluationResponse is an API response containing a list of Evaluations
type EvaluationResponse struct {
	BaseResponse
	Data []*Evaluation `json:"data"`
}

// List retrieves a list of Evaluations filtered by the given parameters
func (p *evaluationService) List(ctx context.Context, plp *EvaluationListParams) (*EvaluationResponse, error) {
	// Build EvaluationListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, "evaluations", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &EvaluationResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Evaluation
func (p *evaluationService) Get(ctx context.Context, uid string) (*Evaluation, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("evaluations/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Evaluation{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
