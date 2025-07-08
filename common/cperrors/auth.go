package cperrors

import "fmt"

var ErrorAuthUserTokenIsNull = fmt.Errorf("%s", "Error1.1")
var ErrorAuthUserTokenIsInvalid = fmt.Errorf("%s", "Error1.2")
var ErrorAuthBadUserToken = fmt.Errorf("%s", "Error1.3")
var ErrorAuthUserBanned = fmt.Errorf("%s", "Error1.4")
var ErrorAuthJwtTokenIsNull = fmt.Errorf("%s", "Error1.5")
var ErrorAuthJwtTokenInvalid = fmt.Errorf("%s", "Error1.6")
var ErrorAuthJwtTokenInvalidHashMethod = fmt.Errorf("%s", "Error1.7")
var ErrorAuthJwtTokenClaimsNotValid = fmt.Errorf("%s", "Error1.8")
var ErrorAuthJwtTokenIdOrStkIsNull = fmt.Errorf("%s", "Error1.9")
var ErrorAuthJwtTokenFailedDecryptSecretToken = fmt.Errorf("%s", "Error1.10")
var ErrorAuthJwtTokenMismatchSecretTokens = fmt.Errorf("%s", "Error1.11")
var ErrorAuthJwtTokenMismatchPinCodes = fmt.Errorf("%s", "Error1.12")
