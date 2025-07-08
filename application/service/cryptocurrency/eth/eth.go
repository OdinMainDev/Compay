package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ethereumClient *ethclient.Client

func Init() {
	eC, err := ethclient.Dial("https://ethereum.publicnode.com")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the Ethereum client: %v\n", err))
	}
	ethereumClient = eC
}

var contractAddresses = map[string]string{}
