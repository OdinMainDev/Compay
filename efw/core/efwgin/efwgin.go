package efwgin

import (
	"efw/config"
	efwmiddleware "efw/core/efwgin/efw_middleware"
	"efw/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	AllowMethods      []string
	AdditionalHeaders []string
}

func CreateRouter(routerOptions RouterOptions) *gin.Engine {
	if !config.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Обозначаем заголовок отвечающий за ip адрес
	router.TrustedPlatform = "X-Forwarded-For"

	// Базовые мидлвейры
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Хидеры безопасности
	router.Use(efwmiddleware.SecurityHeaders())

	// Доверенные прокси
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	if len(routerOptions.AllowMethods) == 0 {
		routerOptions.AllowMethods = []string{"GET", "POST"}
	}

	allowHeaders := []string{
		"Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"X-Forwarded-For",
		"X-Telegram-Bot-Api-Secret-Token",
		"Referer",
	}
	if len(routerOptions.AdditionalHeaders) == 0 {
		allowHeaders = append(allowHeaders, routerOptions.AdditionalHeaders...)
		allowHeaders = utils.RemoveDuplicates(allowHeaders)
	}

	corsConfig := cors.Config{
		AllowMethods:     routerOptions.AllowMethods,
		AllowHeaders:     allowHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}

	router.Use(cors.New(corsConfig))
	return router
}
