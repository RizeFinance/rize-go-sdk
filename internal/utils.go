package internal

import (
	"fmt"
	"log"
	"os"
)

// Logger wraps Println in a DEBUG env check
func Logger(message string) {
	if os.Getenv("debug") == "true" {
		fmt.Println(message)
	}
}

// CheckEnvVariable is a helper function to check for the existence of an environment variable
func CheckEnvVariable(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Missing '%s' environment variable. Exiting...", key)
	return ""
}
