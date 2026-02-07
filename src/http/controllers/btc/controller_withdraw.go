package btc

import (
	"compay/src/services/cryptocurrencies/btc"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/gin-gonic/gin"
)

type requestWithdraw struct {
	Wif                    string        `json:"wif" binding:"required"`
	Address                string        `json:"address" binding:"required"`
	EstimatedFeeBlockCount int64         `json:"estimated_fee_block_count" binding:"required"`
	Amount                 entity.Amount `json:"amount" binding:"required,gt=0"`
}

func Withdraw(context *gin.Context) {
	var input requestWithdraw
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	response, err := btc.Withdraw(input.Wif, input.Address, input.EstimatedFeeBlockCount, btcutil.Amount(input.Amount))
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(err))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{
		"data": response,
	})
}
