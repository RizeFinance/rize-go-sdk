package main

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk/internal"
	rize "github.com/rizefinance/rize-go-sdk/platform"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.RizeConfig{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewRizeClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List Transactions
	tlp := rize.TransactionListParams{
		CustomerUID:                    "uKxmLxUEiSj5h4M3",
		PoolUID:                        "wTSMX1GubP21ev2h",
		DebitCardUID:                   "MYNGv45UK6HWBHHf",
		SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
		DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
		SyntheticAccountUID:            "4XkJnsfHsuqrxmeX",
		Type:                           "card_refund",
		ShowDeniedAuths:                false,
		ShowExpired:                    true,
		Status:                         "failed",
		SearchDescription:              "Transfer%2A",
		IncludeZero:                    true,
		Limit:                          100,
		Offset:                         0,
		Sort:                           "id_asc",
	}
	tl, err := rc.Transactions.List(&tlp)
	if err != nil {
		log.Fatal("Error fetching transactions\n", err)
	}
	output, _ := json.Marshal(tl)
	log.Println("List Transactions:", string(output))

	// Get Transaction
	tg, err := rc.Transactions.Get("SMwKC1osz77DTEiu")
	if err != nil {
		log.Fatal("Error fetching transaction\n", err)
	}
	output, _ = json.Marshal(tg)
	log.Println("Get Transaction:", string(output))

	// List Transaction Events
	tep := rize.TransactionEventListParams{
		SourceCustodialAccountUID:      "dmRtw1xkS9ghrntB",
		DestinationCustodialAccountUID: "W55zKgvAk3zkpGM3",
		CustodialAccountUID:            "dmRtw1xkS9ghrntB",
		Type:                           "odfi_ach_withdrawal",
		TransactionUID:                 "SMwKC1osz77DTEiu",
		Limit:                          100,
		Offset:                         0,
		Sort:                           "created_at_asc",
	}
	te, err := rc.Transactions.ListTransactionEvents(&tep)
	if err != nil {
		log.Fatal("Error fetching Transaction Events\n", err)
	}
	output, _ = json.Marshal(te)
	log.Println("List Transaction Events:", string(output))

	// Get Transaction Event
	teg, err := rc.Transactions.GetTransactionEvents("MB2yqBrm3c4bUbou")
	if err != nil {
		log.Fatal("Error fetching Transaction Event\n", err)
	}
	output, _ = json.Marshal(teg)
	log.Println("Get Transaction Event:", string(output))

	// List Synthetic Line Items
	slp := rize.SyntheticLineItemListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		PoolUID:             "wTSMX1GubP21ev2h",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              0,
		TransactionUID:      "SMwKC1osz77DTEiu",
		Status:              "in_progress",
		Sort:                "created_at_asc",
	}
	sl, err := rc.Transactions.ListSyntheticLineItems(&slp)
	if err != nil {
		log.Fatal("Error fetching Synthetic Line Items\n", err)
	}
	output, _ = json.Marshal(sl)
	log.Println("List Synthetic Line Items:", string(output))

	// Get Synthetic Line Item
	sg, err := rc.Transactions.GetSyntheticLineItems("j56aHgLBqkNu1KwK")
	if err != nil {
		log.Fatal("Error fetching Synthetic Line Item\n", err)
	}
	output, _ = json.Marshal(sg)
	log.Println("Get Synthetic Line Item:", string(output))

	// List Custodial Line Items
	clp := rize.CustodialLineItemListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		CustodialAccountUID: "wTSMX1GubP21ev2h",
		Status:              "voided",
		USDollarAmountMax:   2,
		USDollarAmountMin:   2,
		TransactionEventUID: "MB2yqBrm3c4bUbou",
		TransactionUID:      "SMwKC1osz77DTEiu",
		Limit:               100,
		Offset:              0,
		Sort:                "created_at_asc",
	}
	cl, err := rc.Transactions.ListCustodialLineItems(&clp)
	if err != nil {
		log.Fatal("Error fetching Custodial Line Items\n", err)
	}
	output, _ = json.Marshal(cl)
	log.Println("List Custodial Line Items:", string(output))

	// Get Custodial Line Item
	cg, err := rc.Transactions.GetCustodialLineItems("j56aHgLBqkNu1KwK")
	if err != nil {
		log.Fatal("Error fetching Custodial Line Item\n", err)
	}
	output, _ = json.Marshal(cg)
	log.Println("Get Custodial Line Item:", string(output))
}
