package ports

import "context"

type StorageRepository interface {
	SetStorage(ctx context.Context, value uint64) error
	GetStorage(ctx context.Context) (uint64, error)
}
