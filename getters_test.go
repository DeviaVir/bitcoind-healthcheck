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
}
