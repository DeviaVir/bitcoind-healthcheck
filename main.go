package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
)

func main() {
	cache := NewCache()

	connCfg := &rpcclient.ConnConfig{
		Host:         GetEnv("RPC_HOST", "localhost:8332"),
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	if GetEnv("RPC_USER", "") != "" {
		connCfg.User = GetEnv("RPC_USER", "")
		connCfg.Pass = GetEnv("RPC_PASS", "")
	} else {
		connCfg.CookiePath = GetEnv("RPC_COOKIE_PATH", "")
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
	}
	defer client.Shutdown()

	cacheExpireSeconds, err := strconv.Atoi(GetEnv("CACHE_EXPIRE_SECONDS", "14"))
	if err != nil {
		log.Fatalf("Could not convert %s to integer: %s", GetEnv("CACHE_EXPIRE_SECONDS", "14"), err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleHealthcheck(w, r, client, time.Duration(time.Duration(cacheExpireSeconds)*time.Second), cache)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", GetEnv("PORT", "8080")), nil))
}
