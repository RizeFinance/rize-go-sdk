package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete CustodialAccount{} response data
var custodialAccount = &rize.CustodialAccount{
	UID:            "EhrQZJNjCd79LLYq",
	ExternalUID:    "partner-generated-id",
	CustomerUID:    "YrfDrfVRgpPgnhF5",
	PoolUID:        "kaxHFJnWvJxRJZxq",
	Type:           "dda",
	Liability:      true,
	Name:           "XYZ Checking Account",
	PrimaryAccount: true,
	Status:         "active",
	AccountErrors: []*rize.CustodialAccountError{{
		ErrorCode:        "FI1234",
		ErrorName:        "DOB does not match",
		ErrorDescription: "The given DOB does not match the known DOB for the SSN provided",
	}},
	NetUSDBalance:          "12.34",
	NetUSDPendingBalance:   "-2.56",
	NetUSDAvailableBalance: "9.78",
	AssetBalances: []*rize.CustodialAccountAssetBalance{{
		AssetQuantity:   "122.11",
		AssetType:       "USD",
		CurrentUSDValue: "122.11",
		Debit:           true,
	}},
	AccountNumber:       "123456789012",
	AccountNumberMasked: "9012",
	RoutingNumber:       "123456789",
	OpenedAt:            time.Now(),
	ClosedAt:            time.Now(),
}

func TestListCustodialAccounts(t *testing.T) {
	params := &rize.CustodialAccountListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
		Liability:   true,
		Type:        "dda",
	}
	resp, err := rc.CustodialAccounts.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Custodial Accounts\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_accounts", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCustodialAccounts(t *testing.T) {
	resp, err := rc.CustodialAccounts.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		t.Fatal("Error fetching Custodial Account\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_accounts/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
