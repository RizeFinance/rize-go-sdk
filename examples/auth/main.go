package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	rize "github.com/rizefinance/rize-go-sdk/platform"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	v := rize.Version()
	fmt.Printf("main.init::Loading Rize SDK version: %s\n", v)
}

func main() {
	config := rize.RizeConfig{
		ProgramUID:  checkEnvVariable("program_uid"),
		HMACKey:     checkEnvVariable("hmac_key"),
		Environment: checkEnvVariable("environment"),
		Debug:       true,
	}

	// Enable debug logging
	os.Setenv("debug", fmt.Sprintf("%t", config.Debug))

	// Create new Rize client
	rc := rize.NewRizeClient(&config)

	fmt.Printf("main::RizeClient  %+v\n", *rc)
}

// Helper function to check for environment variables
func checkEnvVariable(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Missing '%s' environment variable. Exiting...", key)
	return ""
}
