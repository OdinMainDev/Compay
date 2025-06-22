package rates

import (
	"compay/application/helpers"
	"compay/application/service/rates"
	"compay/core/utils"

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
		helpers.GinSuccessResponse(context, gin.H{"data": []string{}})
		return
	}

	if !checkApiKey(input.ApiKey) {
		helpers.GinSuccessResponse(context, gin.H{"data": []string{}})
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"data": rates.GetRates()})
}

func GetRate(context *gin.Context) {
	var input requestGetRate
	if err := context.ShouldBindQuery(&input); err != nil {
		helpers.GinSuccessResponse(context, gin.H{"data": []string{}})
		return
	}

	if !checkApiKey(input.ApiKey) {
		helpers.GinSuccessResponse(context, gin.H{"data": []string{}})
		return
	}

	if input.CurrencyFrom == input.CurrencyTo {
		helpers.GinSuccessResponse(context, gin.H{"data": 1})
		return
	}

	data := rates.GetRates()
	if _, ok := data[input.CurrencyFrom]; !ok {
		helpers.GinErrorResponse(context, gin.H{"error": "Invalid fcur"})
		return
	}

	if _, ok := data[input.CurrencyFrom][input.CurrencyTo]; !ok {
		helpers.GinErrorResponse(context, gin.H{"error": "Invalid tcur or fcur-tcur pair"})
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"data": data[input.CurrencyFrom][input.CurrencyTo]})
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
