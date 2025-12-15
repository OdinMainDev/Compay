package app

import (
	"compay/common/config"
	"context"
	"efw/core/efwgin"
	efwmiddleware "efw/core/efwgin/efw_middleware"
	"efw/core/efwlog"
	"net/http"

	controller "compay/src/http/controllers/merchant"

	"github.com/gin-gonic/gin"
)

var merchantServer *http.Server

func StartMerchant() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{
		AdditionalHeaders: []string{
			"T-Totp",
		},
	})

	apiRoutes := router.Group("/api")
	// apiRoutes.Use(middleware.SaveRequestBodyToContext())
	// apiRoutes.Use(middleware.CheckTokenAndCasino())
	// apiRoutes.Use(middleware.CheckUserCountry(true))

	/* +++++++++++++++ PANEL ROUTES +++++++++++++++ */
	// Groups
	panelRoutes := apiRoutes.Group("/panel")
	authRoutes := panelRoutes.Group("/auth")
	// publicRoutes := apiRoutes.Group("/public")
	// privateRoutes := apiRoutes.Group("/private")
	// privateRoutes.Use(middleware.CheckAuth())

	// Routes
	merchantAuthRoutes(authRoutes)

	/* --------------- PANEL ROUTES --------------- */

	// v1Routes := apiRoutes.Group("/v1")

	merchantServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := merchantServer.ListenAndServe(); err != nil {
			efwlog.LogError("Merchant server started with error:", err.Error())
		}
	}()
}

func StopMerchant() {
	err := merchantServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Merchant server ended with error:", err.Error())
	}
}

func merchantAuthRoutes(group *gin.RouterGroup) {
	group.POST("/login", efwmiddleware.CheckThrottling(5), controller.AuthLogin)
	// group.POST("/login", efwmiddleware.CheckThrottling(5), controller.AuthLogin)
	// group.POST("/totp", efwmiddleware.CheckThrottling(5), middleware.CheckAuth(true), controller.AuthTotp)
}
