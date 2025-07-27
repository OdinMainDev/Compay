package app

import (
	controller "compay/application/http/controllers/merchant"
	"compay/application/http/middleware"
	"compay/application/service/tgbot"
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var merchantServer *http.Server

func StartMerchant() {
	router := generateRouter("merchant")

	router.POST("/"+config.AppRouterTelegramPath, controller.ResponseTelegram)

	authGroup := router.Group("/auth", middleware.CheckThrottling(5))
	merchantAuthRoutes(authGroup)

	privateGroup := router.Group("/private")
	privateGroup.Use(middleware.CheckMerchantAuth())

	merchantProfileRoutes(privateGroup)

	adminGroup := privateGroup.Group("/admin")
	adminGroup.Use(middleware.CheckMerchantIsAdmin())

	merchantAdminCurrencyRoutes(adminGroup)
	merchantAdminMerchantRoutes(adminGroup)
	merchantAdminUserRoutes(adminGroup)

	// apiGroup := router.Group("/v1")

	tgbot.SetWebHook()
	merchantServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := merchantServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Merchant Server started on port:", config.AppPort)
}

func merchantAuthRoutes(group *gin.RouterGroup) {
	group.POST("/login", middleware.CheckThrottling(5), controller.AuthLogin)
	group.POST("/totp", middleware.CheckThrottling(5), middleware.CheckMerchantAuth(true), controller.AuthTotp)
}

func merchantProfileRoutes(group *gin.RouterGroup) {
	group.GET("/get/me", controller.GetMe)
	group.GET("/get/my/has-2-fa", controller.GetMyHas2Fa)
	group.POST("/get/my/2-fa-uri", controller.GetMy2FaUri)
	group.POST("/set/my/2-fa", controller.SetMy2Fa)
}

func merchantAdminCurrencyRoutes(group *gin.RouterGroup) {
	currenciesGroup := group.Group("/currencies")

	currenciesGroup.GET("/list", controller.CurrenciesList)
	currenciesGroup.GET("/list/filters", controller.CurrenciesListFilters)
	currenciesGroup.GET("/get/:id", controller.CurrencyGet)
	currenciesGroup.POST("/process", controller.CurrencyProcess)
	currenciesGroup.POST("/delete/:id", controller.CurrencyDelete)
	currenciesGroup.POST("/import", controller.CurrencyImport)

	currenciesGroup.GET("/get", controller.GetCurrencies)
}

func merchantAdminMerchantRoutes(group *gin.RouterGroup) {
	merchantsGroup := group.Group("/merchants")

	merchantsGroup.GET("/list", controller.MerchantsList)
	merchantsGroup.GET("/list/filters", controller.MerchantsListFilters)
	merchantsGroup.GET("/get/:id", controller.MerchantGet)
	merchantsGroup.POST("/process", controller.MerchantProcess)

	merchantsGroup.GET("/get", controller.GetMerchants)
}

func merchantAdminUserRoutes(group *gin.RouterGroup) {
	usersGroup := group.Group("/users")

	usersGroup.GET("/list", controller.UsersList)
	usersGroup.GET("/list/filters", controller.UsersListFilters)
	usersGroup.GET("/get/:id", controller.UserGet)
	usersGroup.POST("/process", controller.UserProcess)

}

func ShutdownMerchant(ctx context.Context) {
	err := merchantServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Merchant Server forced to shutdown:", err)
	} else {
		eslog.Log("Merchant Server stopped")
	}
}
