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

	"github.com/gin-gonic/gin"
)

var merchantServer *http.Server

func StartMerchant() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	apiRoutes := router.Group("/api")
	// apiRoutes.Use(middleware.SaveRequestBodyToContext())
	// apiRoutes.Use(middleware.CheckUserCountry(true))

	/* +++++++++++++++ AUTH ROUTES +++++++++++++++ */
	authRoutes := apiRoutes.Group("/auth")

	merchantAuthRoutes(authRoutes)
	/* --------------- AUTH ROUTES --------------- */

	/* +++++++++++++++ PANEL ROUTES +++++++++++++++ */
	panelRoutes := apiRoutes.Group("/panel")
	panelRoutes.Use(efwmiddleware.CheckAuth(), middleware.CheckUser(), middleware.CheckAuthTotp())

	merchantPanelUserRoutes(panelRoutes)

	/* +++++++++++++++ ADMIN ROUTES +++++++++++++++ */
	adminRoutes := panelRoutes.Group("/admin")
	adminRoutes.Use(middleware.CheckIsUserAdmin())
	merchantAdminUsersRoutes(adminRoutes)
	/* --------------- ADMIN ROUTES --------------- */

	/* --------------- PANEL ROUTES --------------- */

	/* +++++++++++++++ MERCHANT ROUTES +++++++++++++++ */
	// merchantRoutes := apiRoutes.Group("/merchant")
	// merchantRoutes.Use(efwmiddleware.CheckAuth(), middleware.CheckUser(), middleware.CheckAuthTotp(), middleware.CheckMerchant())

	/* --------------- MERCHANT ROUTES --------------- */

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
	group.POST("/login", efwmiddleware.RateLimit(5), controller.AuthLogin)
	// group.POST("/totp", efwmiddleware.RateLimit(5), efwmiddleware.CheckAuth(true), controller.AuthTotp)
}

func merchantPanelUserRoutes(group *gin.RouterGroup) {
	group.GET("/get/me", controller.UserGetMe)
}

func merchantAdminUsersRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/users")

	subGroup.GET("/list", controller.UserGetMe)
	subGroup.GET("/list/filters", controller.UserGetMe)
	subGroup.GET("/get/:id", controller.UserGetMe)
	subGroup.GET("/process", controller.UserGetMe)
	subGroup.GET("/delete/:id", controller.UserGetMe)
}
