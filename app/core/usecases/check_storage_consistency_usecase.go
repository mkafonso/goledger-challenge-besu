package usecases

import (
	"context"

	"github.com/mkafonso/goledger-challenge-besu/core/errors"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
)

type CheckStorageConsistencyRequest struct{}

type CheckStorageConsistencyResponse struct {
	Consistent bool `json:"consistent"`
}

type CheckStorageConsistency struct {
	blockchain ports.StorageBlockchainProvider
	repository ports.StorageRepository
}

func NewCheckStorageConsistency(
	blockchain ports.StorageBlockchainProvider,
	repository ports.StorageRepository,
) *CheckStorageConsistency {
	return &CheckStorageConsistency{
		blockchain: blockchain,
		repository: repository,
	}
}

func (uc *CheckStorageConsistency) Execute(ctx context.Context, _ *CheckStorageConsistencyRequest) (*CheckStorageConsistencyResponse, error) {
	blockchainValue, err := uc.blockchain.GetStorageFromBlockchain(ctx)
	if err != nil {
		return nil, errors.NewErrorUnableToReadFromBlockchain()
	}

	dbValue, err := uc.repository.GetStorage(ctx)
	if err != nil {
		return nil, errors.NewInternalError()
	}

	return &CheckStorageConsistencyResponse{
		Consistent: blockchainValue == dbValue,
	}, nil
}
