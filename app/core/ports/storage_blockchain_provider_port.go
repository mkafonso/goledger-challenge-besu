package ports

import "context"

type StorageBlockchainProvider interface {
	GetStorageFromBlockchain(ctx context.Context) (uint64, error)
	SetStorageOnBlockchain(ctx context.Context, value uint64) error
}
