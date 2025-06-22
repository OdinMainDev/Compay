package app

import (
	controller "compay/application/http/controllers/rates"
	"compay/application/service/rates"
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"
)

var ratesServer *http.Server

func StartRates() {
	go rates.Initialize()
	router := generateRouter("rates")

	router.GET("/get-rates", controller.GetRates)
	router.GET("/get-rate", controller.GetRate)

	ratesServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := ratesServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Rates Server started on port:", config.AppPort)
}

func ShutdownRates(ctx context.Context) {
	err := ratesServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Server forced to shutdown:", err)
	} else {
		eslog.Log("Rates Server stopped")
	}
}
