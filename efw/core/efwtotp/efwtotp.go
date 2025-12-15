package efwtotp

import "fmt"

var (
	TotpHeaderName        = "X-Totp-Token"
	TotpSessionTime int64 = 1200 // время в секундах
)

var (
	ErrorTotpNilHeader       = fmt.Errorf("%s", "[EFW: TOTP] NilHeader")
	ErrorTotpTokenParseError = func(err error) error {
		return fmt.Errorf("%s: %s", "[EFW: TOTP] TokenParseError", err.Error())
	}
	ErrorTotpTokenExpired = fmt.Errorf("%s", "[EFW: TOTP] TokenExpired")
)
