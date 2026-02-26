package constants

import "compay/common/types"

// Log Auth
const (
	LogAuthLogin types.LogAuthAction = "Login"
	LogAuthTotp  types.LogAuthAction = "Totp"
)

var LogAuthActions = []types.LogAuthAction{
	LogAuthLogin,
	LogAuthTotp,
}

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
	LogUserAdminMerchantCreate          types.LogUserAction = "AdminMerchantCreate"
	LogUserAdminMerchantProcess         types.LogUserAction = "AdminMerchantProcess"
	LogUserAdminMerchantDisable         types.LogUserAction = "AdminMerchantDisable"
	LogUserAdminMerchantEnable          types.LogUserAction = "AdminMerchantEnable"
	LogUserAdminMerchantCurrencyProcess types.LogUserAction = "AdminMerchantCurrencyProcess"

	// ADMIN CURRENCY
	LogUserAdminCurrencyProcess types.LogUserAction = "AdminCurrencyProcess"
	LogUserAdminCurrencyImport  types.LogUserAction = "AdminCurrencyImport"
	LogUserAdminCurrencyDisable types.LogUserAction = "AdminCurrencyDisable"
	LogUserAdminCurrencyEnable  types.LogUserAction = "AdminCurrencyEnable"

	// ADMIN MERCHANT WALLET
	LogUserAdminMerchantWalletGetPrivateKey    types.LogUserAction = "AdminMerchantWalletGetPrivateKey"
	LogUserAdminMerchantWalletGetPublicKey     types.LogUserAction = "AdminMerchantWalletGetPublicKey"
	LogUserAdminMerchantWalletGetPublicKeyFull types.LogUserAction = "AdminMerchantWalletGetPublicKeyFull"

	// DEV ACTIONS
	LogUserAdminBlockchainStateProcess types.LogUserAction = "AdminBlockchainStateProcess"
	LogUserAdminBlockchainStateDelete  types.LogUserAction = "AdminBlockchainStateDelete"

	// MERCHANT
	// MERCHANT SETTINGS
	LogUserMerchantSettingsProcess        types.LogUserAction = "MerchantSettingsProcess"
	LogUserMerchantSettingsGenerateSecret types.LogUserAction = "MerchantSettingsGenerateSecret"

	// MERCHANT WALLETS
	LogUserMerchantWithdraw types.LogUserAction = "MerchantWithdraw"
)

var LogUserActions = []types.LogUserAction{
	LogUserProfileUpdate,
	LogUserNotifyRead,
	LogUserNotifyReadAll,
	LogUserAdminUserCreate,
	LogUserAdminUserProcess,
	LogUserAdminUserDelete,
	LogUserAdminUserDisableTotp,
	LogUserAdminUserBan,
	LogUserAdminUserUnban,
	LogUserAdminMerchantCreate,
	LogUserAdminMerchantProcess,
	LogUserAdminMerchantDisable,
	LogUserAdminMerchantEnable,
	LogUserAdminMerchantCurrencyProcess,
	LogUserAdminCurrencyProcess,
	LogUserAdminCurrencyImport,
	LogUserAdminCurrencyDisable,
	LogUserAdminCurrencyEnable,
	LogUserAdminMerchantWalletGetPrivateKey,
	LogUserAdminMerchantWalletGetPublicKey,
	LogUserAdminMerchantWalletGetPublicKeyFull,
	LogUserMerchantSettingsProcess,
	LogUserMerchantSettingsGenerateSecret,
	LogUserMerchantWithdraw,
}
