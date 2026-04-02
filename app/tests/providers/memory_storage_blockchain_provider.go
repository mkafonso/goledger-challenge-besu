package memory_providers

import (
	"context"
	"errors"
	"sync"
)

type MemoryStorageBlockchainProvider struct {
	sync.Mutex
	value      uint64
	ForceError bool
}

func NewMemoryStorageBlockchainProvider(initialValue uint64) *MemoryStorageBlockchainProvider {
	return &MemoryStorageBlockchainProvider{
		value: initialValue,
	}
}

func (m *MemoryStorageBlockchainProvider) GetStorageFromBlockchain(ctx context.Context) (uint64, error) {
	m.Lock()
	defer m.Unlock()

	if m.ForceError {
		return 0, errors.New("read error")
	}

	return m.value, nil
}

func (m *MemoryStorageBlockchainProvider) SetStorageOnBlockchain(ctx context.Context, value uint64) error {
	m.Lock()
	defer m.Unlock()

	if m.ForceError {
		return errors.New("write error")
	}

	m.value = value
	return nil
}
