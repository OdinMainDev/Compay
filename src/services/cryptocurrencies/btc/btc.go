package btc

import "github.com/btcsuite/btcd/rpcclient"

var client *rpcclient.Client

func NewService() error {
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
