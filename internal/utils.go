package internal

import (
	"fmt"
	"os"
)

// Logger wraps Println in a DEBUG env check
func Logger(message string) {
	if os.Getenv("debug") == "true" {
		fmt.Println(message)
	}
}
