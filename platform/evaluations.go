package platform

import (
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
	UID       string              `json:"uid"`
	Outcome   string              `json:"outcome"`
	CreatedAt time.Time           `json:"created_at"`
	Flags     *EvaluationFlag     `json:"flags"`
	PIIMatch  *EvaluationPIIMatch `json:"pii_match"`
}

// EvaluationFlag provides a mapping of categories to outcomes for those categories
type EvaluationFlag struct {
	DocumentQualityCheck bool `json:"Document Quality Check"`
	FraudCheck           bool `json:"Fraud Check"`
	FinancialCheck       bool `json:"Financial Check"`
	WatchListCheck       bool `json:"Watch List Check"`
}

// EvaluationPIIMatch provides a mapping of KYC categories to results returned from various services
type EvaluationPIIMatch struct {
	DOBMatch     bool `json:"DOB Match"`
	SSNMatch     bool `json:"SSN Match"`
	NameMatch    bool `json:"Name Match"`
	EmailMatch   bool `json:"Email Match"`
	PhoneMatch   bool `json:"Phone Match"`
	AddressMatch bool `json:"Address Match"`
}

// EvaluationListParams builds the query parameters used in querying Evaluations
type EvaluationListParams struct {
	CustomerUID string `url:"customer_uid,omitempty"`
	Latest      bool   `url:"latest,omitempty"`
}

// EvaluationResponse is an API response containing a list of Evaluations
type EvaluationResponse struct {
	BaseResponse
	Data []*Evaluation `json:"data"`
}

// List retrieves a list of Evaluations filtered by the given parameters
func (p *evaluationService) List(plp *EvaluationListParams) (*EvaluationResponse, error) {
	// Build EvaluationListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := p.rizeClient.doRequest(http.MethodGet, "evaluations", v, nil)
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
func (p *evaluationService) Get(uid string) (*Evaluation, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("evaluations/%s", uid), nil, nil)
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
