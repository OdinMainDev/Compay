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
