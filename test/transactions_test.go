package rize_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/rizefinance/rize-go-sdk"
)

// Complete Transaction{} response data
var transaction = &rize.Transaction{
	AdjustmentUID:                  "KM2eKbR98t4tdAyZ",
	CustomerUID:                    "Trzqy9t6j6tFGoG3",
	CreatedAt:                      time.Now(),
	CustodialAccountUIDs:           []string{"Gw4gr1T81YrvLT6M", "66gso857LyQWC6XN"},
	DebitCardUID:                   "h9MzupcjtA3LPW2e",
	DenialReason:                   "insufficient_funds",
	Description:                    "External transfer from \"Bank ABC\" to \"Emergency Cupcake Funds\"",
	DestinationSyntheticAccountUID: "15gTLPAEzHxtSrU6",
	ID:                             27,
	InitialActionAt:                time.Now(),
	MCC:                            "5200",
	MerchantLocation:               "SPRINGFIELD VA",
	MerchantName:                   "BOB'S WIDGETS",
	MerchantNumber:                 "010000034261221",
	NetAsset:                       "positive",
	SettledAt:                      time.Now(),
	SettledIndex:                   8,
	SourceSyntheticAccountUID:      "rJuYjZuLei2TZji9",
	Status:                         "settled",
	TransactionEventUIDs:           []string{"MB2yqBrm3c4bUbou"},
	TransferUID:                    "1qVSAEjsV55vDZxX",
	Type:                           "external_transfer",
	UID:                            "SMwKC1osz77DTEiu",
	USDollarAmount:                 "5.21",
}

// Complete TransactionEvent{} response data
var transactionEvent = &rize.TransactionEvent{
	UID:                            "MB2yqBrm3c4bUbou",
	SettledIndex:                   1,
	TransactionUIDs:                []string{"SMwKC1osz77DTEiu", "Wfp6uBnhKDxUtCeu"},
	SourceCustodialAccountUID:      "66gso857LyQWC6XN",
	DestinationCustodialAccountUID: "Gw4gr1T81YrvLT6M",
	CustodialLineItemUIDs:          []string{"y4r8oTATb23MdGDF", "m9bED9iicUUk8YAc"},
	Status:                         "settled",
	USDollarAmount:                 "5.21",
	Type:                           "odfi_ach_deposit",
	DebitCardUID:                   "h9MzupcjtA3LPW2e",
	NetAsset:                       "positive",
	Description:                    "Transfer from Bank ABC 123 to Rize 999",
	CreatedAt:                      time.Now(),
	SettledAt:                      time.Now(),
}

// Complete SyntheticLineItem{} response data
var syntheticLineItem = &rize.SyntheticLineItem{
	UID:                    "j56aHgLBqkNu1KwK",
	SettledIndex:           10,
	TransactionUID:         "YBHNH3BgykqrjLgz",
	SyntheticAccountUID:    "exMDShw6yM3NHLYV",
	Status:                 "settled",
	USDollarAmount:         "-12.34",
	RunningUSDollarBalance: "4.21",
	RunningAssetBalance:    "4.21",
	AssetQuantity:          "-12.34",
	AssetType:              "USD",
	ClosingPrice:           "12.50",
	CustodialAccountUID:    "WPsSKYyKy9v4RQ6P",
	CustodialAccountName:   "Checking Account",
	Description:            "Deposit from Bank ABC 123",
	CreatedAt:              time.Now(),
	SettledAt:              time.Now(),
}

// Complete CustodialLineItem{} response data
var custodialLineItem = &rize.CustodialLineItem{
	UID:                    "y4r8oTATb23MdGDF",
	SettledIndex:           1,
	TransactionUID:         "SMwKC1osz77DTEiu",
	TransactionEventUID:    "MB2yqBrm3c4bUbou",
	CustodialAccountUID:    "Gw4gr1T81YrvLT6M",
	DebitCardUID:           "h9MzupcjtA3LPW2e",
	Status:                 "settled",
	USDollarAmount:         "5.21",
	RunningUSDollarBalance: "34.21",
	RunningAssetBalance:    "34.21",
	AssetQuantity:          "5.21",
	AssetType:              "USD",
	ClosingPrice:           "5.00",
	Type:                   "transaction category 8348",
	Description:            "Deposit from Bank ABC 123",
	CreatedAt:              time.Now(),
	OccurredAt:             time.Now(),
	SettledAt:              time.Now(),
}

