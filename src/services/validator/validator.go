package validator

import "github.com/go-playground/validator/v10"

func CustomValidators() []func() (string, validator.Func) {
	var resp []func() (string, validator.Func)
	resp = append(resp, validateMerchantExist)
	return resp
}
