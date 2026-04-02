package http_test

import (
	"context"
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

func TestSyncStorageHandler_ShouldSyncSuccessfully(t *testing.T) {
	blockchain := memory_providers.NewMemoryStorageBlockchainProvider(42)
	repo := memory_repositories.NewMemoryStorageRepositoryProvider(0)
	uc := usecases.NewSyncStorageToDatabase(blockchain, repo)
	handlers := &httpinfra.Handlers{
		SyncStorageUsecase: uc,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/storage/sync", nil)
	w := httptest.NewRecorder()

	handlers.SyncStorage(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]bool
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.True(t, body["success"])

	dbVal, _ := repo.GetStorage(context.Background())
	assert.Equal(t, uint64(42), dbVal)
}
