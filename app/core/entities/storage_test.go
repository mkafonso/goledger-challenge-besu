package entities_test

import (
	"testing"
	"time"

	"github.com/mkafonso/goledger-challenge-besu/core/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	value := uint64(42)

	storage, err := entities.NewStorage(value)

	assert.NoError(t, err)
	assert.NotNil(t, storage)

	assert.Equal(t, value, storage.Value)
	assert.NotZero(t, storage.ID)
	assert.WithinDuration(t, time.Now().UTC(), storage.UpdatedAt, time.Second)
}
