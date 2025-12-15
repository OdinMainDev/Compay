package cperrors

import "fmt"

var ErrorAuthUserNotFound = fmt.Errorf("%s", "AuthUserNotFound")
var ErrorAuthPasswordDoNotMatch = fmt.Errorf("%s", "AuthPasswordDoNotMatch")
