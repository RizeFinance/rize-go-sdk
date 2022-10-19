package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/examples"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	v := rize.Version()
	log.Printf("Loading Rize SDK version: %s\n", v)
}

func main() {
	// Create new Rize client for examples
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	examples.ExampleAuthService_GetToken(rc)

}
