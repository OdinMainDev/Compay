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

const (
	CurrencyCodeAlphaBtc      types.CurrencyCodeAlpha = "BTC"
	CurrencyMinAmountBtc      types.Amount            = 0.0001
	CurrencyMaxAmountBtc      types.Amount            = 5
	CurrencyConfirmationBtc   types.Confirmation      = 3
	CurrencyFeeTypeBtc        types.FeeType           = FeeTypeOnlyWithdraw
	CurrencyFeePercentTypeBtc types.FeeType           = FeeTypeOnlyWithdraw
	CurrencyFeeBtc            types.FeeAmount         = 0.0001
	CurrencyFeePercentBtc     types.FeePercentAmount  = 0.25
)
