package constants

import "compay/common/types"

// Log Auth
const (
	LogAuthLogin types.LogAuthAction = "Login"
)

var LogAuthActions = []types.LogAuthAction{LogAuthLogin}

// Log User
const (
	LogUserProfileUpdate types.LogUserAction = "ProfileUpdate"

	LogUserAdminUserCreate      types.LogUserAction = "AdminUserCreate"
	LogUserAdminUserProcess     types.LogUserAction = "AdminUserProcess"
	LogUserAdminUserDelete      types.LogUserAction = "AdminUserDelete"
	LogUserAdminUserDisableTotp types.LogUserAction = "AdminUserDisableTotp"
	LogUserAdminUserBan         types.LogUserAction = "AdminUserBan"
	LogUserAdminUserUnban       types.LogUserAction = "AdminUserUnban"

	LogUserAdminMerchantCreate  types.LogUserAction = "AdminMerchantCreate"
	LogUserAdminMerchantProcess types.LogUserAction = "AdminMerchantProcess"

	LogUserAdminCurrencyProcess types.LogUserAction = "AdminCurrencyProcess"
	LogUserAdminCurrencyDelete  types.LogUserAction = "AdminCurrencyDelete"
	LogUserAdminCurrencyImport  types.LogUserAction = "AdminCurrencyImport"

	LogUserNotifyRead    types.LogUserAction = "NotifyRead"
	LogUserNotifyReadAll types.LogUserAction = "NotifyReadAll"
)

var LogUserActions = []types.LogUserAction{
	LogUserProfileUpdate,
	LogUserAdminUserCreate,
	LogUserAdminUserProcess,
	LogUserAdminUserDelete,
	LogUserAdminUserDisableTotp,
	LogUserNotifyRead,
}
