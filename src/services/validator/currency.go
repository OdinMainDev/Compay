package validator

import (
	"compay/common/types"
	"compay/src/models/currency"

	"github.com/go-playground/validator/v10"
)

func validateCurrencyExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.CurrencyCodeAlpha)
		if ok {
			return currency.GetModel(currency.QueryHelper{CodeAlpha: data}) != nil
		}
		return false
	}
	return "currencyExist", fn
}
