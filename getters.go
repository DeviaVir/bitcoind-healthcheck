package main

import (
	"github.com/btcsuite/btcd/btcjson"
)

type BlockChainInfoGetter interface {
	GetBlockChainInfo() (*btcjson.GetBlockChainInfoResult, error)
}

func getBlockChainInfo(client BlockChainInfoGetter) *float64 {
	info, err := client.GetBlockChainInfo()
	if err != nil || info == nil {
		return nil
	}

	return &info.VerificationProgress
}
