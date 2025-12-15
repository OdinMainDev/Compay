package rates

import (
	"compay/src/services/ratesparser"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/utils"

	"github.com/gin-gonic/gin"
)

type requestGetRates struct {
	ApiKey string `form:"apiKey"`
}

type requestGetRate struct {
	ApiKey       string `form:"apiKey"`
	CurrencyFrom string `form:"fcur"`
	CurrencyTo   string `form:"tcur"`
}

func GetRates(context *gin.Context) {
	var input requestGetRates
	if err := context.ShouldBindQuery(&input); err != nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": []string{}})
		return
	}

	if !checkApiKey(input.ApiKey) {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": []string{}})
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": ratesparser.GetRates()})
}

func GetRate(context *gin.Context) {
	var input requestGetRate
	if err := context.ShouldBindQuery(&input); err != nil {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": []string{}})
		return
	}

	if !checkApiKey(input.ApiKey) {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": []string{}})
		return
	}

	if input.CurrencyFrom == input.CurrencyTo {
		efwginhelpers.ResponseSuccess(context, gin.H{"data": 1})
		return
	}

	data := ratesparser.GetRates()
	if _, ok := data[input.CurrencyFrom]; !ok {
		efwginhelpers.ResponseError(context, gin.H{"error": "Invalid fcur"})
		return
	}

	if _, ok := data[input.CurrencyFrom][input.CurrencyTo]; !ok {
		efwginhelpers.ResponseError(context, gin.H{"error": "Invalid tcur or fcur-tcur pair"})
		return
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": data[input.CurrencyFrom][input.CurrencyTo]})
}

func checkApiKey(apiKey string) bool {
	if apiKey == "" {
		return false
	}

	if !utils.InArray(apiKey, []string{
		"w1ZvhUJI89XvIeCMF3RXWmeBxArrSDWY",
		"szxzLzUB6hvL2eJLBjqr5Qeeu5d7hAN1",
		"BcZkXoolIT7RVOvTbb9pyIBMWY78q0Ck",
	}) {
		return false
	}

	return true
}
