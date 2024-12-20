package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
)

func caller(key string, client *rpcclient.Client, cache *Cache, expiration time.Duration, fn func(BlockChainInfoGetter) (*float64, error)) float64 {
	result, exists := cache.Get(key)
	if !exists {
		resp, err := fn(client)
		if resp == nil || err != nil {
			log.Printf("Unable to fetch %s (nil or err): %s", key, err)
			result = float64(0)
		} else {
			result = *resp
			cache.Set(key, result, expiration)
		}
	}
	return result
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request, client *rpcclient.Client, expiration time.Duration, waitForTxIndex bool, waitForFeeEstimation bool, cache *Cache) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]bool)
	allTrue := true

	vLog("handler.go: handling healthcheck")
	vLog("handler.go: handling txindex")
	if waitForTxIndex {
		indexRes := caller("indexinfo", client, cache, expiration, getIndexInfo)
		resp["gettxindexinfo"] = indexRes > 0.0
		if !resp["gettxindexinfo"] {
			allTrue = false
		}
	}

	vLog("handler.go: handling fee estimation")
	if waitForFeeEstimation {
		feeRes := caller("estimatefee", client, cache, expiration, getFeeEstimation)
		resp["estimatesmartfee"] = feeRes > 0.0
		if !resp["estimatesmartfee"] {
			allTrue = false
		}
	}

	vLog("handler.go: handling blockchain info")
	progressRes := caller("verificationprogress", client, cache, expiration, getBlockChainInfo)
	resp["verificationprogress"] = progressRes > 0.9999
	if !resp["verificationprogress"] {
		allTrue = false
	}

	vLog("handler.go: sending response")
	jsonResp, _ := json.Marshal(resp)
	if !allTrue {
		vLog("handler.go: not all checks passed")
		http.Error(w, string(jsonResp), http.StatusServiceUnavailable)
		return
	}

	w.Write(jsonResp)
}
