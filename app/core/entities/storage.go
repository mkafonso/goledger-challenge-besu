package entities

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	ID        uuid.UUID
	Value     uint64
	UpdatedAt time.Time
}

func NewStorage(value uint64) (*Storage, error) {
	return &Storage{
		ID:        uuid.New(),
		Value:     value,
		UpdatedAt: time.Now().UTC(),
	}, nil
}
