package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validateEmail validator.Func = func(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if ok {
		return ValidateEmail(data)
	}
	return false
}

var (
	emailUsernameRegexp = regexp.MustCompile("^[a-zA-Z0-9_.-]+$")
	emailHostRegexp     = regexp.MustCompile(`^[a-zA-Z][^\s]+\.[^\s]+$`)
	emailDotRegexp      = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
)

func ValidateEmail(email string) bool {
	if len(email) < 6 || len(email) > 254 {
		return false
	}

	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return false
	}

	user := email[:at]
	host := email[at+1:]

	if len(user) > 64 {
		return false
	}
	if emailDotRegexp.MatchString(user) || !emailUsernameRegexp.MatchString(user) || !emailHostRegexp.MatchString(host) {
		return false
	}

	return true
}
