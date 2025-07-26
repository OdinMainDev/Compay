package totp

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

const (
	SecretSize = 32
	Period     = 30
	Skew       = 0
)

func GenerateSecret(accountName, issuer string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
		SecretSize:  SecretSize,
		Period:      Period,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate TOTP key: %w", err)
	}
	return key.Secret(), nil
}

// генерирует otpauth:// URI
// AccountName это емейл или юзернейм
func OtpUri(secret, accountName, issuer string) string {
	label := url.QueryEscape(fmt.Sprintf("%s:%s", issuer, accountName))
	params := url.Values{}
	params.Set("secret", secret)
	params.Set("issuer", issuer)
	params.Set("algorithm", "SHA1")
	params.Set("digits", "6")
	params.Set("period", "30")

	return fmt.Sprintf("otpauth://totp/%s?%s", label, params.Encode())
}

// Проверка кода
func ValidateCode(secret, code string) (bool, error) {
	opts := totp.ValidateOpts{
		Period:    Period,
		Skew:      Skew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	}
	ok, err := totp.ValidateCustom(code, secret, time.Now().UTC(), opts)
	if err != nil {
		return false, fmt.Errorf("error validating TOTP code: %w", err)
	}
	return ok, nil
}
