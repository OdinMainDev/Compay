package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func GetTransactions(address string) ([]types.Log, error) {
	addr := common.HexToAddress(address)

	filter := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	logs, err := ethereumClient.FilterLogs(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	txHashObj := common.HexToHash(txHash)

	tx, isPending, err := ethereumClient.TransactionByHash(context.Background(), txHashObj)
	if err != nil {
		return nil, false, err
	}

	return tx, isPending, nil
}
