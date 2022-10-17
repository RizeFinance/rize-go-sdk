package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize client
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// List Pinwheel Jobs
	plp := rize.PinwheelJobListParams{
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              0,
	}
	pl, err := rc.PinwheelJobs.List(context.Background(), &plp)
	if err != nil {
		log.Fatal("Error fetching Pinwheel Jobs\n", err)
	}
	output, _ := json.MarshalIndent(pl, "", "\t")
	log.Println("List Pinwheel Jobs:", string(output))

	// Create Pinwheel Job
	pcp := rize.PinwheelJobCreateParams{
		JobNames:             []string{"direct_deposit_switch"},
		SyntheticAccountUID:  "4XkJnsfHsuqrxmeX",
		Amount:               1000,
		DisablePartialSwitch: false,
		OrganizationName:     "Chipotle Mexican Grill, Inc.",
		SkipWelcomeScreen:    false,
	}
	pc, err := rc.PinwheelJobs.Create(context.Background(), &pcp)
	if err != nil {
		log.Fatal("Error creating Pinwheel Job\n", err)
	}
	output, _ = json.MarshalIndent(pc, "", "\t")
	log.Println("Create Pinwheel Job:", string(output))

	// Get PinwheelJob
	pg, err := rc.PinwheelJobs.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Pinwheel Job\n", err)
	}
	output, _ = json.MarshalIndent(pg, "", "\t")
	log.Println("Get Pinwheel Job:", string(output))
}
