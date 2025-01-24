package main

import (
	"fmt"
	"log"
	"os"
	"time"
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

// waitForFile waits for a file to exist and have content
func waitForFile(filePath string, timeout time.Duration) error {
	vLog("Waiting for file %s to exist and have content...", filePath)
	startTime := time.Now()
	for {
		fileInfo, err := os.Stat(filePath)
		if err == nil && fileInfo.Size() > 0 {
			vLog("File %s exists and has content", filePath)
			return nil
		}
		elapsedTime := time.Since(startTime)
		if elapsedTime > timeout {
			return fmt.Errorf("timeout waiting for file %s", filePath)
		}
		vLog("File %s does not exist or is empty, retrying in 1 second...", filePath)
		time.Sleep(1 * time.Second)
	}
}
