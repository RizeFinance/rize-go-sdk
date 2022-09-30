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

	// List Debit Cards
	dlp := rize.DebitCardListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		ExternalUID: "client-generated-id",
		Limit:       100,
		Offset:      0,
		PoolUID:     "wTSMX1GubP21ev2h",
		Locked:      false,
		Status:      "queued",
	}
	dl, err := rc.DebitCards.List(&dlp)
	if err != nil {
		log.Fatal("Error fetching Debit Cards\n", err)
	}
	output, _ := json.Marshal(dl)
	log.Println("List Debit Cards:", string(output))

	// Create Debit Card
	dcp := rize.DebitCardCreateParams{
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
	dc, err := rc.DebitCards.Create(&dcp)
	if err != nil {
		log.Fatal("Error creating Debit Card\n", err)
	}
	output, _ = json.Marshal(dc)
	log.Println("Create Debit Card:", string(output))

	// Get Debit Card
	dg, err := rc.DebitCards.Get("EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Debit Card\n", err)
	}
	output, _ = json.Marshal(dg)
	log.Println("Get Debit Card:", string(output))

	// Activate Debit Card
	dap := rize.DebitCardActivateParams{
		CardLastFourDigits: "1234",
		CVV:                "012",
		ExpiryDate:         "2023-08",
	}
	da, err := rc.DebitCards.Activate("h9MzupcjtA3LPW2e", &dap)
	if err != nil {
		log.Fatal("Error activating Debit Card\n", err)
	}
	output, _ = json.Marshal(da)
	log.Println("Activate Debit Card:", string(output))

	// Lock Debit Card
	dlk, err := rc.DebitCards.Lock("Lt6qjTNnYLjFfEWL", "lost")
	if err != nil {
		log.Fatal("Error locking Debit Card\n", err)
	}
	output, _ = json.Marshal(dlk)
	log.Println("Lock Debit Card:", string(output))

	// Unlock Debit Card
	duk, err := rc.DebitCards.Unlock("Lt6qjTNnYLjFfEWL")
	if err != nil {
		log.Fatal("Error unlocking Debit Card\n", err)
	}
	output, _ = json.Marshal(duk)
	log.Println("Unlock Debit Card:", string(output))

	// Reissue Debit Card
	drp := rize.DebitCardReissueParams{
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
	dr, err := rc.DebitCards.Reissue("h9MzupcjtA3LPW2e", &drp)
	if err != nil {
		log.Fatal("Error reissuing Debit Card\n", err)
	}
	output, _ = json.Marshal(dr)
	log.Println("Reissue Debit Card:", string(output))

	// Get Debit Card PIN Token
	dpt, err := rc.DebitCards.GetPINToken("Lt6qjTNnYLjFfEWL", true)
	if err != nil {
		log.Fatal("Error fetching Debit Card PIN Token\n", err)
	}
	output, _ = json.Marshal(dpt)
	log.Println("Get Debit Card PIN Token:", string(output))

	// Get Debit Card Access Token
	dat, err := rc.DebitCards.GetAccessToken("Lt6qjTNnYLjFfEWL")
	if err != nil {
		log.Fatal("Error fetching Debit Card Access Token\n", err)
	}
	output, _ = json.Marshal(dat)
	log.Println("Get Debit Card Access Token:", string(output))

	// Migrate a Virtual Debit Card to a Physical Debit Card
	vmp := rize.VirtualDebitCardMigrateParams{
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
	vm, err := rc.DebitCards.MigrateVirtualDebitCard("h9MzupcjtA3LPW2e", &vmp)
	if err != nil {
		log.Fatal("Error migrating Debit Card\n", err)
	}
	output, _ = json.Marshal(vm)
	log.Println("Migrating Virtual Debit Card:", string(output))

	// Get Virtual Debit Card Image
	vi, err := rc.DebitCards.GetVirtualDebitCardImage("Lt6qjTNnYLjFfEWL", "h9MzupcjtA3LPW2e")
	if err != nil {
		log.Fatal("Error fetching Virtual Debit Card Image\n", err)
	}
	output, _ = json.Marshal(vi)
	log.Println("Get Virtual Debit Card Image:", string(output))

}
