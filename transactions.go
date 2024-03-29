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

// Handles all Transaction operations
type transactionService service

// Transaction data type
type Transaction struct {
	AdjustmentUID                  string    `json:"adjustment_uid,omitempty"`
	CustomerUID                    string    `json:"customer_uid,omitempty"`
	CreatedAt                      time.Time `json:"created_at,omitempty"`
	CustodialAccountUIDs           []string  `json:"custodial_account_uids,omitempty"`
	DebitCardUID                   string    `json:"debit_card_uid,omitempty"`
	DenialReason                   string    `json:"denial_reason,omitempty"`
	Description                    string    `json:"description,omitempty"`
	DestinationSyntheticAccountUID string    `json:"destination_synthetic_account_uid,omitempty"`
	ID                             int       `json:"id,omitempty"`
	InitialActionAt                time.Time `json:"initial_action_at,omitempty"`
	MCC                            string    `json:"mcc,omitempty"`
	MerchantLocation               string    `json:"merchant_location,omitempty"`
	MerchantName                   string    `json:"merchant_name,omitempty"`
	MerchantNumber                 string    `json:"merchant_number,omitempty"`
	NetAsset                       string    `json:"net_asset,omitempty"`
	SettledAt                      time.Time `json:"settled_at,omitempty"`
	SettledIndex                   int       `json:"settled_index,omitempty"`
	SourceSyntheticAccountUID      string    `json:"source_synthetic_account_uid,omitempty"`
	Status                         string    `json:"status,omitempty"`
	TransactionEventUIDs           []string  `json:"transaction_event_uids,omitempty"`
	TransferUID                    string    `json:"transfer_uid,omitempty"`
	Type                           string    `json:"type,omitempty"`
	UID                            string    `json:"uid,omitempty"`
	USDollarAmount                 string    `json:"us_dollar_amount,omitempty"`
}

// TransactionEvent data type
type TransactionEvent struct {
	UID                            string    `json:"uid,omitempty"`
	SettledIndex                   int       `json:"settled_index,omitempty"`
	TransactionUIDs                []string  `json:"transaction_uids,omitempty"`
	SourceCustodialAccountUID      string    `json:"source_custodial_account_uid,omitempty"`
	DestinationCustodialAccountUID string    `json:"destination_custodial_account_uid,omitempty"`
	CustodialLineItemUIDs          []string  `json:"custodial_line_item_uids,omitempty"`
	Status                         string    `json:"status,omitempty"`
	USDollarAmount                 string    `json:"us_dollar_amount,omitempty"`
	Type                           string    `json:"type,omitempty"`
	DebitCardUID                   string    `json:"debit_card_uid,omitempty"`
	NetAsset                       string    `json:"net_asset,omitempty"`
	Description                    string    `json:"description,omitempty"`
	CreatedAt                      time.Time `json:"created_at,omitempty"`
	SettledAt                      time.Time `json:"settled_at,omitempty"`
}

// SyntheticLineItem data type
type SyntheticLineItem struct {
	UID                    string    `json:"uid,omitempty"`
	SettledIndex           int       `json:"settled_index,omitempty"`
	TransactionUID         string    `json:"transaction_uid,omitempty"`
	SyntheticAccountUID    string    `json:"synthetic_account_uid,omitempty"`
	Status                 string    `json:"status,omitempty"`
	USDollarAmount         string    `json:"us_dollar_amount,omitempty"`
	RunningUSDollarBalance string    `json:"running_us_dollar_balance,omitempty"`
	RunningAssetBalance    string    `json:"running_asset_balance,omitempty"`
	AssetQuantity          string    `json:"asset_quantity,omitempty"`
	AssetType              string    `json:"asset_type,omitempty"`
	ClosingPrice           string    `json:"closing_price,omitempty"`
	CustodialAccountUID    string    `json:"custodial_account_uid,omitempty"`
	CustodialAccountName   string    `json:"custodial_account_name,omitempty"`
	Description            string    `json:"description,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
	SettledAt              time.Time `json:"settled_at,omitempty"`
}

// CustodialLineItem data type
type CustodialLineItem struct {
	UID                    string    `json:"uid,omitempty"`
	SettledIndex           int       `json:"settled_index,omitempty"`
	TransactionUID         string    `json:"transaction_uid,omitempty"`
	TransactionEventUID    string    `json:"transaction_event_uid,omitempty"`
	CustodialAccountUID    string    `json:"custodial_account_uid,omitempty"`
	DebitCardUID           string    `json:"debit_card_uid,omitempty"`
	Status                 string    `json:"status,omitempty"`
	USDollarAmount         string    `json:"us_dollar_amount,omitempty"`
	RunningUSDollarBalance string    `json:"running_us_dollar_balance,omitempty"`
	RunningAssetBalance    string    `json:"running_asset_balance,omitempty"`
	AssetQuantity          string    `json:"asset_quantity,omitempty"`
	AssetType              string    `json:"asset_type,omitempty"`
	ClosingPrice           string    `json:"closing_price,omitempty"`
	Type                   string    `json:"type,omitempty"`
	Description            string    `json:"description,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
	OccurredAt             time.Time `json:"occurred_at,omitempty"`
	SettledAt              time.Time `json:"settled_at,omitempty"`
}

