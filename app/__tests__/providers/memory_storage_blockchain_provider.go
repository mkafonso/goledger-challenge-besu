package memory_providers

import (
	"context"
	"sync"
)

type MemoryStorageBlockchainProvider struct {
	sync.Mutex
	value uint64
}

func NewMemoryStorageBlockchainProvider(initialValue uint64) *MemoryStorageBlockchainProvider {
	return &MemoryStorageBlockchainProvider{
		value: initialValue,
	}
}

func (m *MemoryStorageBlockchainProvider) GetStorageFromBlockchain(ctx context.Context) (uint64, error) {
	m.Lock()
	defer m.Unlock()

	return m.value, nil
}
