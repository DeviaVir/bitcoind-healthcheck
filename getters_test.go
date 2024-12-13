package main

import (
	"errors"
	"log"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetBlockChainInfo() (*btcjson.GetBlockChainInfoResult, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*btcjson.GetBlockChainInfoResult), args.Error(1)
}

func (m *MockClient) GetIndexInfo() (*btcjson.GetIndexInfoResult, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*btcjson.GetIndexInfoResult), args.Error(1)
}

func (m *MockClient) EstimateSmartFee(blocks int64, mode *btcjson.EstimateSmartFeeMode) (*btcjson.EstimateSmartFeeResult, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*btcjson.EstimateSmartFeeResult), args.Error(1)
}

func TestGetBlockChainInfo(t *testing.T) {
	log.SetOutput(nil)
	mockClient := new(MockClient)

	verificationProgress := float64(0.9)
	mockClient.On("GetBlockChainInfo").Return(&btcjson.GetBlockChainInfoResult{VerificationProgress: verificationProgress}, nil).Once()
	result, err := getBlockChainInfo(mockClient)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, verificationProgress, *result)

	mockClient.On("GetBlockChainInfo").Return(nil, errors.New("Failed to fetch")).Once()
	result, err = getBlockChainInfo(mockClient)
	assert.NotNil(t, err)
	assert.Nil(t, result)

	mockClient.On("GetIndexInfo").Return(&btcjson.GetIndexInfoResult{TxIndex: &btcjson.TxIndex{Synced: true, BestBlockHeight: 0}}, nil).Once()
	result, err = getIndexInfo(mockClient)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, *result)

	mockClient.On("GetIndexInfo").Return(nil, errors.New("Failed to fetch")).Once()
	result, err = getIndexInfo(mockClient)
	assert.NotNil(t, err)
	assert.Nil(t, result)

	fee := 0.04094321
	mockClient.On("EstimateSmartFee").Return(&btcjson.EstimateSmartFeeResult{FeeRate: &fee}, nil).Once()
	result, err = getFeeEstimation(mockClient)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 1.0, *result)

	mockClient.On("EstimateSmartFee").Return(&btcjson.EstimateSmartFeeResult{}, nil).Once()
	result, err = getFeeEstimation(mockClient)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 0.0, *result)

	mockClient.On("EstimateSmartFee").Return(nil, errors.New("Failed to fetch")).Once()
	result, err = getFeeEstimation(mockClient)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
