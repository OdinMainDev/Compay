package app

import (
	"compay/common/config"
	"compay/src/services/ratesparser"
	"context"
	"efw/core/efwgin"
	"efw/core/efwlog"
	"errors"
	"net/http"

	controller "compay/src/http/controllers/rates"
)

var ratesServer *http.Server

func StartRates() {
	go ratesparser.Initialize()
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.GET("/get-rates", controller.GetRates)
	router.GET("/get-rate", controller.GetRate)

	ratesServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := ratesServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			efwlog.LogError("Rates server started with error: ", err.Error())
		}
	}()
}

func StopRates() {
	err := ratesServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Rates server ended with error:", err.Error())
	}
}
