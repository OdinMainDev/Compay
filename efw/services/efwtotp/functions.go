package efwtotp

import (
	"fmt"
	"net/url"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
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
		return "", ErrorGenerateFail(err)
	}
	return key.Secret(), nil
}

// генерирует otpauth:// URI
// AccountName это емейл или юзернейм
// issuer - имя проекта
func ConstructUri(secret, accountName, issuer string) string {
	label := url.QueryEscape(fmt.Sprintf("%s:%s", issuer, accountName))
	params := url.Values{}
	params.Set("secret", secret)
	params.Set("issuer", issuer)
	params.Set("algorithm", "SHA1")
	params.Set("digits", "6")
	params.Set("period", "30")

	return fmt.Sprintf("otpauth://totp/%s?%s", label, params.Encode())
}
