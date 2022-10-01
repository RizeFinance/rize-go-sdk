package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all Transfer operations
type transferService service

// Transfer data type
type Transfer struct {
	UID                            string    `json:"uid,omitempty"`
	ExternalUID                    string    `json:"external_uid,omitempty"`
	SourceSyntheticAccountUID      string    `json:"source_synthetic_account_uid,omitempty"`
	DestinationSyntheticAccountUID string    `json:"destination_synthetic_account_uid,omitempty"`
	InitiatingCustomerUID          string    `json:"initiating_customer_uid,omitempty"`
	Status                         string    `json:"status,omitempty"`
	CreatedAt                      time.Time `json:"created_at,omitempty"`
	USDTransferAmount              string    `json:"usd_transfer_amount,omitempty"`
	USDRequestedAmount             string    `json:"usd_requested_amount,omitempty"`
}

// TransferListParams builds the query parameters used in querying Transfers
type TransferListParams struct {
	CustomerUID         string `url:"customer_uid,omitempty"`
	ExternalUID         string `url:"external_uid,omitempty"`
	PoolUID             string `url:"pool_uid,omitempty"`
	SyntheticAccountUID string `url:"synthetic_account_uid,omitempty"`
	Limit               int    `url:"limit,omitempty"`
	Offset              int    `url:"offset,omitempty"`
}

// TransferCreateParams are the body params used when creating a new Transfer
type TransferCreateParams struct {
	ExternalUID                    string `json:"external_uid,omitempty"`
	SourceSyntheticAccountUID      string `json:"source_synthetic_account_uid"`
	DestinationSyntheticAccountUID string `json:"destination_synthetic_account_uid"`
	InitiatingCustomerUID          string `json:"initiating_customer_uid"`
	USDTransferAmount              string `json:"usd_transfer_amount"`
}

// TransferResponse is an API response containing a list of Transfers
type TransferResponse struct {
	BaseResponse
	Data []*Transfer `json:"data"`
}

// List retrieves a list of Transfers filtered by the given parameters
func (t *transferService) List(ctx context.Context, tlp *TransferListParams) (*TransferResponse, error) {
	// Build TransferListParams into query string params
	v, err := query.Values(tlp)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, "transfers", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &TransferResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create will initiate a Transfer between two Synthetic Accounts
func (t *transferService) Create(ctx context.Context, tc *TransferCreateParams) (*Transfer, error) {
	if tc.SourceSyntheticAccountUID == "" ||
		tc.DestinationSyntheticAccountUID == "" ||
		tc.InitiatingCustomerUID == "" ||
		tc.USDTransferAmount == "" {
		return nil, fmt.Errorf("SourceSyntheticAccountUID, DestinationSyntheticAccountUID, InitiatingCustomerUID and USDTransferAmount are required")
	}

	bytesMessage, err := json.Marshal(tc)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodPost, "transfers", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Transfer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Transfer
func (t *transferService) Get(ctx context.Context, uid string) (*Transfer, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("transfers/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Transfer{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
