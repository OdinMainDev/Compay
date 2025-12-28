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

var ErrorBindingSyntaxError = fmt.Errorf("%s", "[EFW: GIN BINDING] SchemaError")
var ErrorBindingInvalidPayload = fmt.Errorf("%s", "[EFW: GIN BINDING] InvalidPayload")
var ErrorBindingUndefinedError = func(fnName, err string) error {
	return fmt.Errorf("[EFW: GIN BINDING (%s)] %s: %s", fnName, "UndefinedError", err)
}

func ShouldBindBodyWithJSON(context *gin.Context, input any) bool {
	return ShouldBindBody(context, input, binding.JSON)
}

func ShouldBindWithForm(context *gin.Context, input any) bool {
	return ShouldBind(context, input, binding.Form)
}

// return ok (true) or not ok (false)
func ShouldBindBody(context *gin.Context, input any, bindingType binding.BindingBody) bool {
	err := context.ShouldBindBodyWith(&input, bindingType)
	if err == nil {
		return true
	}

	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		efwginhelpers.ResponseError(context, ErrorBindingSyntaxError)
		return false
	}

	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		efwginhelpers.ResponseError(context, ErrorBindingInvalidPayload)
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

	if err != nil {
		efwginhelpers.ResponseError(context, ErrorBindingUndefinedError("ShouldBindBody", err.Error()))
		return false
	}

	return true
}

func ShouldBind(context *gin.Context, input any, bindingType binding.Binding) bool {
	err := context.ShouldBindWith(&input, bindingType)
	if err == nil {
		return true
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

	if err != nil {
		efwginhelpers.ResponseError(context, ErrorBindingUndefinedError("ShouldBind", err.Error()))
		return false
	}

	return true
}
