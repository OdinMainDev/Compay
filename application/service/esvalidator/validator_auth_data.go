package esvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/tyler-smith/go-bip39"
)

var pinCode validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if len(data) != 6 {
		return false
	}

	for _, r := range data {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

var seedPhrase validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	if !bip39.IsMnemonicValid(data) {
		return false
	}

	return true
}
