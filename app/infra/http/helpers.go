package http

import (
	"encoding/json"
	"net/http"

	"github.com/mkafonso/goledger-challenge-besu/core/errors"
)

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if err == nil {
		WriteJSON(w, http.StatusInternalServerError, errors.NewInternalError())
		return
	}

	if appErr, ok := err.(*errors.AppError); ok {
		WriteJSON(w, status, appErr)
		return
	}

	if status >= http.StatusInternalServerError {
		WriteJSON(w, status, errors.NewInternalError())
		return
	}

	WriteJSON(w, status, errors.NewBadRequestError(err.Error()))
}
