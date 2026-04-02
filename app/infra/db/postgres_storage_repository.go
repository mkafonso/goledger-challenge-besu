package repository

import (
	"context"

	db "github.com/mkafonso/goledger-challenge-besu/infra/db/sqlc"
)

type StorageRepository struct {
	q *db.Queries
}

func NewStorageRepository(q *db.Queries) *StorageRepository {
	return &StorageRepository{
		q: q,
	}
}

func (r *StorageRepository) SetStorage(ctx context.Context, value uint64) error {
	_, err := r.q.SetStorage(ctx, int64(value))
	return err
}

func (r *StorageRepository) GetStorage(ctx context.Context) (uint64, error) {
	res, err := r.q.GetStorage(ctx)
	if err != nil {
		return 0, err
	}

	return uint64(res.Value), nil
}
