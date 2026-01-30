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

// BTC
const (
	BtcNodeUrl           string                  = "http://localhost:8081"
	BtcCodeAlpha         types.CurrencyCodeAlpha = "BTC"
	BtcMinAmount         types.Amount            = 10000 // 0.0001
	BtcMaxAmount         types.Amount            = 0
	BtcMinAmountWithdraw types.Amount            = 10000     // 0.0001
	BtcMaxAmountWithdraw types.Amount            = 500000000 // 5 BTC
	BtcConfirmation      types.Confirmation      = 3
	BtcFeeType           types.FeeType           = FeeTypeOnlyWithdraw
	BtcFeeFixed          types.FeeAmount         = 10000 // 0.0001
	BtcFeePercent        types.FeePercentAmount  = 25    // 0.25
)
