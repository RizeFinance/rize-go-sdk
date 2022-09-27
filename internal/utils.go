package internal

import (
	"io"
	"log"
	"os"
)

// EnableLogging will customize or disable the standard Go log package
func EnableLogging(enabled bool) {
	if enabled {
		log.SetPrefix("[INFO] ")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
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
