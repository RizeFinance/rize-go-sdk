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

// Handles all PinwheelJob operations
type pinwheelJobService service

// PinwheelJob data type
type PinwheelJob struct {
	UID                 string    `json:"uid,omitempty"`
	SyntheticAccountUID string    `json:"synthetic_account_uid,omitempty"`
	Status              string    `json:"status,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	StatusUpdatedAt     time.Time `json:"status_updated_at,omitempty"`
	CustomerUID         string    `json:"customer_uid,omitempty"`
	LinkToken           string    `json:"link_token,omitempty"`
	ExpiresAt           time.Time `json:"expires_at,omitempty"`
}

// PinwheelJobListParams builds the query parameters used in querying Pinwheel Jobs
type PinwheelJobListParams struct {
	CustomerUID         string `url:"customer_uid,omitempty"`
	SyntheticAccountUID string `url:"synthetic_account_uid,omitempty"`
	Limit               int    `url:"limit,omitempty"`
	Offset              int    `url:"offset,omitempty"`
}

// PinwheelJobCreateParams are the body params used when creating a new Pinwheel Job
type PinwheelJobCreateParams struct {
	JobNames             []string `json:"job_names"`
	SyntheticAccountUID  string   `json:"synthetic_account_uid"`
	Amount               int      `json:"amount,omitempty"`
	DisablePartialSwitch bool     `json:"disable_partial_switch,omitempty"`
	OrganizationName     string   `json:"organization_name,omitempty"`
	SkipWelcomeScreen    bool     `json:"skip_welcome_screen,omitempty"`
}

// PinwheelJobResponse is an API response containing a list of Pinwheel Jobs
type PinwheelJobResponse struct {
	BaseResponse
	Data []*PinwheelJob `json:"data"`
}

// List retrieves a list of Pinwheel Jobs filtered by the given parameters
func (p *pinwheelJobService) List(plp *PinwheelJobListParams) (*PinwheelJobResponse, error) {
	// Build PinwheelJobListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := p.client.doRequest(http.MethodGet, "pinwheel_jobs", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &PinwheelJobResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create is used to initialize a new Pinwheel Job and return a pinwheel_link_token to be used with the Pinwheel Link SDK
func (p *pinwheelJobService) Create(pcp *PinwheelJobCreateParams) (*PinwheelJob, error) {
	if len(pcp.JobNames) == 0 || pcp.SyntheticAccountUID == "" {
		return nil, fmt.Errorf("JobNames and SyntheticAccountUID are required")
	}

	bytesMessage, err := json.Marshal(pcp)
	if err != nil {
		return nil, err
	}

	res, err := p.client.doRequest(http.MethodPost, "pinwheel_jobs", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &PinwheelJob{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single PinwheelJob
func (p *pinwheelJobService) Get(uid string) (*PinwheelJob, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.client.doRequest(http.MethodGet, fmt.Sprintf("pinwheel_jobs/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &PinwheelJob{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
