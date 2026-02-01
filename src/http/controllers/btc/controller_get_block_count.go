package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func GetBlockCount(context *gin.Context) {
	blockCount, err := btc.GetBlockCount()
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedGetBlockCount(err.Error())))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{
		"data": blockCount,
	})
}