// TransactionListParams builds the query parameters used in querying Transactions
type TransactionListParams struct {
	CustomerUID                    string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	PoolUID                        string `url:"pool_uid,omitempty" json:"pool_uid,omitempty"`
	DebitCardUID                   string `url:"debit_card_uid,omitempty" json:"debit_card_uid,omitempty"`
	SourceSyntheticAccountUID      string `url:"source_synthetic_account_uid,omitempty" json:"source_synthetic_account_uid,omitempty"`
	DestinationSyntheticAccountUID string `url:"destination_synthetic_account_uid,omitempty" json:"destination_synthetic_account_uid,omitempty"`
	Type                           string `url:"type,omitempty" json:"type,omitempty"`
	SyntheticAccountUID            string `url:"synthetic_account_uid,omitempty" json:"synthetic_account_uid,omitempty"`
	ShowDeniedAuths                bool   `url:"show_denied_auths,omitempty" json:"show_denied_auths,omitempty"`
	ShowExpired                    bool   `url:"show_expired,omitempty" json:"show_expired,omitempty"`
	Status                         string `url:"status,omitempty" json:"status,omitempty"`
	SearchDescription              string `url:"search_description,omitempty" json:"search_description,omitempty"`
	IncludeZero                    bool   `url:"include_zero,omitempty" json:"include_zero,omitempty"`
	Limit                          int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset                         int    `url:"offset,omitempty" json:"offset,omitempty"`
	Sort                           string `url:"sort,omitempty" json:"sort,omitempty"`
}

// TransactionEventListParams builds the query parameters used in querying TransactionEvents
type TransactionEventListParams struct {
	SourceCustodialAccountUID      string `url:"source_custodial_account_uid,omitempty" json:"source_custodial_account_uid,omitempty"`
	DestinationCustodialAccountUID string `url:"destination_custodial_account_uid,omitempty" json:"destination_custodial_account_uid,omitempty"`
	CustodialAccountUID            string `url:"custodial_account_uid,omitempty" json:"custodial_account_uid,omitempty"`
	Type                           string `url:"type,omitempty" json:"type,omitempty"`
	TransactionUID                 string `url:"transaction_uid,omitempty" json:"transaction_uid,omitempty"`
	Limit                          int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset                         int    `url:"offset,omitempty" json:"offset,omitempty"`
	Sort                           string `url:"sort,omitempty" json:"sort,omitempty"`
}

// SyntheticLineItemListParams builds the query parameters used in querying SyntheticLineItems
type SyntheticLineItemListParams struct {
	CustomerUID         string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	PoolUID             string `url:"pool_uid,omitempty" json:"pool_uid,omitempty"`
	SyntheticAccountUID string `url:"synthetic_account_uid,omitempty" json:"synthetic_account_uid,omitempty"`
	Limit               int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset              int    `url:"offset,omitempty" json:"offset,omitempty"`
	TransactionUID      string `url:"transaction_uid,omitempty" json:"transaction_uid,omitempty"`
	Status              string `url:"status,omitempty" json:"status,omitempty"`
	Sort                string `url:"sort,omitempty" json:"sort,omitempty"`
}

