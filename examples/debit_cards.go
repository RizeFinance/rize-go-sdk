package examples

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rizefinance/rize-go-sdk"
)

// List Debit Cards
func ExampleDebitCardService_List(rc *rize.Client) {
	params := &rize.DebitCardListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      10,
		PoolUID:     "wTSMX1GubP21ev2h",
		Locked:      true,
		Status:      "queued",
	}
	resp, err := rc.DebitCards.List(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Debit Cards\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("List Debit Cards:", string(output))
}

// Create Debit Card
func ExampleDebitCardService_Create(rc *rize.Client) {
	params := &rize.DebitCardCreateParams{
		ExternalUID:    "partner-generated-id",
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		CustomerUID:    "uKxmLxUEiSj5h4M3",
		PoolUID:        "wTSMX1GubP21ev2h",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.Create(context.Background(), params)
	if err != nil {
		log.Fatal("Error creating Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Create Debit Card:", string(output))
}

// Get Debit Card
func ExampleDebitCardService_Get(rc *rize.Client) {
	resp, err := rc.DebitCards.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Debit Card:", string(output))
}

// Activate Debit Card
func ExampleDebitCardService_Activate(rc *rize.Client) {
	params := &rize.DebitCardActivateParams{
		CardLastFourDigits: "1234",
		CVV:                "012",
		ExpiryDate:         "2023-08",
	}
	resp, err := rc.DebitCards.Activate(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error activating Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Activate Debit Card:", string(output))
}

// Lock Debit Card
func ExampleDebitCardService_Lock(rc *rize.Client) {
	params := &rize.DebitCardLockParams{
		LockReason: "Fraud detected",
	}
	resp, err := rc.DebitCards.Lock(context.Background(), "Lt6qjTNnYLjFfEWL", params)
	if err != nil {
		log.Fatal("Error locking Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Lock Debit Card:", string(output))
}

// Unlock Debit Card
func ExampleDebitCardService_Unlock(rc *rize.Client) {
	resp, err := rc.DebitCards.Unlock(context.Background(), "Lt6qjTNnYLjFfEWL")
	if err != nil {
		log.Fatal("Error unlocking Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Unlock Debit Card:", string(output))
}

// Reissue Debit Card
func ExampleDebitCardService_Reissue(rc *rize.Client) {
	params := &rize.DebitCardReissueParams{
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		ReissueReason:  "damaged",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.Reissue(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error reissuing Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Reissue Debit Card:", string(output))
}

// Get Debit Card PIN Token
func ExampleDebitCardService_GetPINToken(rc *rize.Client) {
	params := &rize.DebitCardGetPINTokenParams{
		ForceReset: true,
	}
	resp, err := rc.DebitCards.GetPINToken(context.Background(), "Lt6qjTNnYLjFfEWL", params)
	if err != nil {
		log.Fatal("Error fetching Debit Card PIN Token\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Debit Card PIN Token:", string(output))
}

// Get Debit Card Access Token
func ExampleDebitCardService_GetAccessToken(rc *rize.Client) {
	resp, err := rc.DebitCards.GetAccessToken(context.Background(), "Lt6qjTNnYLjFfEWL")
	if err != nil {
		log.Fatal("Error fetching Debit Card Access Token\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Get Debit Card Access Token:", string(output))
}

// Migrate a Virtual Debit Card to a Physical Debit Card
func ExampleDebitCardService_MigrateVirtualDebitCard(rc *rize.Client) {
	params := &rize.VirtualDebitCardMigrateParams{
		ExternalUID:    "partner-generated-id",
		CardArtworkUID: "EhrQZJNjCd79LLYq",
		ShippingAddress: &rize.DebitCardShippingAddress{
			Street1:    "123 Abc St",
			Street2:    "Apt 2",
			City:       "Chicago",
			State:      "IL",
			PostalCode: "12345",
		},
	}
	resp, err := rc.DebitCards.MigrateVirtualDebitCard(context.Background(), "h9MzupcjtA3LPW2e", params)
	if err != nil {
		log.Fatal("Error migrating Debit Card\n", err)
	}
	output, _ := json.MarshalIndent(resp, "", "\t")
	log.Println("Migrating Virtual Debit Card:", string(output))
}

// Get Virtual Debit Card Image
func ExampleDebitCardService_GetVirtualDebitCardImage(rc *rize.Client) {
	params := &rize.DebitCardAccessToken{
		Token:    "VmU27goFku4DyxfsdyoH5G1mlztvwskBywKrskVN9jQOh50Yy7",
		ConfigID: "1",
	}
	resp, err := rc.DebitCards.GetVirtualDebitCardImage(context.Background(), params)
	if err != nil {
		log.Fatal("Error fetching Virtual Debit Card Image\n", err)
	}
	log.Println("Get Virtual Debit Card Image:", resp.Status)
}
