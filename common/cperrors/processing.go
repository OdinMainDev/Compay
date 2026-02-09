package cperrors

import "fmt"

// PROCESSING +++++++++++++++++++
/* START PAYMENT */
var ErrorProcessingStartPaymentRequestError = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingStartPaymentRequestError:", err)
}
var ErrorProcessingStartPaymentUnmarshalError = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingStartPaymentUnmarshalError:", err)
}
var ErrorProcessingStartPaymentResponseError = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingStartPaymentResponseError:", err)
}

/* AMOUNT */
var ErrorProcessingDepositAmountIsRequired = fmt.Errorf("%s", "ProcessingDepositAmountIsRequired")
var ErrorProcessingDepositAmountLessThanMin = fmt.Errorf("%s", "ProcessingDepositAmountLessThanMin")
var ErrorProcessingDepositAmountBiggerThanMax = fmt.Errorf("%s", "ProcessingDepositAmountBiggerThanMax")

// PROCESSING -------------------

// MONITORING +++++++++++++++++++
var ErrorMonitoringBadResponseStatusCode = func(code int) error {
	return fmt.Errorf("%s %d", "MonitoringBadResponseStatusCode:", code)
}
var ErrorMonitoringBadResponseBody = func(err string) error {
	return fmt.Errorf("%s %s", "MonitoringBadResponseBody:", err)
}
var ErrorMonitoringErrorOnUpdateBtcBlock = func(err string) error {
	return fmt.Errorf("%s %s", "MonitoringErrorOnUpdateBtcBlock:", err)
}
var ErrorMonitoringErrorOnGetBtcTransactionsMarshal = func(err string) error {
	return fmt.Errorf("%s %s", "MonitoringErrorOnGetBtcTransactionsMarshal:", err)
}
var ErrorMonitoringErrorOnGetBtcTransactions = func(err string) error {
	return fmt.Errorf("%s %s", "MonitoringErrorOnGetBtcTransactions:", err)
}

// MONITORING -------------------

/* WITHDRAW */
var ErrorProcessingWithdrawAmountLessThanMinimal = fmt.Errorf("%s", "ProcessingWithdrawAmountLessThanMinimal")
var ErrorProcessingWithdrawAmountGreaterThanMaximum = fmt.Errorf("%s", "ProcessingWithdrawAmountGreaterThanMaximum")
var ErrorProcessingWithdrawHasUncompletedWithdraw = fmt.Errorf("%s", "ProcessingWithdrawHasUncompletedWithdraw")
var ErrorProcessingErrorOnConstructWithdrawRequest = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingErrorOnConstructWithdrawRequest:", err)
}
var ErrorProcessingWithdrawUnmarshalError = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingWithdrawUnmarshalError:", err)
}
var ErrorProcessingWithdrawResponseError = func(err string) error {
	return fmt.Errorf("%s %s", "ProcessingWithdrawResponseError:", err)
}
