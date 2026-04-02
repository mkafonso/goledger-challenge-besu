package usecases_test

import (
	"context"
	"testing"

	memory_providers "github.com/mkafonso/goledger-challenge-besu/__tests__/providers"
	usecase "github.com/mkafonso/goledger-challenge-besu/core/usecases"

	"github.com/stretchr/testify/assert"
)

func TestSetStorageOnBlockchain_ShouldSetValue(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(0)
	uc := usecase.NewSetStorageOnBlockchain(provider)

	request := &usecase.SetStorageOnBlockchainRequest{
		Value: 42,
	}

	response, err := uc.Execute(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)

	value, _ := provider.GetStorageFromBlockchain(context.Background())
	assert.Equal(t, uint64(42), value)
}

func TestSetStorageOnBlockchain_ShouldMapProviderError(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(0)
	provider.ForceError = true

	uc := usecase.NewSetStorageOnBlockchain(provider)

	request := &usecase.SetStorageOnBlockchainRequest{
		Value: 123,
	}

	response, err := uc.Execute(context.Background(), request)

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Equal(t, "unable to write storage to blockchain", err.Error())
}
