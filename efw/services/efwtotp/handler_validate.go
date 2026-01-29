package efwtotp

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func ValidateCode(secret, code string) error {
	opts := totp.ValidateOpts{
		Period:    Period,
		Skew:      Skew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	}
	_, err := totp.ValidateCustom(code, secret, time.Now(), opts)
	if err != nil {
		return ErrorValidateFail(err)
	}
	return nil
}
