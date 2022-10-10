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

// JSONKeys will extract key values from a JSON object
func JSONKeys(data map[string]interface{}) []string {
	keys := []string{}

	for key, value := range data {
		// Check for json object
		if val, ok := value.(map[string]interface{}); ok {
			keys = append(keys, key)
			for _, k := range JSONKeys(val) {
				keys = append(keys, k)
			}
			// Check for json array
		} else if val, ok := value.([]interface{}); ok {
			keys = append(keys, key)
			for _, v := range val {
				// Check for json object
				if subObject, ok := v.(map[string]interface{}); ok {
					for _, subObjectKey := range JSONKeys(subObject) {
						keys = append(keys, subObjectKey)
					}
				}
			}
		} else {
			keys = append(keys, key)
		}
	}
	return keys
}

// Difference will find any string from a[] that are not present in b[]
func Difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
