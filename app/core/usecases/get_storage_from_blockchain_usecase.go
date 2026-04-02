package usecase

import (
	"context"

	"github.com/mkafonso/goledger-challenge-besu/core/errors"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
)

type GetStorageFromBlockchainRequest struct{}

type GetStorageFromBlockchainResponse struct {
	Value uint64
}

type GetStorageFromBlockchain struct {
	blockchain ports.StorageBlockchainProvider
}

func NewGetStorageFromBlockchain(blockchain ports.StorageBlockchainProvider) *GetStorageFromBlockchain {
	return &GetStorageFromBlockchain{
		blockchain: blockchain,
	}
}

func (uc *GetStorageFromBlockchain) Execute(ctx context.Context, _ *GetStorageFromBlockchainRequest) (*GetStorageFromBlockchainResponse, error) {
	value, err := uc.blockchain.GetStorageFromBlockchain(ctx)
	if err != nil {
		return nil, errors.NewErrorUnableToReadFromBlockchain()
	}

	return &GetStorageFromBlockchainResponse{
		Value: value,
	}, nil
}
