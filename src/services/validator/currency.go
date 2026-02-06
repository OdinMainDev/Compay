package validator

import (
	"compay/src/models/currency"
	"efw/components/entity"

	"github.com/go-playground/validator/v10"
)

func validateCurrencyExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(entity.CurrencyCodeAlpha)
		if ok {
			return currency.GetModel(currency.QueryHelper{CodeAlpha: data}) != nil
		}
		return false
	}
	return "currencyExist", fn
}
