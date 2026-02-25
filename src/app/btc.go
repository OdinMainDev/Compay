package app

import (
	"context"
	"efw/core/efwgin"
	"efw/core/efwlog"
	"errors"
	"net/http"

	"compay/common/config"
	"compay/src/helpers"
	controller "compay/src/http/controllers/btc"
	"compay/src/services/cryptocurrencies/btc"
)

var btcServer *http.Server

func StartBtc() {
	err := btc.Start()
	if err != nil {
		helpers.NotificationError("FATAL ERROR ON STARTUP BTC NODE: " + err.Error())
		panic("FATAL ERROR ON STARTUP BTC NODE: " + err.Error())
	}
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.POST("/generate-wallet", controller.GenerateWallet)
	router.POST("/get-transactions", controller.GetTransactions)
	router.POST("/get-block-count", controller.GetBlockCount)
	router.POST("/withdraw", controller.Withdraw)

	btcServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := btcServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			efwlog.LogError("Btc server started with error: ", err.Error())
		}
	}()
}

func StopBtc() {
	err := btcServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("Btc server ended with error:", err.Error())
	}
}
