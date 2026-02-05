package cperrors

import "fmt"

var ErrorCConverterOnStart = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterOnStart:", err)
}
var ErrorCConverterRequestFail = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterRequestFail:", err)
}
var ErrorCConverterInvalidResponse = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterInvalidResponse:", err)
}
var ErrorCConverterBadResponse = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterBadResponse:", err)
}
var ErrorCConverterCurrencyNotExist = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterCurrencyNotExist:", err)
}
var ErrorCConverterUsdForCurrencyNotExist = func(err string) error {
	return fmt.Errorf("%s %s", "CConverterUsdForCurrencyNotExist:", err)
}
