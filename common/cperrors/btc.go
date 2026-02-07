package cperrors

import "fmt"

var ErrorBtcFailedGeneratePrivateKey = func(err string) error {
	return fmt.Errorf("%s %s", "BtcFailedGeneratePrivateKey:", err)
}

var ErrorBtcFailedCreateWif = func(err string) error {
	return fmt.Errorf("%s %s", "BtcFailedCreateWif:", err)
}

var ErrorBtcFailedCreateAddress = func(err string) error {
	return fmt.Errorf("%s %s", "BtcFailedCreateAddress:", err)
}
var ErrorBtcFailedGetBlockCount = func(err string) error {
	return fmt.Errorf("%s %s", "BtcFailedGetBlockCount:", err)
}
var ErrorBtcFailedGetTransactions = func(err string) error {
	return fmt.Errorf("%s %s", "BtcFailedGetTransactions:", err)
}
var ErrorBtcWithdrawFailedDecodePrivateWif = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedDecodePrivateWif:", err)
}
var ErrorBtcWithdrawFailedCreateNewAddress = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedCreateNewAddress:", err)
}
var ErrorBtcWithdrawFailedGetListUnspent = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedGetListUnspent:", err)
}
var ErrorBtcWithdrawFailedBadResponseListUnspent = fmt.Errorf("%s", "BtcWithdrawFailedBadResponseListUnspent:")
var ErrorBtcWithdrawFailedGetEstimatedFee = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedGetEstimatedFee:", err)
}
var ErrorBtcWithdrawInsufficientFunds = func(have, need int64) error {
	return fmt.Errorf("%s %d %s %d", "BtcWithdrawInsufficientFunds: Have", have, "Need", need)
}
var ErrorBtcWithdrawInsufficientFundsForFee = func(have, need, fee int64) error {
	return fmt.Errorf("%s %d %s %d %s %d", "BtcWithdrawInsufficientFunds: Have", have, "Need", need, "Fee", fee)
}
var ErrorBtcWithdrawFailedToParseTxId = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToParseTxId:", err)
}
var ErrorBtcWithdrawFailedToDecodeDesitnationAddress = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToDecodeDesitnationAddress:", err)
}
var ErrorBtcWithdrawFailedToCreateDestinationScript = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToCreateDestinationScript:", err)
}
var ErrorBtcWithdrawFailedToCreateChangeScript = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToCreateChangeScript:", err)
}
var ErrorBtcWithdrawFailedToDecodeScriptPublicKey = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToDecodeScriptPublicKey:", err)
}
var ErrorBtcWithdrawFailedToCreateWitnessSignature = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToCreateWitnessSignature:", err)
}
var ErrorBtcWithdrawFailedToSerializeTransaction = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToSerializeTransaction:", err)
}
var ErrorBtcWithdrawFailedToSendTransaction = func(err string) error {
	return fmt.Errorf("%s %s", "BtcWithdrawFailedToSendTransaction:", err)
}
