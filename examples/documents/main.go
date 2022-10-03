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

	// List Documents
	dlp := rize.DocumentListParams{
		DocumentType:        "monthly_statement",
		Month:               1,
		Year:                2020,
		CustodialAccountUID: "yqyYk5b1xgXFFrXs",
		CustomerUID:         "uKxmLxUEiSj5h4M3",
		SyntheticAccountUID: "4XkJnsfHsuqrxmeX",
		Limit:               100,
		Offset:              0,
	}
	dl, err := rc.Documents.List(context.Background(), &dlp)
	if err != nil {
		log.Fatal("Error fetching Documents\n", err)
	}
	output, _ := json.MarshalIndent(dl, "", "\t")
	log.Println("List Documents:", string(output))

	// Get Document
	dg, err := rc.Documents.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Document\n", err)
	}
	output, _ = json.MarshalIndent(dg, "", "\t")
	log.Println("Get Document:", string(output))

	// View Document
	dv, err := rc.Documents.View(context.Background(), "u8EHFJnWvJxRJZxa")
	if err != nil {
		log.Fatal("Error viewing document\n", err)
	}
	output, _ = json.MarshalIndent(dv, "", "\t")
	log.Println("View Document:", string(output))
}
