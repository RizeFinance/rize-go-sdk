package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Transactions
func (e Example) ExampleTransactionService_List(rc *rize.Client) {
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
		log.Fatal("Error fetching Transactions\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Transactions:", string(output))
}

// Get Transaction
func (e Example) ExampleTransactionService_Get(rc *rize.Client) {
	resp, err := rc.Transactions.Get(context.Background(), "SMwKC1osz77DTEiu")
	if err != nil {
		log.Fatal("Error fetching Transaction\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Transaction:", string(output))
}

// List Transaction Events
func (e Example) ExampleTransactionService_ListTransactionEvents(rc *rize.Client) {
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
		log.Fatal("Error fetching Transaction Events\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Transaction Events:", string(output))
}

// Get Transaction Event
func (e Example) ExampleTransactionService_GetTransactionEvent(rc *rize.Client) {
	resp, err := rc.Transactions.GetTransactionEvent(context.Background(), "MB2yqBrm3c4bUbou")
	if err != nil {
		log.Fatal("Error fetching Transaction Event\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Transaction Event:", string(output))
}

// List Synthetic Line Items
func (e Example) ExampleTransactionService_ListSyntheticLineItems(rc *rize.Client) {
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
		log.Fatal("Error fetching Synthetic Line Items\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Synthetic Line Items:", string(output))
}

// Get Synthetic Line Item
func (e Example) ExampleTransactionService_GetSyntheticLineItem(rc *rize.Client) {
	resp, err := rc.Transactions.GetSyntheticLineItem(context.Background(), "j56aHgLBqkNu1KwK")
	if err != nil {
		log.Fatal("Error fetching Synthetic Line Item\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Synthetic Line Item:", string(output))
}

// List Custodial Line Items
func (e Example) ExampleTransactionService_ListCustodialLineItems(rc *rize.Client) {
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
		log.Fatal("Error fetching Custodial Line Items\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Custodial Line Items:", string(output))
}

// Get Custodial Line Item
func (e Example) ExampleTransactionService_GetCustodialLineItem(rc *rize.Client) {
	resp, err := rc.Transactions.GetCustodialLineItem(context.Background(), "j56aHgLBqkNu1KwK")
	if err != nil {
		log.Fatal("Error fetching Custodial Line Item\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Custodial Line Item:", string(output))
}
