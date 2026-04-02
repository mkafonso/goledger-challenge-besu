package bootstrap

import (
	"log"

	"github.com/mkafonso/goledger-challenge-besu/config"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
	"github.com/mkafonso/goledger-challenge-besu/infra/blockchain"
)

func MustInitBlockchain() ports.StorageBlockchainProvider {
	provider, err := blockchain.NewBesuStorageProvider(
		config.Env.BlockChainRPCURL,
		config.Env.BlockChainContractAddress,
	)
	if err != nil {
		log.Fatal("failed to init blockchain provider:", err)
	}

	return provider
}
