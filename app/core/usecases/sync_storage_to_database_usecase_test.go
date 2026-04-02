package usecase_test

import (
	"context"
	"testing"

	memory_providers "github.com/mkafonso/goledger-challenge-besu/__tests__/providers"
	memory_repositories "github.com/mkafonso/goledger-challenge-besu/__tests__/repositories"
	usecase "github.com/mkafonso/goledger-challenge-besu/core/usecases"

	"github.com/stretchr/testify/assert"
)

func TestSyncStorageToDatabase_ShouldSyncSuccessfully(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(42)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(0)

	uc := usecase.NewSyncStorageToDatabase(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.SyncStorageToDatabaseRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)
}

func TestSyncStorageToDatabase_ShouldMapBlockchainError(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(0)
	blockchain.ForceError = true

	repository := memory_repositories.NewMemoryStorageRepositoryProvider(0)
	uc := usecase.NewSyncStorageToDatabase(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.SyncStorageToDatabaseRequest{})

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Equal(t, "unable to read storage from blockchain", err.Error())
}

func TestSyncStorageToDatabase_ShouldMapRepositoryError(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(42)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(0)
	repository.ForceError = true

	uc := usecase.NewSyncStorageToDatabase(blockchain, repository)

	response, err := uc.Execute(context.Background(), &usecase.SyncStorageToDatabaseRequest{})

	assert.Nil(t, response)
	assert.Error(t, err)
	assert.Equal(t, "internal server error", err.Error())
}
