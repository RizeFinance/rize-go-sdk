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

	// List Products
	pl, err := rc.Products.List(context.Background(), "pQtTCSXz57fuefzp")
	if err != nil {
		log.Fatal("Error fetching products\n", err)
	}
	output, _ := json.MarshalIndent(pl, "", "\t")
	log.Println("List Products:", string(output))

	// Get Product
	pg, err := rc.Products.Get(context.Background(), "f9VncZny4ejhcPF4")
	if err != nil {
		log.Fatal("Error fetching product\n", err)
	}
	output, _ = json.MarshalIndent(pg, "", "\t")
	log.Println("Get Product:", string(output))
}
