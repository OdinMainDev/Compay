package validator

import (
	"compay/common/constants"
	"compay/common/types"
	"efw/utils"

	"github.com/go-playground/validator/v10"
)

func validateLogAuthActionExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.LogAuthAction)
		if ok {
			return utils.InArray(data, constants.LogAuthActions)
		}
		return false
	}
	return "logAuthActionExist", fn
}

func validateLogUserActionExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(types.LogUserAction)
		if ok {
			return utils.InArray(data, constants.LogUserActions)
		}
		return false
	}
	return "logUserActionExist", fn
}
