package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestCheckBalance struct {
	Address          string `json:"address" binding:"required"`
	MinConfirmations int    `json:"min_confirmations"`
}

type responseCheckBalance struct {
	Address          string  `json:"address"`
	Balance          float64 `json:"balance"`
	BalanceSatoshi   int64   `json:"balance_satoshi"`
	MinConfirmations int     `json:"min_confirmations"`
}

// CheckBalance проверяет баланс BTC адреса
func CheckBalance(context *gin.Context) {
	var input requestCheckBalance
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	if input.MinConfirmations < 0 {
		input.MinConfirmations = 0
	}

	service := btc.NewBTCService()
	balance, err := service.CheckAddressBalance(input.Address, input.MinConfirmations)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedCheckBalance(err.Error())))
		return
	}

	balanceSatoshi := btc.BTCToSatoshi(balance)

	efwginhelpers.ResponseSuccess(context, gin.H{"data": responseCheckBalance{
		Address:          input.Address,
		Balance:          balance,
		BalanceSatoshi:   int64(balanceSatoshi),
		MinConfirmations: input.MinConfirmations,
	}})
}
