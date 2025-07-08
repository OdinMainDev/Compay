package wallet

import (
	"compay/application/entities/currency"
	"compay/application/helpers"

	"github.com/gin-gonic/gin"
)

func GetCurrencies(context *gin.Context) {
	helpers.GinSuccessResponse(context, gin.H{"data": currency.GetModels(currency.QueryHelper{})})
}
