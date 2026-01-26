package constants

import "compay/common/types"

const (
	CurrencyNetworkBtc types.CurrencyNetwork = "bitcoin"
	CurrencyNetworkEth types.CurrencyNetwork = "ethereum"
	CurrencyNetworkSol types.CurrencyNetwork = "solana"
	CurrencyNetworkTrx types.CurrencyNetwork = "tron"
)

var CurrencyNetworks = []types.CurrencyNetwork{
	CurrencyNetworkBtc,
	CurrencyNetworkEth,
	CurrencyNetworkSol,
	CurrencyNetworkTrx,
}
