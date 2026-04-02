package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mkafonso/goledger-challenge-besu/config"
	"github.com/mkafonso/goledger-challenge-besu/core/ports"
)

type BesuStorageProvider struct {
	client   *ethclient.Client
	contract common.Address
}

func NewBesuStorageProvider(rpcURL string, contractAddress string) (ports.StorageBlockchainProvider, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to besu: %w", err)
	}

	return &BesuStorageProvider{
		client:   client,
		contract: common.HexToAddress(contractAddress),
	}, nil
}

func (b *BesuStorageProvider) GetStorageFromBlockchain(ctx context.Context) (uint64, error) {
	data, err := b.client.StorageAt(ctx, b.contract, common.Hash{}, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to read storage: %w", err)
	}

	value := new(big.Int).SetBytes(data)
	return value.Uint64(), nil
}

func (b *BesuStorageProvider) SetStorageOnBlockchain(ctx context.Context, value uint64) error {
	privateKeyHex := strings.TrimPrefix(config.Env.BlockChainPrivateKey, "0x")
	if privateKeyHex == "" {
		return fmt.Errorf("missing env BLOCKCHAIN_PRIVATE_KEY")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("invalid private key: %w", err)
	}

	chainID, ok := new(big.Int).SetString(config.Env.BlockChainChainID, 10)
	if !ok {
		return fmt.Errorf("invalid env BLOCKCHAIN_CHAIN_ID: %q", config.Env.BlockChainChainID)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"set","outputs":[],"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("failed to parse ABI: %w", err)
	}

	data, err := parsedABI.Pack("set", new(big.Int).SetUint64(value))
	if err != nil {
		return fmt.Errorf("failed to pack tx data: %w", err)
	}

	nonce, err := b.client.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	gasPrice, err := b.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	tx := types.NewTransaction(
		nonce,
		b.contract,
		big.NewInt(0),
		300000,
		gasPrice,
		data,
	)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return fmt.Errorf("failed to sign tx: %w", err)
	}

	err = b.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}

	return nil
}
