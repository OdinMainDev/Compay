package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init(customValidations func() [](func() (string, validator.Func))) {
	if validatorEngine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validatorEngine.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		validatorEngine.RegisterValidation("email", validateEmail)
		validatorEngine.RegisterValidation("phone", validatePhone)
		if customValidations != nil {
			for _, v := range customValidations() {
				validatorEngine.RegisterValidation(v())
			}
		}
	}
}
