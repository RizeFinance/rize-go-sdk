package rize

import (
	"context"
	"net/http"
	"testing"
	"time"
)

// Complete SyntheticAccount{} response data
var syntheticAccount = &SyntheticAccount{
	UID:                      "exMDShw6yM3NHLYV",
	ExternalUID:              "60689018-94e9-4870-970a-cc22f52c9c65",
	Name:                     "Spinach Fund",
	PoolUID:                  "wTSMX1GubP21ev2h",
	CustomerUID:              "h9MzupcjtA3LPW2e",
	SyntheticAccountTypeUID:  "fRMwt6H14ovFUz1s",
	SyntheticAccountCategory: "general",
	Status:                   "active",
	Liability:                true,
	NetUSDBalance:            "769.65",
	NetUSDPendingBalance:     "343.16",
	NetUSDAvailableBalance:   "701.46",
	AssetBalances: []*SyntheticAccountAssetBalance{{
		AssetQuantity:        "769.65",
		AssetType:            "USD",
		CurrentUSDValue:      "769.65",
		CustodialAccountUID:  "4uJMJjNd5wjzPaCj",
		CustodialAccountName: "Second Checking",
		Debit:                true,
	}},
	MasterAccount:               true,
	AccountNumber:               "123456789",
	AccountNumberLastFour:       "1234",
	RoutingNumber:               "0000000",
	OpenedAt:                    time.Now(),
	ClosedAt:                    time.Now(),
	ClosedToSyntheticAccountUID: "4XkJnsfHsuqrxmeX",
}

// Complete SyntheticAccountType{} response data
var syntheticAccountType = &SyntheticAccountType{
	UID:                      "EhrQZJNjCd79LLYq",
	Name:                     "New Resource Name",
	Description:              "This synthetic_account_type will be used to open synthetic_accounts for our customers that will only contain a USD asset type.",
	ProgramUID:               "kaxHFJnWvJxRJZxq",
	SyntheticAccountCategory: "general",
	TargetAnnualYieldPercent: 2.25,
}

func TestListSyntheticAccounts(t *testing.T) {
	params := &SyntheticAccountListParams{
		CustomerUID:              "uKxmLxUEiSj5h4M3",
		ExternalUID:              "client-generated-id",
		PoolUID:                  "wTSMX1GubP21ev2h",
		Limit:                    100,
		Offset:                   10,
		SyntheticAccountTypeUID:  "q4mdMxMtjXfdbrjn",
		SyntheticAccountCategory: "general",
		Liability:                true,
		Status:                   "active",
		Sort:                     "name_asc",
	}
	resp, err := rc.SyntheticAccounts.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Synthetic Accounts\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_accounts", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCreateSyntheticAccount(t *testing.T) {
	params := &SyntheticAccountCreateParams{
		ExternalUID:             "partner-generated-id",
		Name:                    "New Resource Name",
		PoolUID:                 "kaxHFJnWvJxRJZxq",
		SyntheticAccountTypeUID: "fRMwt6H14ovFUz1s",
		AccountNumber:           "123456789012",
		RoutingNumber:           "123456789",
		PlaidProcessorToken:     "processor-sandbox-96d86f35-ef58-4e4a-826f-4870b5d677f2",
		ExternalProcessorToken:  "processor-sandbox-96d86f35-ef58-4e4a-826f-4870b5d677f2",
	}
	resp, err := rc.SyntheticAccounts.Create(context.Background(), params)
	if err != nil {
		t.Fatal("Error creating Synthetic Account\n", err)
	}

	if err := validateSchema(http.MethodPost, "/synthetic_accounts", http.StatusCreated, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
func TestGetSyntheticAccount(t *testing.T) {
	resp, err := rc.SyntheticAccounts.Get(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		t.Fatal("Error fetching Synthetic Account\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_accounts/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUpdateSyntheticAccount(t *testing.T) {
	params := &SyntheticAccountUpdateParams{
		Name: "New Resource Name",
		Note: "note",
	}
	resp, err := rc.SyntheticAccounts.Update(context.Background(), "EhrQZJNjCd79LLYq", params)
	if err != nil {
		t.Fatal("Error updating Synthetic Account\n", err)
	}

	if err := validateSchema(http.MethodPut, "/synthetic_accounts/{uid}", http.StatusOK, nil, params, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDeleteSyntheticAccount(t *testing.T) {
	_, err := rc.SyntheticAccounts.Delete(context.Background(), "exMDShw6yM3NHLYV")
	if err != nil {
		t.Fatal("Error deleting Synthetic Account\n", err)
	}

	if err := validateSchema(http.MethodDelete, "/synthetic_accounts/{uid}", http.StatusNoContent, nil, nil, nil); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestListAccountTypes(t *testing.T) {
	params := &SyntheticAccountTypeListParams{
		ProgramUID: "EhrQZJNjCd79LLYq",
		Limit:      100,
		Offset:     10,
	}
	resp, err := rc.SyntheticAccounts.ListAccountTypes(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Synthetic Account Types\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_account_types", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetAccountType(t *testing.T) {
	resp, err := rc.SyntheticAccounts.GetAccountType(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Synthetic Account Type\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_account_types/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
