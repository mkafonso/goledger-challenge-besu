package main

import (
	"context"

	"github.com/mkafonso/goledger-challenge-besu/cmd/bootstrap"
)

func main() {
	ctx := context.Background()

	// Bootstrap dependencies
	blockchainProvider := bootstrap.MustInitBlockchain()
	storageRepository := bootstrap.MustInitDatabase(ctx)

	// Bootstrap app
	handlers := bootstrap.AppBuildHandlers(blockchainProvider, storageRepository)

	// HTTP server
	bootstrap.StartServer(handlers)
}
