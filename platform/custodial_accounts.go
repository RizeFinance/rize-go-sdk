package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

// Handles all CustodialAccount operations
type custodialAccountService service

// CustodialAccount data type
type CustodialAccount struct {
	UID            string `json:"uid,omitempty"`
	ExternalUID    string `json:"external_uid,omitempty"`
	CustomerUID    string `json:"customer_uid,omitempty"`
	PoolUID        string `json:"pool_uid,omitempty"`
	Type           string `json:"type,omitempty"`
	Liability      bool   `json:"liability,omitempty"`
	Name           string `json:"name,omitempty"`
	PrimaryAccount bool   `json:"primary_account,omitempty"`
	Status         string `json:"status,omitempty"`
	AccountErrors  []struct {
		ErrorCode        string `json:"error_code,omitempty"`
		ErrorName        string `json:"error_name,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	} `json:"account_errors,omitempty"`
	NetUsdBalance          string `json:"net_usd_balance,omitempty"`
	NetUsdPendingBalance   string `json:"net_usd_pending_balance,omitempty"`
	NetUsdAvailableBalance string `json:"net_usd_available_balance,omitempty"`
	AssetBalances          []struct {
		AssetQuantity   string `json:"asset_quantity,omitempty"`
		AssetType       string `json:"asset_type,omitempty"`
		CurrentUsdValue string `json:"current_usd_value,omitempty"`
		Debit           bool   `json:"debit,omitempty"`
	} `json:"asset_balances,omitempty"`
	AccountNumber       string    `json:"account_number,omitempty"`
	AccountNumberMasked string    `json:"account_number_masked,omitempty"`
	RoutingNumber       string    `json:"routing_number,omitempty"`
	OpenedAt            time.Time `json:"opened_at,omitempty"`
	ClosedAt            time.Time `json:"closed_at,omitempty"`
}

// CustodialAccountListParams builds the query parameters used in querying Custodial Accounts
type CustodialAccountListParams struct {
	CustomerUID string `url:"customer_uid,omitempty"`
	ExternalUID string `url:"external_uid,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
	Liability   bool   `url:"liability,omitempty"`
	Type        string `url:"type,omitempty"`
}

// CustodialAccountResponse is an API response containing a list of Custodial Accounts
type CustodialAccountResponse struct {
	BaseResponse
	Data []*CustodialAccount `json:"data"`
}

// List retrieves a list of Custodial Accounts filtered by the given parameters
func (c *custodialAccountService) List(plp *CustodialAccountListParams) (*CustodialAccountResponse, error) {
	// Build CustodialAccountListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := c.rizeClient.doRequest(http.MethodGet, "custodial_accounts", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialAccountResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Custodial Account
func (c *custodialAccountService) Get(uid string) (*CustodialAccount, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := c.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("custodial_accounts/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &CustodialAccount{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}