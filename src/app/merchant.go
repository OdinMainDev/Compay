package app

import (
	"compay/common/config"
	"context"
	"efw/core/efwgin"
	efwmiddleware "efw/core/efwgin/efw_middleware"
	"efw/core/efwlog"
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
	merchantAdminCurrenciesRoutes(adminRoutes)

	merchantAdminApiRoutes(adminRoutes)

	// LOGS
	logGroup := adminRoutes.Group("/dev/logs")
	merchantAdminLogAuthRoutes(logGroup)
	merchantAdminLogUserRoutes(logGroup)
	/* --------------- ADMIN ROUTES --------------- */

	/* --------------- PANEL ROUTES --------------- */

	/* +++++++++++++++ MERCHANT ROUTES +++++++++++++++ */
	merchantRoutes := panelRoutes.Group("/merchant")
	merchantRoutes.Use(middleware.CheckMerchant())

	merchantMerchantRoutes(merchantRoutes)

	/* --------------- MERCHANT ROUTES --------------- */

	/* +++++++++++++++ V1 ROUTES +++++++++++++++ */
	v1 := apiRoutes.Group("/v1")

	merchantPublicApiV1Routes(v1)
	/* --------------- V1 ROUTES --------------- */
	merchantServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := merchantServer.ListenAndServe(); err != nil {
			efwlog.LogError("Merchant server started with error:", err.Error())
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
	group.POST("/totp", efwmiddleware.RateLimit(5), efwmiddleware.CheckAuth(), controller.AuthTotp)
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
	// subGroup.POST("/action/disable", controller.AdminMerchantsActionDisable)
	// subGroup.POST("/action/enable", controller.AdminMerchantsActionEnable)
}

func merchantAdminApiRoutes(group *gin.RouterGroup) {
	getGroup := group.Group("/get/data")
	actionGroup := group.Group("/action")

	getGroup.POST("/fee-types", controller.AdminApiFeeTypes)

	actionGroup.POST("/clear-cache", controller.AdminApiClearCache)
}

func merchantAdminCurrenciesRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/currencies")

	subGroup.POST("/list", controller.AdminCurrencyList)
	subGroup.POST("/list/filters", controller.AdminCurrencyListFilters)
	subGroup.POST("/get/:id", controller.AdminCurrencyGet)
	subGroup.POST("/process", controller.AdminCurrencyProcess)
	subGroup.POST("/search", controller.AdminCurrencySearch)
	subGroup.POST("/action/import", controller.AdminCurrencyActionImport)
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

// MERCHANT ROUTES

func merchantMerchantRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/merchant")

	subGroup.POST("/get", controller.MerchantSettingsGet)
	subGroup.POST("/process", controller.MerchantSettingsProcess)
	subGroup.POST("/action/get-token", controller.MerchantSettingsActionGetToken)
	subGroup.POST("/action/generate-secret", controller.MerchantSettingsActionGenerateSecret)
}

// PUBLIC API

func merchantPublicApiV1Routes(group *gin.RouterGroup) {
	group.POST("/payment-create", controller.V1PaymentCreate)
}
