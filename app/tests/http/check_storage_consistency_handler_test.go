package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	usecases "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"
	memory_providers "github.com/mkafonso/goledger-challenge-besu/tests/providers"
	memory_repositories "github.com/mkafonso/goledger-challenge-besu/tests/repositories"

	"github.com/stretchr/testify/assert"
)

func TestCheckConsistencyHandler_ShouldReturnConsistent(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(100)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(100)

	checkUC := usecases.NewCheckStorageConsistency(blockchain, repository)
	handlers := &httpinfra.Handlers{
		CheckConsistencyUsecase: checkUC,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/storage/consistency", nil)
	w := httptest.NewRecorder()

	handlers.CheckConsistency(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]bool
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, true, body["consistent"])
}

func TestCheckConsistencyHandler_ShouldReturnInconsistent(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(100)
	repository := memory_repositories.NewMemoryStorageRepositoryProvider(50)

	checkUC := usecases.NewCheckStorageConsistency(blockchain, repository)
	handlers := &httpinfra.Handlers{
		CheckConsistencyUsecase: checkUC,
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/storage/consistency", nil)
	w := httptest.NewRecorder()

	handlers.CheckConsistency(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]bool
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.Equal(t, false, body["consistent"])
}
