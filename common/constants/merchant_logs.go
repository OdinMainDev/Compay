package constants

import "compay/common/types"

// Log Auth
const (
	LogAuthLogin types.LogAuthAction = "Login"
	LogAuthTotp  types.LogAuthAction = "Totp"
)

var LogAuthActions = []types.LogAuthAction{LogAuthLogin}

// Log User
const (
	// COMMON
	LogUserProfileUpdate types.LogUserAction = "ProfileUpdate"
	LogUserNotifyRead    types.LogUserAction = "NotifyRead"
	LogUserNotifyReadAll types.LogUserAction = "NotifyReadAll"

	// ADMIN
	// ADMIN USER
	LogUserAdminUserCreate      types.LogUserAction = "AdminUserCreate"
	LogUserAdminUserProcess     types.LogUserAction = "AdminUserProcess"
	LogUserAdminUserDelete      types.LogUserAction = "AdminUserDelete"
	LogUserAdminUserDisableTotp types.LogUserAction = "AdminUserDisableTotp"
	LogUserAdminUserBan         types.LogUserAction = "AdminUserBan"
	LogUserAdminUserUnban       types.LogUserAction = "AdminUserUnban"

	// ADMIN MERCHANT
	LogUserAdminMerchantCreate  types.LogUserAction = "AdminMerchantCreate"
	LogUserAdminMerchantProcess types.LogUserAction = "AdminMerchantProcess"
	LogUserAdminMerchantDisable types.LogUserAction = "AdminMerchantDisable"
	LogUserAdminMerchantEnable  types.LogUserAction = "AdminMerchantEnable"

	// ADMIN CURRENCY
	LogUserAdminCurrencyProcess types.LogUserAction = "AdminCurrencyProcess"
	LogUserAdminCurrencyImport  types.LogUserAction = "AdminCurrencyImport"

	// MERCHANT
	// MERCHANT SETTINGS
	LogUserMerchantSettingsProcess        types.LogUserAction = "MerchantSettingsProcess"
	LogUserMerchantSettingsGenerateSecret types.LogUserAction = "MerchantSettingsGenerateSecret"
)

var LogUserActions = []types.LogUserAction{
	LogUserProfileUpdate,
	LogUserAdminUserCreate,
	LogUserAdminUserProcess,
	LogUserAdminUserDelete,
	LogUserAdminUserDisableTotp,
	LogUserNotifyRead,
}
