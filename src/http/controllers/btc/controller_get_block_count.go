package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

// GetBlockCount получает текущую высоту блокчейна
func GetBlockCount(context *gin.Context) {
	client := btc.GetBTCClient()
	blockCount, err := client.GetBlockCount()
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedGetBlockCount(err.Error())))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": gin.H{
		"block_count": blockCount,
	}})
}
