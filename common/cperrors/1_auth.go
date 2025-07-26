package cperrors

import "fmt"

var ErrorAuthMerchantBadAuthToken = fmt.Errorf("%s", "Error1.1")
var ErrorAuthMerchantInvalidAuthToken = fmt.Errorf("%s", "Error1.2")
var ErrorAuthMerchantInvalidAuthToken2 = fmt.Errorf("%s", "Error1.3")
var ErrorAuthMerchantInvalidAuthToken3 = fmt.Errorf("%s", "Error1.4")
var ErrorAuthMerchantExpiredAuthToken = fmt.Errorf("%s", "Error1.5")
var ErrorAuthMerchantInvalidUser = fmt.Errorf("%s", "Error1.6")
var ErrorAuthTotpBadRequest = fmt.Errorf("%s", "Error1.7")
var ErrorAuthTotpTokenRequired = fmt.Errorf("%s", "Error1.8")
var ErrorAuthTotpDecryptDataError = fmt.Errorf("%s", "Error1.9")
var ErrorAuthTotpInvalidToken = fmt.Errorf("%s", "Error1.10")
var ErrorAuthTotpTokenExpired = fmt.Errorf("%s", "Error1.11")
