package app

import (
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"
)

var apiServer *http.Server

func StartApi() {
	router := generateRouter("api")

	// router.POST("/"+config.RouterTelegramPath, controller.ResponseTelegram)

	// router.POST("/auth/login", middleware.CheckThrottling(5), controller.AuthLogin)
	// router.POST("/auth/totp", middleware.CheckThrottling(5), middleware.CheckAuth(true), controller.AuthTotp)

	// group := router.Group("/api")
	// group.Use(middleware.CheckAuth())

	// group.GET("/get/me", controller.GetMe)
	// group.GET("/get/my/has-2-fa", controller.GetMyHas2Fa)
	// group.POST("/get/my/2-fa-uri", controller.GetMy2FaUri)
	// group.POST("/set/my/2-fa", controller.SetMy2Fa)

	// group.GET("/users/list", controller.UserList)
	// group.GET("/users/list/filters", controller.UserListFilters)
	// group.GET("/users/get/:id", controller.UserGet)
	// group.POST("/users/process", controller.UserProcess)

	// group.GET("/wallets/list", controller.WalletList)
	// group.GET("/wallets/list/filters", controller.WalletListFilters)
	// group.GET("/wallets/get/:id", controller.WalletGet)
	// group.POST("/wallets/create", controller.WalletProcess)

	apiServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Api Server started on port:", config.AppPort)
}

func ShutdownApi(ctx context.Context) {
	err := apiServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Server forced to shutdown:", err)
	} else {
		eslog.Log("Api Server stopped")
	}
}
