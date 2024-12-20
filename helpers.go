package main

import (
	"log"
	"os"
)

// GetEnv with default fallback helper function
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// vLog logs a message if the verbose flag is set
func vLog(format string, a ...interface{}) {
	if verbose {
		log.Printf(format, a...)
	}
}
