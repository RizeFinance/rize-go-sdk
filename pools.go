package rize

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// Handles all Pool operations
type poolService service

// Pool data type
type Pool struct {
	UID              string   `json:"uid,omitempty"`
	Name             string   `json:"name,omitempty"`
	OwnerCustomerUID string   `json:"owner_customer_uid,omitempty"`
	CustomerUIDs     []string `json:"customer_uids,omitempty"`
}

// PoolListParams builds the query parameters used in querying Pools
type PoolListParams struct {
	CustomerUID string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	ExternalUID string `url:"external_uid,omitempty" json:"external_uid,omitempty"`
	Limit       int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty" json:"offset,omitempty"`
}

// PoolResponse is an API response containing a list of Pools
type PoolResponse struct {
	BaseResponse
	Data []*Pool `json:"data"`
}

// List retrieves a list of Pools filtered by the given parameters
func (p *poolService) List(ctx context.Context, plp *PoolListParams) (*PoolResponse, error) {
	// Build PoolListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, "pools", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &PoolResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Pool
func (p *poolService) Get(ctx context.Context, uid string) (*Pool, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := p.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("pools/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Pool{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
