package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mkafonso/goledger-challenge-besu/core/usecases"
)

type Handlers struct {
	GetStorageUsecase       *usecases.GetStorageFromBlockchain
	SetStorageUsecase       *usecases.SetStorageOnBlockchain
	SyncStorageUsecase      *usecases.SyncStorageToDatabase
	CheckConsistencyUsecase *usecases.CheckStorageConsistency
}

func (h *Handlers) GetStorage(w http.ResponseWriter, r *http.Request) {
	resp, err := h.GetStorageUsecase.Execute(context.Background(), &usecases.GetStorageFromBlockchainRequest{})
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) SetStorage(w http.ResponseWriter, r *http.Request) {
	var req usecases.SetStorageOnBlockchainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.SetStorageUsecase.Execute(context.Background(), &req)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) SyncStorage(w http.ResponseWriter, r *http.Request) {
	resp, err := h.SyncStorageUsecase.Execute(context.Background(), &usecases.SyncStorageToDatabaseRequest{})
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}

func (h *Handlers) CheckConsistency(w http.ResponseWriter, r *http.Request) {
	resp, err := h.CheckConsistencyUsecase.Execute(context.Background(), &usecases.CheckStorageConsistencyRequest{})
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, resp)
}
