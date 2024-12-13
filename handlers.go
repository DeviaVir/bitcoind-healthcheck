package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
)

func caller(key string, client *rpcclient.Client, cache *Cache, expireSeconds time.Duration, fn func(BlockChainInfoGetter) (*float64, error)) float64 {
	result, exists := cache.Get(key)
	if !exists {
		resp, err := fn(client)
		if resp == nil || err != nil {
			log.Printf("Unable to fetch %s (nil or err): %s", key, err)
			result = float64(0)
		} else {
			result = *resp
			cache.Set(key, result, expireSeconds)
		}
	}
	return result
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request, client *rpcclient.Client, expireSeconds time.Duration, waitForTxIndex bool, waitForFeeEstimation bool, cache *Cache) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]bool)
	allTrue := true

	if waitForTxIndex {
		indexRes := caller("indexinfo", client, cache, expireSeconds, getIndexInfo)
		resp["gettxindexinfo"] = indexRes > 0.0
		if !resp["gettxindexinfo"] {
			allTrue = false
		}
	}

	if waitForFeeEstimation {
		feeRes := caller("estimatefee", client, cache, expireSeconds, getFeeEstimation)
		resp["estimatesmartfee"] = feeRes > 0.0
		if !resp["estimatesmartfee"] {
			allTrue = false
		}
	}

	progressRes := caller("verificationprogress", client, cache, expireSeconds, getBlockChainInfo)
	resp["verificationprogress"] = progressRes > 0.0
	if !resp["verificationprogress"] {
		allTrue = false
	}

	jsonResp, _ := json.Marshal(resp)
	if !allTrue {
		http.Error(w, string(jsonResp), http.StatusServiceUnavailable)
		return
	}

	w.Write(jsonResp)
}
