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

	// List Customer Products
	cpp := rize.CustomerProductListParams{
		ProgramUID:  "pQtTCSXz57fuefzp",
		ProductUID:  "zbJbEa72eKMgbbBv",
		CustomerUID: "uKxmLxUEiSj5h4M3",
	}
	pl, err := rc.CustomerProducts.List(context.Background(), &cpp)
	if err != nil {
		log.Fatal("Error fetching Customer Products\n", err)
	}
	output, _ := json.Marshal(pl)
	log.Println("List Customer Products:", string(output))

	// Create Customer Product
	ccp := rize.CustomerProductCreateParams{
		CustomerUID: "S62MaHx6WwsqG9vQ",
		ProductUID:  "pQtTCSXz57fuefzp",
	}
	cc, err := rc.CustomerProducts.Create(context.Background(), &ccp)
	if err != nil {
		log.Fatal("Error creating Customer Product\n", err)
	}
	output, _ = json.Marshal(cc)
	log.Println("Create Customer Product:", string(output))

	// Get Customer Product
	cp, err := rc.CustomerProducts.Get(context.Background(), "Tegvs2E4TQgVYYMj")
	if err != nil {
		log.Fatal("Error fetching Customer Product\n", err)
	}
	output, _ = json.Marshal(cp)
	log.Println("Get Customer Product:", string(output))
}
