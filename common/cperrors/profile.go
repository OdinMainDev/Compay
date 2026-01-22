package cperrors

import "fmt"

var ErrorProfileInvalidOldUsername = fmt.Errorf("%s", "ProfileInvalidOldUsername")
var ErrorProfileNewUsernameIsTaken = fmt.Errorf("%s", "ProfileNewUsernameIsTaken")
