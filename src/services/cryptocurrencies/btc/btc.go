package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

var client *rpcclient.Client
var networkConfig *chaincfg.Params = &chaincfg.TestNet3Params

func Start() error {
	c, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         "bitcoin-testnet.core.chainstack.com:443/ff918ee18e1e5abbe9267aa44be1d08d",
		User:         "goofy-mclean",
		Pass:         "tartar-barbed-rented-gloss-hazing-cross",
		HTTPPostMode: true,
		DisableTLS:   false,
	}, nil)
	if err != nil {
		return err
	}

	client = c
	return nil
}

// GetNetworkConfig возвращает конфигурацию сети
func GetNetworkConfig() *chaincfg.Params {
	return networkConfig
}
