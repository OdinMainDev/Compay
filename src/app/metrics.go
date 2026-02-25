package app

import (
	"context"
	"efw/core/efwgin"
	"efw/core/efwlog"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var metricsServer *http.Server

func StartMetrics() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	metricsServer = &http.Server{Addr: ":10666", Handler: router.Handler()}
	go func() {
		if err := metricsServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			efwlog.LogError("Metrics server started with error: ", err.Error())
		}
	}()
}

func StopMetrics() {
	err := metricsServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Metrics server ended with error:", err.Error())
	}
}
