package usecases

import (
	"context"

	"github.com/mkafonso/goledger-challenge-besu/core/errors"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
)

type SyncStorageToDatabaseRequest struct{}

type SyncStorageToDatabaseResponse struct {
	Success bool `json:"success"`
}

type SyncStorageToDatabase struct {
	blockchain ports.StorageBlockchainProvider
	repository ports.StorageRepository
}

func NewSyncStorageToDatabase(
	blockchain ports.StorageBlockchainProvider,
	repository ports.StorageRepository,
) *SyncStorageToDatabase {
	return &SyncStorageToDatabase{
		blockchain: blockchain,
		repository: repository,
	}
}

func (uc *SyncStorageToDatabase) Execute(ctx context.Context, _ *SyncStorageToDatabaseRequest) (*SyncStorageToDatabaseResponse, error) {
	value, err := uc.blockchain.GetStorageFromBlockchain(ctx)
	if err != nil {
		return nil, errors.NewErrorUnableToReadFromBlockchain()
	}

	if err := uc.repository.SetStorage(ctx, value); err != nil {
		return nil, errors.NewInternalError()
	}

	return &SyncStorageToDatabaseResponse{
		Success: true,
	}, nil
}
