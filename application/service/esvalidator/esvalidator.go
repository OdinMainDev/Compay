package esvalidator

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Initialize() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		v.RegisterValidation("pinCode", pinCode)
		v.RegisterValidation("seedPhrase", seedPhrase)
	}
}

func ValidatorResponse(err error) map[string][]string {
	ret := make(map[string][]string)
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		ret["_schema"] = []string{"SyntaxError"}
		return ret
	}
	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		ret[unmarshalError.Field] = []string{"WrongType"}
		return ret
	}
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, validatorErr := range validationErrors {
			ret[validatorErr.Field()] = append(ret[validatorErr.Field()], validatorErr.Tag())
		}
	}
	return ret
}
