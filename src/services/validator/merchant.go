package validator

import (
	"compay/common/types"
	"compay/src/models/merchant"

	"github.com/go-playground/validator/v10"
)

func validateMerchantExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.MerchantId)
		if ok {
			return merchant.GetModel(merchant.QueryHelper{Id: data}) != nil
		}
		return false
	}
	return "merchantExist", fn
}
