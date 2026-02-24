package btc

import (
	"compay/common/config"
	"sync"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
)

var client *rpcclient.Client
var clientMutex sync.RWMutex
var networkConfig *chaincfg.Params = &chaincfg.TestNet3Params

func Start() error {
	if config.BtcNetworkMainnet {
		networkConfig = &chaincfg.MainNetParams
	}

	c, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         config.BtcNodeHost,
		User:         config.BtcNodeUsername,
		Pass:         config.BtcNodePassword,
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
