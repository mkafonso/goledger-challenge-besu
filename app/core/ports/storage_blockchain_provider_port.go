package ports

import "context"

type StorageBlockchainProvider interface {
	GetStorageFromBlockchain(ctx context.Context) (uint64, error)
}
