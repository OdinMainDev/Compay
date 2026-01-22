package cperrors

import "fmt"

var ErrorUserUsernameIsTaken = fmt.Errorf("%s", "UserUsernameIsTaken")
var ErrorUserCreateAdminWithMerchantIdIsNotAllowed = fmt.Errorf("%s", "UserCreateAdminWithMerchantIdIsNotAllowed")
var ErrorUserCreateUserWithoutMerchantIdIsNotAllowed = fmt.Errorf("%s", "UserCreateUserWithoutMerchantIdIsNotAllowed")
