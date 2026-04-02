package usecases

import (
	"context"

	"github.com/mkafonso/goledger-challenge-besu/core/errors"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
)

type SetStorageOnBlockchainRequest struct {
	Value uint64
}

type SetStorageOnBlockchainResponse struct {
	Success bool
}

type SetStorageOnBlockchain struct {
	blockchain ports.StorageBlockchainProvider
}

func NewSetStorageOnBlockchain(blockchain ports.StorageBlockchainProvider) *SetStorageOnBlockchain {
	return &SetStorageOnBlockchain{
		blockchain: blockchain,
	}
}

func (uc *SetStorageOnBlockchain) Execute(ctx context.Context, data *SetStorageOnBlockchainRequest) (*SetStorageOnBlockchainResponse, error) {
	if err := uc.blockchain.SetStorageOnBlockchain(ctx, data.Value); err != nil {
		return nil, errors.NewErrorUnableToWriteToBlockchain()
	}

	return &SetStorageOnBlockchainResponse{
		Success: true,
	}, nil
}