// CustodialLineItemListParams builds the query parameters used in querying CustodialLineItems
type CustodialLineItemListParams struct {
	CustomerUID         string `url:"customer_uid,omitempty" json:"customer_uid,omitempty"`
	CustodialAccountUID string `url:"custodial_account_uid,omitempty" json:"custodial_account_uid,omitempty"`
	Status              string `url:"status,omitempty" json:"status,omitempty"`
	USDollarAmountMax   int    `url:"us_dollar_amount_max,omitempty" json:"us_dollar_amount_max,omitempty"`
	USDollarAmountMin   int    `url:"us_dollar_amount_min,omitempty" json:"us_dollar_amount_min,omitempty"`
	TransactionEventUID string `url:"transaction_event_uid,omitempty" json:"transaction_event_uid,omitempty"`
	TransactionUID      string `url:"transaction_uid,omitempty" json:"transaction_uid,omitempty"`
	Limit               int    `url:"limit,omitempty" json:"limit,omitempty"`
	Offset              int    `url:"offset,omitempty" json:"offset,omitempty"`
	Sort                string `url:"sort,omitempty" json:"sort,omitempty"`
}

// TransactionListResponse is an API response containing a list of Transactions
type TransactionListResponse struct {
	ListResponse
	Data []*Transaction `json:"data"`
}

// TransactionEventListResponse is an API response containing a list of TransactionEvents
type TransactionEventListResponse struct {
	ListResponse
	Data []*TransactionEvent `json:"data"`
}

// SyntheticLineItemListResponse is an API response containing a list of SyntheticLineItems
type SyntheticLineItemListResponse struct {
	ListResponse
	Data []*SyntheticLineItem `json:"data"`
}

// CustodialLineItemListResponse is an API response containing a list of CustodialLineItems
type CustodialLineItemListResponse struct {
	ListResponse
	Data []*CustodialLineItem `json:"data"`
}

// List retrieves a list of Transactions filtered by the given parameters
func (t *transactionService) List(ctx context.Context, params *TransactionListParams) (*TransactionListResponse, error) {
	// Build TransactionListParams into query string params
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, "transactions", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &TransactionListResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Transaction
func (t *transactionService) Get(ctx context.Context, uid string) (*Transaction, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("transactions/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Transaction{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListTransactionEvents retrieves a list of Transaction Events filtered by the given parameters
func (t *transactionService) ListTransactionEvents(ctx context.Context, params *TransactionEventListParams) (*TransactionEventListResponse, error) {
	// Build TransactionEventListParams into query string params
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, "transaction_events", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &TransactionEventListResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetTransactionEvent returns a single Transaction Event
func (t *transactionService) GetTransactionEvent(ctx context.Context, uid string) (*TransactionEvent, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("transaction_events/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &TransactionEvent{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListSyntheticLineItems retrieves a list of Synthetic Line Items filtered by the given parameters
func (t *transactionService) ListSyntheticLineItems(ctx context.Context, params *SyntheticLineItemListParams) (*SyntheticLineItemListResponse, error) {
	// Build SyntheticLineItemListParams into query string params
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, "synthetic_line_items", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticLineItemListResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetSyntheticLineItem returns a single Synthetic Line Item
func (t *transactionService) GetSyntheticLineItem(ctx context.Context, uid string) (*SyntheticLineItem, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("synthetic_line_items/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticLineItem{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ListCustodialLineItems retrieves a list of Custodial Line Items filtered by the given parameters
func (t *transactionService) ListCustodialLineItems(ctx context.Context, params *CustodialLineItemListParams) (*CustodialLineItemListResponse, error) {
	// Build CustodialLineItemListParams into query string params
	v, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, "custodial_line_items", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialLineItemListResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetCustodialLineItem returns a single Custodial Line Item
func (t *transactionService) GetCustodialLineItem(ctx context.Context, uid string) (*CustodialLineItem, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := t.client.doRequest(ctx, http.MethodGet, fmt.Sprintf("custodial_line_items/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialLineItem{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
