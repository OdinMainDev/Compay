package constants

import "compay/common/types"

// Log Auth
const (
	LogAuthLogin types.LogAuthAction = "Login"
)

// Log User
const (
	LogUserProfileUpdate        types.LogUserAction = "ProfileUpdate"
	LogUserAdminUserCreate      types.LogUserAction = "AdminUserCreate"
	LogUserAdminUserProcess     types.LogUserAction = "AdminUserProcess"
	LogUserAdminUserDelete      types.LogUserAction = "AdminUserDelete"
	LogUserAdminUserDisableTotp types.LogUserAction = "AdminUserDisableTotp"
)
