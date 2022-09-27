package internal

import (
	"io"
	"log"
	"os"
)

// EnableLogging disables the Go log package when enabled is false
func EnableLogging(enabled bool) {
	if !enabled {
		log.SetFlags(0)
		log.SetOutput(io.Discard)
	} else {
		log.SetPrefix("[INFO] ")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
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
