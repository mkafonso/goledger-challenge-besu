package usecases_test

import (
	"context"
	"testing"

	usecase "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	memory_providers "github.com/mkafonso/goledger-challenge-besu/tests/providers"
	memory_repositories "github.com/mkafonso/goledger-challenge-besu/tests/repositories"

	"github.com/stretchr/testify/assert"
)

func TestCheckStorageConsistency_ShouldBeConsistent(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(100)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(100)

	uc := usecase.NewCheckStorageConsistency(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.CheckStorageConsistencyRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Consistent)
}

func TestCheckStorageConsistency_ShouldBeInconsistent(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(100)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(50)

	uc := usecase.NewCheckStorageConsistency(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.CheckStorageConsistencyRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.False(t, response.Consistent)
}

func TestCheckStorageConsistency_ShouldMapBlockchainError(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(0)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(0)
	blockchain.ForceError = true

	uc := usecase.NewCheckStorageConsistency(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.CheckStorageConsistencyRequest{})

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Equal(t, "unable to read storage from blockchain", err.Error())
}

func TestCheckStorageConsistency_ShouldMapRepositoryError(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(100)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(0)
	repository.ForceError = true

	uc := usecase.NewCheckStorageConsistency(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.CheckStorageConsistencyRequest{})

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Equal(t, "internal server error", err.Error())
}
