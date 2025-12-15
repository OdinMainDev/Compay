package cperrors

import "fmt"

var ErrorAuthUserNotFound = fmt.Errorf("%s", "AuthUserNotFound")
var ErrorAuthPasswordDoNotMatch = fmt.Errorf("%s", "AuthPasswordDoNotMatch")
var ErrorAuthTotpExpired = fmt.Errorf("%s", "AuthTotpExpired")
