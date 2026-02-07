package cperrors

import "fmt"

var ErrorApiV1BadCurrency = fmt.Errorf("%s", "ApiV1BadCurrency")
var ErrorApiV1BadMerchantToken = fmt.Errorf("%s", "ApiV1BadMerchantToken")
var ErrorApiV1InvalidSign = fmt.Errorf("%s", "ApiV1InvalidSign")
var ErrorApiV1CurrencyOnMaintenance = fmt.Errorf("%s", "CurrencyOnMaintenance")
var ErrorApiV1WithdrawAmountLessThanMinimal = fmt.Errorf("%s", "ErrorApiV1WithdrawAmountLessThanMinimal")
var ErrorApiV1WithdrawAmountGreaterThanMaximum = fmt.Errorf("%s", "ErrorApiV1WithdrawAmountGreaterThanMaximum")
var ErrorApiV1ErrorOnConstructWithdrawRequest = func(err string) error {
	return fmt.Errorf("%s %s", "ApiV1ErrorOnConstructWithdrawRequest:", err)
}
