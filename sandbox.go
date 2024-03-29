package rize

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Handles all Sandbox operations
type sandboxService service

// SandboxCreateParams are the body params used when creating a new Sandbox transaction
type SandboxCreateParams struct {
	TransactionType  string  `json:"transaction_type"`
	CustomerUID      string  `json:"customer_uid"`
	DebitCardUID     string  `json:"debit_card_uid"`
	DenialReason     string  `json:"denial_reason,omitempty"`
	USDollarAmount   float64 `json:"us_dollar_amount"`
	Mcc              string  `json:"mcc,omitempty"`
	MerchantLocation string  `json:"merchant_location,omitempty"`
	MerchantName     string  `json:"merchant_name,omitempty"`
	MerchantNumber   string  `json:"merchant_number,omitempty"`
	Description      string  `json:"description,omitempty"`
}

// SandboxResponse is an API response
type SandboxResponse struct {
	Success string `json:"success"`
}

// Create a Transaction by simulating the attributes that would be expected from reading an actual transaction received from a third party system
func (s *sandboxService) Create(ctx context.Context, params *SandboxCreateParams) (*SandboxResponse, error) {
	if params.TransactionType == "" ||
		params.CustomerUID == "" ||
		params.DebitCardUID == "" ||
		params.USDollarAmount == 0 {
		return nil, fmt.Errorf("TransactionType, CustomerUID, DebitCardUID and USDollarAmount are required")
	}

	bytesMessage, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := s.client.doRequest(ctx, http.MethodPost, "sandbox/mock_transactions", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SandboxResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
