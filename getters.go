package main

import (
	"github.com/btcsuite/btcd/btcjson"
)

type BlockChainInfoGetter interface {
	GetBlockChainInfo() (*btcjson.GetBlockChainInfoResult, error)
	GetIndexInfo() (*btcjson.GetIndexInfoResult, error)
	EstimateSmartFee(int64, *btcjson.EstimateSmartFeeMode) (*btcjson.EstimateSmartFeeResult, error)
}

func getBlockChainInfo(client BlockChainInfoGetter) (*float64, error) {
	vLog("getters.go: Getting blockchain info")
	info, err := client.GetBlockChainInfo()
	if err != nil || info == nil {
		vLog("getters.go: Error getting blockchain info: %s", err)
		return nil, err
	}

	vLog("getters.go: Blockchain info: %f", info.VerificationProgress)
	return &info.VerificationProgress, nil
}

func getIndexInfo(client BlockChainInfoGetter) (*float64, error) {
	vLog("getters.go: Getting index info")
	info, err := client.GetIndexInfo()
	if err != nil || info == nil {
		vLog("getters.go: Error getting index info: %s", err)
		return nil, err
	}

	var result float64
	result = 0.0
	if info.Synced {
		vLog("getters.go: Index is synced")
		result = 1.0
	}
	return &result, nil
}

func getFeeEstimation(client BlockChainInfoGetter) (*float64, error) {
	vLog("getters.go: Getting fee estimation")
	info, err := client.EstimateSmartFee(1, nil)
	if err != nil || info == nil {
		vLog("getters.go: Error getting fee estimation: %s", err)
		return nil, err
	}

	var result float64
	result = 0.0
	if info.FeeRate != nil {
		vLog("getters.go: Fee estimation: %d", info.FeeRate)
		result = 1.0
	}
	return &result, nil
}
