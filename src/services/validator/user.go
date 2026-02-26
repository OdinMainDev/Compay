package validator

import (
	"compay/src/models/user"
	"efw/components/entity"

	"github.com/go-playground/validator/v10"
)

func validateUserExist() (string, validator.Func) {
	var fn validator.Func = func(fl validator.FieldLevel) bool {
		data, ok := fl.Field().Interface().(entity.UserId)
		if ok {
			return user.GetModel(user.QueryHelper{Id: data, DontCheckMerchantId: true}) != nil
		}
		return false
	}
	return "userExist", fn
}
