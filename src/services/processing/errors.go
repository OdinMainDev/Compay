package processing

import "fmt"

// START PAYMENT
var ErrorProcessingStartPaymentRequestError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentRequestError:", err)
}
var ErrorProcessingStartPaymentUnmarshalError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentUnmarshalError:", err)
}
var ErrorProcessingStartPaymentResponseError = func(err string) error {
	return fmt.Errorf("%s %s", "StartPaymentResponseError:", err)
}
