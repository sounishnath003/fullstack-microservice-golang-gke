package utils

import (
	"log"
	"os"
)

// GetEnv function to get environment variable.
// If the environment variable is not found, it will return the fallback value.
func GetEnv(key string, fallback any) any {
	if val, ok := os.LookupEnv(key); ok {
		log.Printf("key %s environment value found\n", key)
		return val
	}
	log.Printf("no environment value found. setting fallback value key=%s\n", key)
	return fallback
}
