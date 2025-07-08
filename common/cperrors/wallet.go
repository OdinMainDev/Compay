package cperrors

import "fmt"

var ErrorWalletCreateError = fmt.Errorf("%s", "Error2.1")
var ErrorWalletCreateError2 = fmt.Errorf("%s", "Error2.2")
var ErrorWalletCreateMasterKey = fmt.Errorf("%s", "Error2.3")
var ErrorWalletInvalidMnemonic = fmt.Errorf("%s", "Error2.4")
var ErrorWalletCreateMasterKeyOnRestore = fmt.Errorf("%s", "Error2.4")
