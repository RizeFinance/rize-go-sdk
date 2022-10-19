package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Transfers
func ExampleTransferService_List(rc *rize.Client) {
	params := &rize.TransferListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		ExternalUID:         "client-generated-id",
		PoolUID:             "wTSMX1GubP21ev2h",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              10,
	}
	resp, err := rc.Transfers.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Transfers\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Transfers:", string(output))
}

// Create Transfer
func ExampleTransferService_Create(rc *rize.Client) {
	params := &rize.TransferCreateParams{
		ExternalUID:                    "partner-generated-id",
		SourceSyntheticAccountUID:      "4XkJnsfHsuqrxmeX",
		DestinationSyntheticAccountUID: "exMDShw6yM3NHLYV",
		InitiatingCustomerUID:          "iDtmSA52zRhgN4iy",
		USDTransferAmount:              "12.34",
	}
	resp, err := rc.Transfers.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Transfer\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Transfer:", string(output))
}

// Get Transfer
func ExampleTransferService_Get(rc *rize.Client) {
	resp, err := rc.Transfers.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Transfer\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Transfer:", string(output))
}
