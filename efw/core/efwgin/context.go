package efwgin

import (
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var ErrorBindingSyntaxError = func(err string) error {
	return fmt.Errorf("[EFW: GIN BINDING] %s: %s", "SchemaError", err)
}
var ErrorBindingInvalidPayload = func(err string) error {
	return fmt.Errorf("[EFW: GIN BINDING] %s: %s", "InvalidPayload", err)
}
var ErrorBindingUndefinedError = func(fnName, err string) error {
	return fmt.Errorf("[EFW: GIN BINDING (%s)] %s: %s", fnName, "UndefinedError", err)
}

func ShouldBindBodyWithJSON(context *gin.Context, input any) bool {
	return shouldBindBody(context, input, binding.JSON)
}

func ShouldBindWithForm(context *gin.Context, input any) bool {
	return shouldBind(context, input, binding.Form)
}

func ShouldBindUri(context *gin.Context, input any) bool {
	return shouldBindUri(context, input)
}

// return ok (true) or not ok (false)
func shouldBindBody(context *gin.Context, input any, bindingType binding.BindingBody) bool {
	err := context.ShouldBindBodyWith(&input, bindingType)
	if err == nil {
		return true
	}

	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		efwginhelpers.ResponseError(context, ErrorBindingSyntaxError(err.Error()))
		return false
	}

	var unmarshalError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalError) {
		efwginhelpers.ResponseError(context, ErrorBindingInvalidPayload(err.Error()))
		return false
	}

	return validationResponse(context, err)
}

// return ok (true) or not ok (false)
func shouldBind(context *gin.Context, input any, bindingType binding.Binding) bool {
	err := context.ShouldBindWith(&input, bindingType)
	if err == nil {
		return true
	}

	return validationResponse(context, err)
}

// return ok (true) or not ok (false)
func shouldBindUri(context *gin.Context, input any) bool {
	err := context.ShouldBindUri(input)
	if err == nil {
		return true
	}

	return validationResponse(context, err)
}

func validationResponse(ctx *gin.Context, err error) bool {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		if err != nil {
			efwginhelpers.ResponseError(
				ctx,
				ErrorBindingUndefinedError("validationResponse", err.Error()),
			)
			return false
		}
		return true
	}

	ret := make(map[string][]string)
	for _, val := range validationErrors {
		ns := val.Namespace()

		nsParts := strings.SplitN(ns, ".", 2)
		if len(nsParts) != 2 {
			continue
		}
		path := nsParts[1]
		parts := strings.Split(path, ".")
		key := strings.Split(parts[0], "[")[0]

		index := ""
		if strings.Contains(parts[0], "[") {
			index = parts[0][strings.Index(parts[0], "["):]
		}

		var message string

		if len(parts) == 1 {
			message = val.Tag()
		} else {
			message = fmt.Sprintf(
				"%s%s;%s",
				parts[1],
				index,
				val.Tag(),
			)
		}

		ret[key] = append(ret[key], message)
	}

	efwginhelpers.ResponseError(ctx, ret)
	return false
}
