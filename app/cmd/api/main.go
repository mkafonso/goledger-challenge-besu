package main

import (
	"fmt"
	"log"
	"net/http"

	memory_repositories "github.com/mkafonso/goledger-challenge-besu/__tests__/repositories"
	"github.com/mkafonso/goledger-challenge-besu/config"
	usecases "github.com/mkafonso/goledger-challenge-besu/core/usecases"
	blockchain "github.com/mkafonso/goledger-challenge-besu/infra/blockchain"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"
)

func main() {
	blockchainProvider, err := blockchain.NewBesuStorageProvider(
		config.Env.BlockChainRPCURL,
		config.Env.BlockChainContractAddress,
	)
	if err != nil {
		log.Fatal("failed to init blockchain provider:", err)
	}

	storageRepository := memory_repositories.NewMemoryStorageRepositoryProvider(100)

	getStorageUC := usecases.NewGetStorageFromBlockchain(blockchainProvider)
	setStorageUC := usecases.NewSetStorageOnBlockchain(blockchainProvider)
	syncStorageUC := usecases.NewSyncStorageToDatabase(blockchainProvider, storageRepository)
	checkConsistencyUC := usecases.NewCheckStorageConsistency(blockchainProvider, storageRepository)

	handlers := &httpinfra.Handlers{
		GetStorageUsecase:       getStorageUC,
		SetStorageUsecase:       setStorageUC,
		SyncStorageUsecase:      syncStorageUC,
		CheckConsistencyUsecase: checkConsistencyUC,
	}

	router := httpinfra.NewRouter(handlers)

	fmt.Println("HTTP Server running on port: ", config.Env.APIHost+":"+config.Env.APIPort)
	if err := http.ListenAndServe(":"+config.Env.APIPort, router); err != nil {
		log.Fatal(err)
	}
}
