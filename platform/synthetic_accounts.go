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

// Handles all Synthetic Account operations
type syntheticAccountService service

// SyntheticAccount data type
type SyntheticAccount struct {
	UID                      string `json:"uid,omitempty"`
	ExternalUID              string `json:"external_uid,omitempty"`
	Name                     string `json:"name,omitempty"`
	PoolUID                  string `json:"pool_uid,omitempty"`
	CustomerUID              string `json:"customer_uid,omitempty"`
	SyntheticAccountTypeUID  string `json:"synthetic_account_type_uid,omitempty"`
	SyntheticAccountCategory string `json:"synthetic_account_category,omitempty"`
	Status                   string `json:"status,omitempty"`
	Liability                bool   `json:"liability,omitempty"`
	NetUsdBalance            string `json:"net_usd_balance,omitempty"`
	NetUsdPendingBalance     string `json:"net_usd_pending_balance,omitempty"`
	NetUsdAvailableBalance   string `json:"net_usd_available_balance,omitempty"`
	AssetBalances            []struct {
		AssetQuantity        string `json:"asset_quantity,omitempty"`
		AssetType            string `json:"asset_type,omitempty"`
		CurrentUsdValue      string `json:"current_usd_value,omitempty"`
		CustodialAccountUID  string `json:"custodial_account_uid,omitempty"`
		CustodialAccountName string `json:"custodial_account_name,omitempty"`
		Debit                bool   `json:"debit,omitempty"`
	} `json:"asset_balances,omitempty"`
	MasterAccount               bool        `json:"master_account,omitempty"`
	AccountNumber               string      `json:"account_number,omitempty"`
	AccountNumberLastFour       string      `json:"account_number_last_four,omitempty"`
	RoutingNumber               string      `json:"routing_number,omitempty"`
	OpenedAt                    time.Time   `json:"opened_at,omitempty"`
	ClosedAt                    interface{} `json:"closed_at,omitempty"`
	ClosedToSyntheticAccountUID interface{} `json:"closed_to_synthetic_account_uid,omitempty"`
}

// SyntheticAccountType data type
type SyntheticAccountType struct {
	UID                      string  `json:"uid,omitempty"`
	Name                     string  `json:"name,omitempty"`
	Description              string  `json:"description,omitempty"`
	ProgramUID               string  `json:"program_uid,omitempty"`
	SyntheticAccountCategory string  `json:"synthetic_account_category,omitempty"`
	TargetAnnualYieldPercent float64 `json:"target_annual_yield_percent,omitempty"`
}

// SyntheticAccountListParams builds the query parameters used in querying Synthetic Accounts
type SyntheticAccountListParams struct {
	CustomerUID              string `url:"customer_uid,omitempty"`
	ExternalUID              string `url:"external_uid,omitempty"`
	PoolUID                  string `url:"pool_uid,omitempty"`
	Limit                    int    `url:"limit,omitempty"`
	Offset                   int    `url:"offset,omitempty"`
	SyntheticAccountTypeUID  string `url:"synthetic_account_type_uid,omitempty"`
	SyntheticAccountCategory string `url:"synthetic_account_category,omitempty"`
	Liability                bool   `url:"liability,omitempty"`
	Status                   string `url:"status,omitempty"`
	Sort                     string `url:"sort,omitempty"`
}

// SyntheticAccountCreateParams are the body params used when creating a new Synthetic Account
type SyntheticAccountCreateParams struct {
	ExternalUID             string `json:"external_uid,omitempty"`
	Name                    string `json:"name"`
	PoolUID                 string `json:"pool_uid"`
	SyntheticAccountTypeUID string `json:"synthetic_account_type_uid"`
	AccountNumber           string `json:"account_number,omitempty"`
	RoutingNumber           string `json:"routing_number,omitempty"`
	ExternalProcessorToken  string `json:"external_processor_token,omitempty"`
}

// SyntheticAccountUpdateParams are the body params used when updating a Synthetic Account
type SyntheticAccountUpdateParams struct {
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
}

// SyntheticAccountTypeListParams builds the query parameters used in querying Synthetic Account Types
type SyntheticAccountTypeListParams struct {
	ProgramUID string `url:"program_uid,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

// SyntheticAccountResponse is an API response containing a list of Synthetic Accounts
type SyntheticAccountResponse struct {
	BaseResponse
	Data []*SyntheticAccount `json:"data"`
}

// SyntheticAccountTypeResponse is an API response containing a list of Synthetic Account Types
type SyntheticAccountTypeResponse struct {
	BaseResponse
	Data []*SyntheticAccountType `json:"data"`
}

// List retrieves a list of Synthetic Account filtered by the given parameters
func (sa *syntheticAccountService) List(plp *SyntheticAccountListParams) (*SyntheticAccountResponse, error) {
	// Build SyntheticAccountListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := sa.rizeClient.doRequest(http.MethodGet, "synthetic_accounts", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccountResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Create a new Synthetic Account in the Pool with the provided specification
func (sa *syntheticAccountService) Create(sac *SyntheticAccountCreateParams) (*SyntheticAccount, error) {
	if sac.Name == "" || sac.PoolUID == "" || sac.SyntheticAccountTypeUID == "" {
		return nil, fmt.Errorf("Name, PoolUID and SyntheticAccountTypeUID are required")
	}

	bytesMessage, err := json.Marshal(sac)
	if err != nil {
		return nil, err
	}

	res, err := sa.rizeClient.doRequest(http.MethodPost, "synthetic_accounts", nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccount{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Get returns a single Synthetic Account resource along with supporting details and account balances
func (sa *syntheticAccountService) Get(uid string) (*SyntheticAccount, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := sa.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("synthetic_accounts/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccount{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Update the Synthetic Account metadata
func (sa *syntheticAccountService) Update(uid string, su *SyntheticAccountUpdateParams) (*SyntheticAccount, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	bytesMessage, err := json.Marshal(su)
	if err != nil {
		return nil, err
	}

	res, err := sa.rizeClient.doRequest(http.MethodPut, fmt.Sprintf("synthetic_accounts/%s", uid), nil, bytes.NewBuffer(bytesMessage))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccount{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// Delete will archive a Synthetic Account
func (sa *syntheticAccountService) Delete(uid string) (*http.Response, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := sa.rizeClient.doRequest(http.MethodDelete, fmt.Sprintf("synthetic_accounts/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

// ListAccountTypes retrieves a list of Synthetic Account Types filtered by the given parameters
func (sa *syntheticAccountService) ListAccountTypes(plp *SyntheticAccountTypeListParams) (*SyntheticAccountTypeResponse, error) {
	// Build SyntheticAccountTypeListParams into query string params
	v, err := query.Values(plp)
	if err != nil {
		return nil, err
	}

	res, err := sa.rizeClient.doRequest(http.MethodGet, "synthetic_account_types", v, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccountTypeResponse{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAccountType returns a single Synthetic Account Type resource along with supporting details
func (sa *syntheticAccountService) GetAccountType(uid string) (*SyntheticAccountType, error) {
	if uid == "" {
		return nil, fmt.Errorf("UID is required")
	}

	res, err := sa.rizeClient.doRequest(http.MethodGet, fmt.Sprintf("synthetic_account_types/%s", uid), nil, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &SyntheticAccountType{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil
}
