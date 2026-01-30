package app

import (
	"context"
	"efw/core/efwgin"
	"efw/core/efwlog"
	"net/http"

	controller "compay/src/http/controllers/btc"
)

var btcServer *http.Server

func StartBtc() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.POST("/generate-wallet", controller.GenerateWallet)

	// TODO: потом вернуть на конфиг
	// btcServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	btcServer = &http.Server{Addr: ":8081", Handler: router.Handler()}
	go func() {
		if err := btcServer.ListenAndServe(); err != nil {
			efwlog.LogError("Btc server started with error:", err.Error())
		}
	}()
}

func StopBtc() {
	err := btcServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Btc server ended with error:", err.Error())
	}
}
