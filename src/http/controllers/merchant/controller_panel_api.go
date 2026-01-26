package merchant

import (
	"compay/common/constants"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

	"github.com/gin-gonic/gin"
)

func PanelApiCurrencyNetworks(context *gin.Context) {
	efwginhelpers.ResponseSuccess(context, gin.H{"data": constants.CurrencyNetworks})
}
