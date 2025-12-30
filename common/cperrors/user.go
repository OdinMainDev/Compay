package cperrors

import "fmt"

var ErrorUserUsernameIsTaken = fmt.Errorf("%s", "UsernameIsTaken")
var ErrorUserCreateAdminWithMerchantIdIsNotAllowed = fmt.Errorf("%s", "CreateAdminWithMerchantIdIsNotAllowed")
var ErrorUserCreateUserWithoutMerchantIdIsNotAllowed = fmt.Errorf("%s", "CreateUserWithoutMerchantIdIsNotAllowed")
