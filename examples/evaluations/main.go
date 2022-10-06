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

	// List Evaluations
	elp := rize.EvaluationListParams{
		CustomerUID: "uKxmLxUEiSj5h4M3",
		Latest:      true,
	}
	el, err := rc.Evaluations.List(context.Background(), &elp)
	if err != nil {
		log.Fatal("Error fetching Evaluations\n", err)
	}
	output, _ := json.MarshalIndent(el, "", "\t")
	log.Println("List Evaluations:", string(output))

	// Get Evaluation
	eg, err := rc.Evaluations.Get(context.Background(), "EhrQZJNjCd79LLYq")
	if err != nil {
		log.Fatal("Error fetching Evaluation\n", err)
	}
	output, _ = json.MarshalIndent(eg, "", "\t")
	log.Println("Get Evaluation:", string(output))
}
