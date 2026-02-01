package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type GetTransactionsRequest struct {
	LastHeight int64 `json:"last_height" binding:"required"`
}

func GetTransactions(context *gin.Context) {
	var input GetTransactionsRequest
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	transactions, newHeight, err := btc.GetTransactions(input.LastHeight)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedGetTransactions(err.Error())))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{
		"data": gin.H{
			"transactions": transactions,
			"new_height":   newHeight,
		},
	})
}
