package memory_repositories

import (
	"context"
	"errors"
	"sync"
)

type MemoryStorageRepositoryProvider struct {
	sync.Mutex
	value      uint64
	ForceError bool
}

func NewMemoryStorageRepositoryProvider(initialValue uint64) *MemoryStorageRepositoryProvider {
	return &MemoryStorageRepositoryProvider{
		value: initialValue,
	}
}

func (m *MemoryStorageRepositoryProvider) SetStorage(ctx context.Context, value uint64) error {
	m.Lock()
	defer m.Unlock()

	if m.ForceError {
		return errors.New("write error")
	}

	m.value = value
	return nil
}

func (m *MemoryStorageRepositoryProvider) GetStorage(ctx context.Context) (uint64, error) {
	m.Lock()
	defer m.Unlock()

	if m.ForceError {
		return 0, errors.New("read error")
	}

	return m.value, nil
}
