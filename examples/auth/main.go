package main

import (
	"flag"
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
}

// Creates a command-line wrapper for testing the SDK locally
func main() {
	var (
		showHelp bool
		command  string
	)

	flag.BoolVar(&showHelp, "h", false, "Show help menu")
	flag.Parse()

	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		showHelp = true
	}

	help := "--- Rize GO SDK Command-Line Utility --- \n\n" +
		"main.go [COMMAND] \n\n" +
		"COMMAND: SDK command to execute \n\n" +
		"Example: \n" +
		"main.go CreateComplianceWorkflow\n"

	if showHelp {
		fmt.Println(help)
		return
	}

	config := rize.RizeConfig{
		HMACKey:     checkEnvVariable("hmac_key"),
		ProgramUID:  checkEnvVariable("program_uid"),
		Environment: checkEnvVariable("environment"),
	}

	// TODO: Execute sdk command
	fmt.Printf("%+v\n", config)
	fmt.Println("Command:", command)

}

// Helper function to check for environment variables
func checkEnvVariable(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Missing '%s' environment variable. Exiting...", key)
	return ""
}
