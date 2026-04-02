package usecase_test

import (
	"context"
	"testing"

	memory_providers "github.com/mkafonso/goledger-challenge-besu/__tests__/providers"
	usecase "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	"github.com/stretchr/testify/assert"
)

func TestGetStorageFromBlockchain_ShouldReturnValue(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(42)
	uc := usecase.NewGetStorageFromBlockchain(provider)

	response, err := uc.Execute(context.Background(), &usecase.GetStorageFromBlockchainRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, uint64(42), response.Value)
}
