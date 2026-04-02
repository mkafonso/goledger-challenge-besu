package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	APIHost                   string
	APIPort                   string
	BlockChainPrivateKey      string
	BlockChainChainID         string
	BlockChainRPCURL          string
	BlockChainContractAddress string
	DatabaseURL               string
}

var Env EnvConfig

func init() {
	_ = godotenv.Load()

	Env = EnvConfig{
		APIHost:                   getEnv("API_HOST", ""),
		APIPort:                   getEnv("API_PORT", ""),
		BlockChainPrivateKey:      getEnv("BLOCKCHAIN_PRIVATE_KEY", ""),
		BlockChainChainID:         getEnv("BLOCKCHAIN_CHAIN_ID", ""),
		BlockChainContractAddress: getEnv("BLOCKCHAIN_CONTRACT_ADDRESS", ""),
		BlockChainRPCURL:          getEnv("BLOCKCHAIN_RPC_URL", ""),
		DatabaseURL:               getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}
