package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Adjustment operations
type adjustmentService service

// Adjustment data type
type Adjustment struct {
	UID                 string          `json:"uid,omitempty"`
	CustomerUID         string          `json:"customer_uid,omitempty"`
	ExternalUID         string          `json:"external_uid,omitempty"`
	USDAdjustmentAmount float64         `json:"usd_adjustment_amount,omitempty"`
	AdjustmentType      *AdjustmentType `json:"adjustment_type,omitempty"`
	CreatedAt           time.Time       `json:"created_at,omitempty"`
	Status              string          `json:"status,omitempty"`
}

// AdjustmentType data type
type AdjustmentType struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Fee         bool   `json:"fee,omitempty"`
}

// AdjustmentListParams builds the query parameters used in querying Adjustments
type AdjustmentListParams struct {
	CustomerUID            string `url:"customer_uid,omitempty"`
	AdjustmentTypeUID      string `url:"adjustment_type_uid,omitempty"`
	ExternalUID            string `url:"external_uid,omitempty"`
	USDAdjustmentAmountMax int    `url:"usd_adjustment_amount_max,omitempty"`
	USDAdjustmentAmountMin int    `url:"usd_adjustment_amount_min,omitempty"`
	Sort                   string `url:"sort,omitempty"`
}

// AdjustmentCreateParams are the body params used when creating a new Adjustment
type AdjustmentCreateParams struct {
	ExternalUID         string  `json:"external_uid,omitempty"`
	CustomerUID         string  `json:"customer_uid"`
	USDAdjustmentAmount float64 `json:"usd_adjustment_amount"`
	AdjustmentTypeUID   string  `json:"adjustment_type_uid"`
}

// AdjustmentTypeListParams builds the query parameters used in querying Adjustment Types
type AdjustmentTypeListParams struct {
	CustomerUID string `url:"customer_uid,omitempty"`
	ProgramUID  string `url:"program_uid,omitempty"`
}

// AdjustmentResponse is an API response containing a list of Adjustments
type AdjustmentResponse struct {
	BaseResponse
	Data []*Adjustment `json:"data"`
}

// AdjustmentTypeResponse is an API response containing a list of Adjustments Types
type AdjustmentTypeResponse struct {
	BaseResponse
	Data []*AdjustmentType `json:"data"`
}

// List retrieves a list of Adjustments filtered by the given parameters
func (a *adjustmentService) List(alp *AdjustmentListParams) (*AdjustmentResponse, error) {
	// Build AdjustmentListParams into query string params
	v, err := query.Values(alp)
	if err != nil {
		return nil, err
	}

	res, err := a.client.doRequest(http.MethodGet, "adjustments", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &AdjustmentResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create a new Adjustment with the provided specification
func (a *adjustmentService) Create(acp *AdjustmentCreateParams) (*Adjustment, error) {
	if acp.CustomerUID == "" ||
		acp.USDAdjustmentAmount == 0 ||
		acp.AdjustmentTypeUID == "" {
		return nil, fmt.Errorf("CustomerUID, USDAdjustmentAmount and AdjustmentTypeUID are required")
	}

	bytesMessage, err := json.Marshal(acp)
	if err != nil {
		return nil, err
	}

	res, err := a.client.doRequest(http.MethodPost, "adjustments", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Adjustment{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Adjustment
func (a *adjustmentService) Get(uid string) (*Adjustment, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := a.client.doRequest(http.MethodGet, fmt.Sprintf("adjustments/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Adjustment{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListAdjustmentTypes retrieves a list of Adjustment Types filtered by the given parameters
func (a *adjustmentService) ListAdjustmentTypes(alp *AdjustmentTypeListParams) (*AdjustmentTypeResponse, error) {
	v, err := query.Values(alp)
	if err != nil {
		return nil, err
	}

	res, err := a.client.doRequest(http.MethodGet, "adjustment_types", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &AdjustmentTypeResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAdjustmentType returns a single Adjustment Type
func (a *adjustmentService) GetAdjustmentType(uid string) (*AdjustmentType, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := a.client.doRequest(http.MethodGet, fmt.Sprintf("adjustment_types/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &AdjustmentType{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
