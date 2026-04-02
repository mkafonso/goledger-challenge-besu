package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	usecases "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"
	memory_providers "github.com/mkafonso/goledger-challenge-besu/tests/providers"

	"github.com/stretchr/testify/assert"
)

func TestGetStorageHandler_ShouldReturnValue(t *testing.T) {
	provider := memory_providers.NewMemoryStorageBlockchainProvider(42)
	uc := usecases.NewGetStorageFromBlockchain(provider)
	handlers := &httpinfra.Handlers{
		GetStorageUsecase: uc,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/storage", nil)
	w := httptest.NewRecorder()

	handlers.GetStorage(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]uint64
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, uint64(42), body["value"])
}
