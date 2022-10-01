package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Handles all Custodial Partner operations
type custodialPartnerService service

// CustodialPartner data type
type CustodialPartner struct {
	UID  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// CustodialPartnerResponse is an API response containing a list of Custodial Partners
type CustodialPartnerResponse struct {
	BaseResponse
	Data []*CustodialPartner `json:"data"`
}

// List retrieves a list of CustodialPartners filtered by the given parameters
func (c *custodialPartnerService) List(ctx context.Context) (*CustodialPartnerResponse, error) {
	res, err := c.client.doRequest(ctx, http.MethodGet, "custodial_partners", nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialPartnerResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single CustodialPartner
func (c *custodialPartnerService) Get(ctx context.Context, uid string) (*CustodialPartner, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := c.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("custodial_partners/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialPartner{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
