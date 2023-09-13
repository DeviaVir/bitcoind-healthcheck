package main

import (
	"os"
)

// GetEnv with default fallback helper function
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
