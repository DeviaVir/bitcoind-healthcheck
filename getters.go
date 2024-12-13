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
	info, err := client.GetBlockChainInfo()
	if err != nil || info == nil {
		return nil, err
	}

	return &info.VerificationProgress, nil
}

func getIndexInfo(client BlockChainInfoGetter) (*float64, error) {
	info, err := client.GetIndexInfo()
	if err != nil || info == nil {
		return nil, err
	}

	var result float64
	result = 0.0
	if info.Synced {
		result = 1.0
	}
	return &result, nil
}

func getFeeEstimation(client BlockChainInfoGetter) (*float64, error) {
	info, err := client.EstimateSmartFee(1, nil)
	if err != nil || info == nil {
		return nil, err
	}

	var result float64
	result = 0.0
	if info.FeeRate != nil {
		result = 1.0
	}
	return &result, nil
}
