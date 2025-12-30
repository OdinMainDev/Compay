package validator

import (
	"compay/common/types"
	"compay/src/models/merchant"

	"github.com/go-playground/validator/v10"
)

func validateMerchantExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(int64)
		if ok {
			return merchant.GetModel(merchant.QueryHelper{Id: types.UserId(data)}) != nil
		}
		return false
	}
	return "merchantExist", fn
}
