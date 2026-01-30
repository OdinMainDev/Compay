package processing

import "fmt"

// START PAYMENT
var ErrorStartPaymentRequestError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentRequestError:", err)
}
var ErrorStartPaymentUnmarshalError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentUnmarshalError:", err)
}
var ErrorStartPaymentResponseError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentResponseError:", err)
}

// AMOUNT
var ErrorDepositAmountIsRequired = fmt.Errorf("%s", "DepositAmountIsRequired")
var ErrorDepositAmountLessThanMin = fmt.Errorf("%s", "DepositAmountLessThanMin")
var ErrorDepositAmountBiggerThanMax = fmt.Errorf("%s", "DepositAmountBiggerThanMax")
