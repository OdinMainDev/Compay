package btc

import (
	"compay/common/cperrors"
	"compay/src/services/cryptocurrencies/btc"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

type requestImportAddress struct {
	Address string `json:"address" binding:"required"`
	Label   string `json:"label"`
	Rescan  bool   `json:"rescan"`
}

// ImportAddress импортирует адрес в BTC ноду для отслеживания
func ImportAddress(context *gin.Context) {
	var input requestImportAddress
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	service := btc.NewBTCService()
	err := service.ImportPaymentAddress(input.Address, input.Label)
	if err != nil {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorBtcFailedImportAddress(err.Error())))
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": gin.H{
		"address": input.Address,
		"label":   input.Label,
		"message": "Address imported successfully",
	}})
}
