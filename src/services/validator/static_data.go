package validator

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/utils"

	"github.com/go-playground/validator/v10"
)

func validateCountryExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.Country)
		if ok {
			return utils.InArray(data, constants.Countries)
		}
		return false
	}
	return "countryExist", fn
}

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

func validateFeeTypeExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.FeeType)
		if ok {
			return utils.InArray(data, constants.FeeTypes)
		}
		return false
	}
	return "feeTypeExist", fn
}

func validatePaymentStatusExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.PaymentStatus)
		if ok {
			return utils.InArray(data, constants.PaymentStatuses)
		}
		return false
	}
	return "paymentStatusExist", fn
}

func validateAmount() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(float64)
		currency := fl.Parent().FieldByName("Currency")
		if !ok {
			return false
		}

		if data > 0 && currency.IsZero() {
			return false
		}

		return true
	}
	return "amount", fn
}
