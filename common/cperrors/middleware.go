package cperrors

import "fmt"

// CheckUser
var ErrorMiddlewareCheckUserNilJwtData = fmt.Errorf("%s", "MiddlewareCheckUserNilJwtData")
var ErrorMiddlewareCheckUserInvalidJwtData = fmt.Errorf("%s", "MiddlewareCheckUserInvalidJwtData")
var ErrorMiddlewareCheckUserNilUserId = fmt.Errorf("%s", "MiddlewareCheckUserNilUserId")
var ErrorMiddlewareCheckUserInvalidUserId = fmt.Errorf("%s", "MiddlewareCheckUserInvalidUserId")
var ErrorMiddlewareCheckUserInvalidUserModel = fmt.Errorf("%s", "MiddlewareCheckUserInvalidUserModel")
var ErrorMiddlewareCheckUserBannedUserModel = fmt.Errorf("%s", "MiddlewareCheckUserBannedUserModel")

// CheckAuthTotp
var ErrorMiddlewareCheckAuthTotpNilUserModel = fmt.Errorf("%s", "MiddlewareCheckAuthTotpNilUserModel")
var ErrorMiddlewareCheckAuthTotpInvalidTypeUserModel = fmt.Errorf("%s", "MiddlewareCheckAuthTotpInvalidTypeUserModel")
