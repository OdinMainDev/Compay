package app

import (
	"compay/common/config"
	"context"
	"efw/core/efwgin"
	efwmiddleware "efw/core/efwgin/efw_middleware"
	"efw/core/efwlog"
	"errors"
	"net/http"

	controller "compay/src/http/controllers/merchant"
	"compay/src/http/middleware"
	"compay/src/services/processing"

	"github.com/gin-gonic/gin"
)

var merchantServer *http.Server

func StartMerchant() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.POST("/init-app", controller.InitApp)

	apiRoutes := router.Group("/api")
	apiRoutes.Use(middleware.CheckIsAppInit())
	apiRoutes.Use(efwmiddleware.CheckClientIp())

	/* +++++++++++++++ AUTH ROUTES +++++++++++++++ */
	authRoutes := apiRoutes.Group("/auth")

	merchantAuthRoutes(authRoutes)
	/* --------------- AUTH ROUTES --------------- */

	/* +++++++++++++++ PANEL ROUTES +++++++++++++++ */
	panelRoutes := apiRoutes.Group("/panel")
	panelRoutes.Use(efwmiddleware.CheckAuth(), middleware.CheckUser(), middleware.CheckAuthTotp())

	merchantPanelProfileRoutes(panelRoutes)
	merchantPanelNotificationsRoutes(panelRoutes)
	merchantPanelApiRoutes(panelRoutes)

	/* +++++++++++++++ ADMIN ROUTES +++++++++++++++ */
	adminRoutes := panelRoutes.Group("/admin")
	adminRoutes.Use(middleware.CheckIsUserAdmin())
	merchantAdminUsersRoutes(adminRoutes)
	merchantAdminMerchantsRoutes(adminRoutes)
	merchantAdminMerchantCurrenciesRoutes(adminRoutes)
	merchantAdminCurrenciesRoutes(adminRoutes)
	merchantAdminPaymentsRoutes(adminRoutes)
	merchantAdminMerchantWithdrawsRoutes(adminRoutes)
	merchantAdminMerchantWalletsRoutes(adminRoutes)

	merchantAdminApiRoutes(adminRoutes)

	// LOGS
	logGroup := adminRoutes.Group("/dev/logs")
	merchantAdminLogAuthRoutes(logGroup)
	merchantAdminLogUserRoutes(logGroup)
	merchantAdminLogTransactionRoutes(logGroup)
	merchantAdminLogMerchantWalletRoutes(logGroup)
	/* --------------- ADMIN ROUTES --------------- */

	/* --------------- PANEL ROUTES --------------- */

	/* +++++++++++++++ MERCHANT ROUTES +++++++++++++++ */
	merchantRoutes := panelRoutes.Group("/merchant")
	merchantRoutes.Use(middleware.CheckMerchant())

	merchantMerchantRoutes(merchantRoutes)
	merchantMerchantWalletRoutes(merchantRoutes)
	merchantMerchantWithdrawsRoutes(merchantRoutes)

	/* --------------- MERCHANT ROUTES --------------- */

	/* +++++++++++++++ V1 ROUTES +++++++++++++++ */
	v1 := apiRoutes.Group("/v1")

	merchantPublicApiV1Routes(v1)
	/* --------------- V1 ROUTES --------------- */
	merchantServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := merchantServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			efwlog.LogError("Merchant server started with error: ", err.Error())
		}
	}()
}

func StopMerchant() {
	processing.StopProcessing()
	err := merchantServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Merchant server ended with error:", err.Error())
	}
}

func merchantAuthRoutes(group *gin.RouterGroup) {
	group.POST("/login", efwmiddleware.RateLimit(5), controller.AuthLogin)
	group.POST("/totp", efwmiddleware.RateLimit(5), efwmiddleware.CheckAuth(), middleware.CheckUser(), controller.AuthTotp)
}

// PANEL ROUTES

func merchantPanelProfileRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/profile")

	subGroup.POST("/get", controller.ProfileGet)
	subGroup.POST("/process", controller.ProfileProcess)
	subGroup.POST("/totp-uri", controller.ProfileTotpUri)
	subGroup.POST("/totp-activate", controller.ProfileTotpActivate)
}

func merchantPanelNotificationsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/notifications")

	subGroup.POST("/list", controller.NotificationsList)
	subGroup.POST("/list/filters", controller.NotificationsListFilters)
	subGroup.POST("/get/unread", controller.NotificationsGetUnread)
	subGroup.POST("/read/:id", controller.NotificationsRead)
	subGroup.POST("/read/all", controller.NotificationsReadAll)
}

func merchantPanelApiRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/get/data")

	subGroup.POST("/currency/networks", controller.PanelApiCurrencyNetworks)
	subGroup.POST("/countries", controller.PanelApiCountries)
}

// ADMIN ROUTES

func merchantAdminUsersRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/users")

	subGroup.POST("/list", controller.AdminUsersList)
	subGroup.POST("/list/filters", controller.AdminUsersListFilters)
	subGroup.POST("/get/:id", controller.AdminUsersGet)
	subGroup.POST("/create", controller.AdminUsersCreate)
	subGroup.POST("/process", controller.AdminUsersProcess)
	subGroup.POST("/search", controller.AdminUsersSearch)
	subGroup.POST("/action/disable-totp", controller.AdminUsersActionDisableTotp)
	subGroup.POST("/action/ban", controller.AdminUsersActionBan)
	subGroup.POST("/action/unban", controller.AdminUsersActionUnban)
}

func merchantAdminMerchantsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/merchants")

	subGroup.POST("/list", controller.AdminMerchantsList)
	subGroup.POST("/list/filters", controller.AdminMerchantsListFilters)
	subGroup.POST("/get/:id", controller.AdminMerchantsGet)
	subGroup.POST("/create", controller.AdminMerchantsCreate)
	subGroup.POST("/process", controller.AdminMerchantsProcess)
	subGroup.POST("/search", controller.AdminMerchantsSearch)
	subGroup.POST("/action/disable", controller.AdminMerchantsActionDisable)
	subGroup.POST("/action/enable", controller.AdminMerchantsActionEnable)
}

func merchantAdminMerchantCurrenciesRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/merchant/currencies")

	subGroup.POST("/get/:merchantId/:id", controller.AdminMerchantCurrenciesGet)
	subGroup.POST("/get/all/:id", controller.AdminMerchantCurrenciesGetAll)
	subGroup.POST("/process", controller.AdminMerchantCurrenciesProcess)
}

func merchantAdminApiRoutes(group *gin.RouterGroup) {
	getGroup := group.Group("/get/data")
	actionGroup := group.Group("/action")

	getGroup.POST("/fee-types", controller.AdminApiFeeTypes)

	actionGroup.POST("/clear-cache", controller.AdminApiClearCache)
	actionGroup.POST("/start-merchant", controller.AdminApiStartMerchant)
	actionGroup.POST("/stop-merchant", controller.AdminApiStopMerchant)
}

func merchantAdminCurrenciesRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/currencies")

	subGroup.POST("/list", controller.AdminCurrenciesList)
	subGroup.POST("/list/filters", controller.AdminCurrenciesListFilters)
	subGroup.POST("/get/:id", controller.AdminCurrenciesGet)
	subGroup.POST("/process", controller.AdminCurrenciesProcess)
	subGroup.POST("/search", controller.AdminCurrenciesSearch)
	subGroup.POST("/action/import", controller.AdminCurrenciesActionImport)
	subGroup.POST("/action/disable", controller.AdminCurrenciesActionDisable)
	subGroup.POST("/action/enable", controller.AdminCurrenciesActionEnable)
}

func merchantAdminPaymentsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/payments")

	subGroup.POST("/get/:id", controller.AdminPaymentsGet)
}

func merchantAdminMerchantWithdrawsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/withdraws")

	subGroup.POST("/get-all/:id", controller.AdminMerchantWithdrawsGetAll)
}

func merchantAdminMerchantWalletsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/merchant-wallets")

	subGroup.POST("/list", controller.AdminMerchantWalletsList)
	subGroup.POST("/list/filters", controller.AdminMerchantWalletsListFilters)
	subGroup.POST("/get/:id", controller.AdminMerchantWalletsGet)
	subGroup.POST("/action/get-private", controller.AdminMerchantActionGetPrivateKey)
	subGroup.POST("/action/get-public", controller.AdminMerchantActionGetPublicKey)
	subGroup.POST("/action/get-public-full", controller.AdminMerchantActionGetPublicKeyFull)
}

func merchantAdminLogAuthRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/auth")

	subGroup.POST("/list", controller.AdminLogAuthList)
	subGroup.POST("/list/filters", controller.AdminLogAuthListFilters)
}

func merchantAdminLogUserRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/user")

	subGroup.POST("/list", controller.AdminLogUserList)
	subGroup.POST("/list/filters", controller.AdminLogUserListFilters)
}

func merchantAdminLogTransactionRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/transactions")

	subGroup.POST("/list/:id", controller.AdminLogTransactionList)
}

func merchantAdminLogMerchantWalletRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/merchant-wallet")

	subGroup.POST("/list/:id", controller.AdminLogMerchantWalletList)
}

// MERCHANT ROUTES

func merchantMerchantRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/merchant")

	subGroup.POST("/get", controller.MerchantSettingsGet)
	subGroup.POST("/process", controller.MerchantSettingsProcess)
	subGroup.POST("/action/get-token", controller.MerchantSettingsActionGetToken)
	subGroup.POST("/action/generate-secret", controller.MerchantSettingsActionGenerateSecret)
}

func merchantMerchantWalletRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/wallets")

	subGroup.POST("/list", controller.MerchantWalletsList)
	subGroup.POST("/list/filters", controller.MerchantWalletsListFilters)
	subGroup.POST("/get/:id", controller.MerchantWalletsGet)
	subGroup.POST("/action/withdraw", controller.MerchantWalletsActionWithdraw)
}

func merchantMerchantWithdrawsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/withdraws")

	subGroup.POST("/list", controller.MerchantWithdrawsList)
	subGroup.POST("/list/filters", controller.MerchantWithdrawsListFilters)
}

// PUBLIC API

func merchantPublicApiV1Routes(group *gin.RouterGroup) {
	group.POST("/payment-create", controller.V1PaymentCreate)
}
