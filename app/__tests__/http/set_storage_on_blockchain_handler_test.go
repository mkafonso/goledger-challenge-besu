package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	memory_providers "github.com/mkafonso/goledger-challenge-besu/__tests__/providers"
	apperrors "github.com/mkafonso/goledger-challenge-besu/core/errors"
	usecases "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"

	"github.com/stretchr/testify/assert"
)

func TestSetStorageHandler_ShouldSetValue(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(0)
	uc := usecases.NewSetStorageOnBlockchain(provider)
	handlers := &httpinfra.Handlers{
		SetStorageUsecase: uc,
	}

	bodyReq, _ := json.Marshal(map[string]uint64{
		"value": 42,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/storage", bytes.NewReader(bodyReq))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.SetStorage(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]bool
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.True(t, body["success"])

	val, _ := provider.GetStorageFromBlockchain(context.Background())
	assert.Equal(t, uint64(42), val)
}

func TestSetStorageHandler_ShouldReturnBadRequest_WhenInvalidValueType(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(0)
	uc := usecases.NewSetStorageOnBlockchain(provider)
	handlers := &httpinfra.Handlers{
		SetStorageUsecase: uc,
	}

	bodyReq, _ := json.Marshal(map[string]int64{
		"value": -1,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/storage", bytes.NewReader(bodyReq))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.SetStorage(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var body apperrors.AppError
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, "BAD_REQUEST", body.Code)
	assert.NotEmpty(t, body.Message)
	assert.NotEmpty(t, body.Action)
}
