package main

import (
	"flag"
	"fmt"

	rize "github.com/rizefinance/rize-go-sdk/v1"
)

var config rize.Config

// Creates a command-line wrapper for testing the SDK locally
func main() {
	var (
		hmac        string
		programUID  string
		environment string
		showHelp    bool
	)

	flag.BoolVar(&showHelp, "h", false, "Show help menu")
	flag.StringVar(&hmac, "k", "", "Rize HMAC key")
	flag.StringVar(&programUID, "p", "", "Program UID")
	flag.StringVar(&environment, "e", "", "Environment Tier")

	flag.Parse()

	help := "--- Rize GO SDK Command-Line Utility --- \n\n" +
		"rize [-k HMAC_KEY] [-p PROGRAM_UID] [-e ENVIRONMENT] \n" +
		"HMAC_KEY: HMAC key for the target environment \n" +
		"PROGRAM_UID: Program UID for the target environment \n" +
		"ENVIRONMENT: The Rize environment to be used: 'sandbox', 'integration' or 'production' \n\n" +
		"Example: \n" +
		"rize -k abcdefg -p 12345 -e sandbox"

	if showHelp || hmac == "" || programUID == "" || environment == "" {
		fmt.Println(help)
		return
	}

	config = rize.Config{
		HMACKey:     hmac,
		ProgramUID:  programUID,
		Environment: environment,
	}

	fmt.Printf("%+v", config)

}
