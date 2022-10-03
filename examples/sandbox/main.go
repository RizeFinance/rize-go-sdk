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

	// Create Sandbox Transaction
	scp := rize.SandboxCreateParams{
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
	sc, err := rc.Sandbox.Create(context.Background(), &scp)
	if err != nil {
		log.Fatal("Error creating Sandbox transactions\n", err)
	}
	output, _ := json.MarshalIndent(sc, "", "\t")
	log.Println("Create Sandbox Transactions:", string(output))
}
