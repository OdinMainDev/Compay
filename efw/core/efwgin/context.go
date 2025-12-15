package efwgin

import (
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ErrorSyntaxError = fmt.Errorf("%s", "[EFW: GIN] SchemaError")
var ErrorInvalidPayload = fmt.Errorf("%s", "[EFW: GIN] InvalidPayload")

func ShouldBindBodyWithJSON(context *gin.Context, input any) bool {
	return ShouldBindBody(context, input, binding.JSON)
}

// return ok (true) or not ok (false)
func ShouldBindBody(context *gin.Context, input any, bindingType binding.BindingBody) bool {
	err := context.ShouldBindBodyWith(&input, bindingType)
	if err == nil {
		return true
	}

	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		efwginhelpers.ResponseError(context, ErrorSyntaxError)
		return false
	}

	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		efwginhelpers.ResponseError(context, ErrorInvalidPayload)
		return false
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		ret := make(map[string][]string)
		for _, validatorErr := range validationErrors {
			ret[validatorErr.Field()] = append(ret[validatorErr.Field()], validatorErr.Tag())
		}
		efwginhelpers.ResponseError(context, ret)
		return false
	}

	return true
}
