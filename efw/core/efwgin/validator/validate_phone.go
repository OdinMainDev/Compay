package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validatePhone validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		re := regexp.MustCompile(`^[1-9]?[0-9]{7,14}$`)
		return re.MatchString(data)
	}
	return false
}
