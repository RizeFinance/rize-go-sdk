package main

import (
	"context"
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

	// List Transfers
	lp := rize.TransferListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		ExternalUID:         "client-generated-id",
		PoolUID:             "wTSMX1GubP21ev2h",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              0,
	}
	tl, err := rc.Transfers.List(context.Background(), &lp)
	if err != nil {
		log.Fatal("Error fetching Transfers\n", err)
	}
	output, _ := json.Marshal(tl)
	log.Println("List Transfers:", string(output))

	// Create Transfer
	cp := rize.TransferCreateParams{
		ExternalUID:                    "partner-generated-id",
		SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
		DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
		InitiatingCustomerUID:          "iDtmSA52zRhgN4iy",
		USDTransferAmount:              "12.34",
	}
	tc, err := rc.Transfers.Create(context.Background(), &cp)
	if err != nil {
		log.Fatal("Error creating Transfer\n", err)
	}
	output, _ = json.Marshal(tc)
	log.Println("Create Transfer:", string(output))

	// Get Transfer
	tg, err := rc.Transfers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Transfer\n", err)
	}
	output, _ = json.Marshal(tg)
	log.Println("Get Transfer:", string(output))
}
