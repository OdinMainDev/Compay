package cperrors

import "fmt"

var ErrorProfileInvalidOldUsername = fmt.Errorf("%s", "ProfileInvalidOldUsername")
var ErrorProfileNewUsernameIsTaken = fmt.Errorf("%s", "ProfileNewUsernameIsTaken")
var ErrorProfileTotpAlreadyConfigured = fmt.Errorf("%s", "ProfileTotpAlreadyConfigured")
var ErrorProfileTotpNotConfigured = fmt.Errorf("%s", "ProfileTotpNotConfigured")
var ErrorProfileTotpAlreadyEnabled = fmt.Errorf("%s", "ProfileTotpAlreadyEnabled")
