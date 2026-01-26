package validator

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/utils"

	"github.com/go-playground/validator/v10"
)

func validateCurrencyNetworkExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.CurrencyNetwork)
		if ok {
			return utils.InArray(data, constants.CurrencyNetworks)
		}
		return false
	}
	return "currencyNetworkExist", fn
}
