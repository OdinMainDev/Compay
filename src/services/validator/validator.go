package validator

import "github.com/go-playground/validator/v10"

func CustomValidators() []func() (string, validator.Func) {
	var resp []func() (string, validator.Func)

	// Merchant
	resp = append(resp, validateMerchantExist)

	// User
	resp = append(resp, validateUserExist)

	// Currency
	resp = append(resp, validateCurrencyExist)

	// LOGS
	resp = append(resp, validateLogAuthActionExist)
	resp = append(resp, validateLogUserActionExist)

	// STATIC DATA VALIDATIONS
	resp = append(resp, validateCountryExist)
	resp = append(resp, validateCurrencyNetworkExist)
	resp = append(resp, validateFeeTypeExist)
	return resp
}
