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
	merchantAdminMerchantsRoutes(adminRoutes)
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

	subGroup.POST("/list", controller.AdminUsersList)
	subGroup.POST("/list/filters", controller.AdminUsersListFilters)
	subGroup.POST("/get/:id", controller.AdminUsersGet)
	subGroup.POST("/create", controller.AdminUsersCreate)
	subGroup.POST("/process", controller.AdminUsersProcess)
	subGroup.POST("/delete", controller.AdminUsersDelete)
}

func merchantAdminMerchantsRoutes(group *gin.RouterGroup) {
	subGroup := group.Group("/dev/merchants")

	// subGroup.POST("/list", controller.AdminUsersList)
	// subGroup.POST("/list/filters", controller.AdminUsersListFilters)
	// subGroup.POST("/get/:id", controller.AdminUsersGet)
	// subGroup.POST("/create", controller.AdminUsersCreate)
	// subGroup.POST("/process", controller.AdminUsersProcess)
	// subGroup.POST("/delete", controller.AdminUsersDelete)
	subGroup.POST("/search", controller.AdminMerchantsSearch)
}
