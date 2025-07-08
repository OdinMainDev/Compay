package app

import (
	controller "compay/application/http/controllers/wallet"
	"compay/application/http/middleware"
	"compay/application/service/start"
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var walletServer *http.Server

func StartWallet() {
	start.FillCurrencies()

	router := generateRouter("wallet")

	apiGroup := router.Group("/v1")

	walletApiRoutes(apiGroup)
	walletAuthRoutes(apiGroup)

	privateGroup := apiGroup.Group("/private")
	privateGroup.Use(middleware.CheckAuthToken())
	walletWalletRoutes(privateGroup)

	walletServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := walletServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Wallet Server started on port:", config.AppPort)
}

func walletApiRoutes(group *gin.RouterGroup) {
	group.POST("/get/currencies", controller.GetCurrencies)
}

func walletAuthRoutes(group *gin.RouterGroup) {
	authGroup := group.Group("/auth")

	authGroup.POST("/register", controller.AuthRegister)

	authGroup.POST("/check/user", middleware.CheckUserToken(), controller.AuthCheckUser)
	authGroup.POST("/check/token", middleware.CheckAuthToken(), controller.AuthCheckToken)
	authGroup.POST("/check/pin", middleware.CheckUserToken(), controller.AuthCheckPinCode)
}

func walletWalletRoutes(group *gin.RouterGroup) {
	walletGroup := group.Group("/wallet")

	walletGroup.POST("/new", controller.AddNewWallet)
}

func ShutdownWallet(ctx context.Context) {
	err := walletServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Wallet Server forced to shutdown:", err)
	} else {
		eslog.Log("Wallet Server stopped")
	}
}
