package bootstrap

import (
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
	"github.com/mkafonso/goledger-challenge-besu/core/usecases"
	dbrepository "github.com/mkafonso/goledger-challenge-besu/infra/db"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"
)

func AppBuildHandlers(
	blockchainProvider ports.StorageBlockchainProvider,
	storageRepository *dbrepository.StorageRepository,
) *httpinfra.Handlers {

	getStorageUC := usecases.NewGetStorageFromBlockchain(blockchainProvider)
	setStorageUC := usecases.NewSetStorageOnBlockchain(blockchainProvider)
	syncStorageUC := usecases.NewSyncStorageToDatabase(blockchainProvider, storageRepository)
	checkConsistencyUC := usecases.NewCheckStorageConsistency(blockchainProvider, storageRepository)

	return &httpinfra.Handlers{
		GetStorageUsecase:       getStorageUC,
		SetStorageUsecase:       setStorageUC,
		SyncStorageUsecase:      syncStorageUC,
		CheckConsistencyUsecase: checkConsistencyUC,
	}
}