func TestListTransactions(t *testing.T) {
	params := &rize.TransactionListParams{
		CustomerUID:                    "uKxmLxUEiSj5h4M3",
		PoolUID:                        "wTSMX1GubP21ev2h",
		DebitCardUID:                   "MYNGv45UK6HWBHHf",
		SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
		DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
		SyntheticAccountUID:            "4XkJnsfHsuqrxmeX",
		Type:                           "card_refund",
		ShowDeniedAuths:                true,
		ShowExpired:                    true,
		Status:                         "failed",
		SearchDescription:              "Transfer%2A",
		IncludeZero:                    true,
		Limit:                          100,
		Offset:                         10,
		Sort:                           "id_asc",
	}
	resp, err := rc.Transactions.List(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Transactions\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transactions", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetTransaction(t *testing.T) {
	resp, err := rc.Transactions.Get(context.Background(), "SMwKC1osz77DTEiu")
	if err != nil {
		t.Fatal("Error fetching Transaction\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transactions/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestListTransactionEvents(t *testing.T) {
	params := &rize.TransactionEventListParams{
		SourceCustodialAccountUID:      "dmRtw1xkS9ghrntB",
		DestinationCustodialAccountUID: "W55zKgvAk3zkpGM3",
		CustodialAccountUID:            "dmRtw1xkS9ghrntB",
		Type:                           "odfi_ach_withdrawal",
		TransactionUID:                 "SMwKC1osz77DTEiu",
		Limit:                          100,
		Offset:                         10,
		Sort:                           "created_at_asc",
	}
	resp, err := rc.Transactions.ListTransactionEvents(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Transaction Events\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transaction_events", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetTransactionEvent(t *testing.T) {
	resp, err := rc.Transactions.GetTransactionEvent(context.Background(), "MB2yqBrm3c4bUbou")
	if err != nil {
		t.Fatal("Error fetching Transaction Event\n", err)
	}

	if err := validateSchema(http.MethodGet, "/transaction_events/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestListSyntheticLineItems(t *testing.T) {
	params := &rize.SyntheticLineItemListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		PoolUID:             "wTSMX1GubP21ev2h",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
		TransactionUID:      "SMwKC1osz77DTEiu",
		Status:              "in_progress",
		Sort:                "created_at_asc",
	}
	resp, err := rc.Transactions.ListSyntheticLineItems(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Synthetic Line Items\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_line_items", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetSyntheticLineItem(t *testing.T) {
	resp, err := rc.Transactions.GetSyntheticLineItem(context.Background(), "j56aHgLBqkNu1KwK")
	if err != nil {
		t.Fatal("Error fetching Synthetic Line Item\n", err)
	}

	if err := validateSchema(http.MethodGet, "/synthetic_line_items/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestListCustodialLineItems(t *testing.T) {
	params := &rize.CustodialLineItemListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		CustodialAccountUID: "wTSMX1GubP21ev2h",
		Status:              "voided",
		USDollarAmountMax:   2,
		USDollarAmountMin:   2,
		TransactionEventUID: "MB2yqBrm3c4bUbou",
		TransactionUID:      "SMwKC1osz77DTEiu",
		Limit:               100,
		Offset:              10,
		Sort:                "created_at_asc",
	}
	resp, err := rc.Transactions.ListCustodialLineItems(context.Background(), params)
	if err != nil {
		t.Fatal("Error fetching Custodial Line Items\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_line_items", http.StatusOK, params, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetCustodialLineItem(t *testing.T) {
	resp, err := rc.Transactions.GetCustodialLineItem(context.Background(), "j56aHgLBqkNu1KwK")
	if err != nil {
		t.Fatal("Error fetching Custodial Line Item\n", err)
	}

	if err := validateSchema(http.MethodGet, "/custodial_line_items/{uid}", http.StatusOK, nil, nil, resp); err != nil {
		t.Fatalf(err.Error())
	}
}
