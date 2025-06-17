package app

import (
	"compay/application/service/rates"
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"
)

var ratesServer *http.Server

func StartRates() {
	router := generateRouter("api")

	rates.Initialize()

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
