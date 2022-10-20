package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// Create Sandbox Transaction
func (e Example) ExampleSandboxService_Create(rc *rize.Client) {
	params := &rize.SandboxCreateParams{
		TransactionType:  "atm_withdrawal",
		CustomerUID:      "uKxmLxUEiSj5h4M3",
		DebitCardUID:     "h9MzupcjtA3LPW2e",
		DenialReason:     "insufficient_funds",
		USDollarAmount:   21.89,
		Mcc:              "5200",
		MerchantLocation: "NEW YORK, NY",
		MerchantName:     "Widgets Incorporated",
		MerchantNumber:   "000067107015968",
		Description:      "test transaction",
	}
	resp, err := rc.Sandbox.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Sandbox transactions\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Sandbox Transactions:", string(output))
}
