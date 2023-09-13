package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
)

func handleHealthcheck(w http.ResponseWriter, r *http.Request, client *rpcclient.Client, expireSeconds time.Duration, cache *Cache) {
	key := "verificationprogress"
	result, exists := cache.Get(key)
	if !exists {
		verificationProgress := getBlockChainInfo(client)
		if verificationProgress == nil {
			log.Printf("Unable to fetch blockchaininfo")
		}
		result = int(*verificationProgress)
		cache.Set(key, result, expireSeconds)
	}
	resp := map[string]bool{"synced": result == 1}
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	if result < 1 {
		http.Error(w, string(jsonResp), http.StatusServiceUnavailable)
		return
	}
	w.Write(jsonResp)
}