package entities

import (
	"time"
)

type Storage struct {
	Value     uint64
	UpdatedAt time.Time
}

func NewStorage(value uint64) (*Storage, error) {
	return &Storage{
		Value:     value,
		UpdatedAt: time.Now().UTC(),
	}, nil
}
