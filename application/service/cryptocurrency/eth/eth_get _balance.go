package eth

import (
	"compay/application/helpers"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func GetETHBalance(address string) (float64, error) {
	addr := common.HexToAddress(address)

	balance, err := ethereumClient.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return 0, err
	}

	return helpers.AmountToFloat(balance.Int64(), 18), nil
}

func GetERC20Balance(address string, contractData struct {
	Address string
	Units   int
}) (float64, error) {
	tokenContract := common.HexToAddress(contractData.Address)

	abiData := `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		return 0, err
	}

	addr := common.HexToAddress(address)
	data, _ := parsedABI.Pack("balanceOf", addr)

	result, err := ethereumClient.CallContract(context.Background(), ethereum.CallMsg{
		To:   &tokenContract,
		Data: data,
	}, nil)
	if err != nil {
		return 0, err
	}

	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return 0, err
	}

	return helpers.AmountToFloat(balance.Int64(), contractData.Units), nil
}
