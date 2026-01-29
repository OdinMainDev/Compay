package efwtotp

import "fmt"

const TotpHeaderName = "X-Totp-Token"

var TotpSessionTime int64 = 1200 // время в секундах

const (
	SecretSize = 32
	Period     = 30
	Skew       = 0
)

var (
	ErrorNilHeader       = fmt.Errorf("%s", "[EFW: TOTP] NilHeader")
	ErrorTokenParseError = func(err error) error {
		return fmt.Errorf("%s: %s", "[EFW: TOTP] TokenParseError", err.Error())
	}
	ErrorTokenExpired = fmt.Errorf("%s", "[EFW: TOTP] TokenExpired")
	ErrorGenerateFail = func(err error) error {
		return fmt.Errorf("%s: %s", "[EFW: TOTP] GenerateFail", err.Error())
	}
	ErrorValidateFail = func(err error) error {
		return fmt.Errorf("%s: %s", "[EFW: TOTP] ValidateFail", err.Error())
	}
)
