package constants

import (
	"compay/common/types"
	"efw/components/entity"
)

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
	BtcApiUrl                 string                   = "http://localhost:8081"
	BtcNodeUrl                string                   = "https://bitcoin-testnet.core.chainstack.com/ff918ee18e1e5abbe9267aa44be1d08d"
	BtcCodeAlpha              entity.CurrencyCodeAlpha = "BTC"
	BtcMinAmount              entity.Amount            = 10000 // 0.0001
	BtcMaxAmount              entity.Amount            = 0
	BtcMinAmountWithdraw      entity.Amount            = 10000     // 0.0001
	BtcMaxAmountWithdraw      entity.Amount            = 500000000 // 5 BTC
	BtcConfirmation           types.Confirmation       = 3
	BtcInvoiceTime            int64                    = 86400 // Время инвойса в секундах
	BtcEstimatedFeeBlockCount int64                    = 6
	BtcFeeType                types.FeeType            = FeeTypeOnlyWithdraw
	BtcFeeFixed               types.FeeAmount          = 10000 // 0.0001
	BtcFeePercent             types.FeePercentAmount   = 25    // 0.25
)
