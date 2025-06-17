package app

import (
	"compay/application/service/start"
	"compay/common/config"
	"compay/core"
	"compay/core/eslog"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func CreateAppInstance(startApps []func(), shutdownApps []func(context.Context)) {
	core.CreateInstance(core.InitData{
		EnvCheckKeys: config.AppEnvVariables,
		EnvCallbacks: start.StartupApplication(),
	})
	start.StartupServices()

	for _, start := range startApps {
		go start()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	eslog.Log("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), config.AppShutdownTimeout)
	defer cancel()

	var wg sync.WaitGroup

	for _, shutdown := range shutdownApps {
		wg.Add(1)
		shutdown := shutdown
		go func() {
			defer wg.Done()
			shutdown(ctx)
		}()
	}
	wg.Wait()
	eslog.Log("Server exited.")
}

func generateRouter(service string) *gin.Engine {
	gin.DefaultWriter = eslog.NewEsGinWriter()
	router := gin.New()
	router.TrustedPlatform = "X-Forwarded-For"

	// базовые мидлвейры
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Forwarded-For")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "C-Token")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Telegram-Bot-Api-Secret-Token")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Referer")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Origin")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "C-Language")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "C-Totp")
	router.Use(cors.New(corsConfig))

	// trust proxies
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// profiler initialize
	router.Use(otelgin.Middleware(service))

	return router
}
