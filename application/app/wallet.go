package app

import (
	"compay/application/service/start"
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"
)

var walletServer *http.Server

func StartWallet() {
	start.FillCurrencies()

	router := generateRouter("wallet")

	walletServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := walletServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Wallet Server started on port:", config.AppPort)
}

func ShutdownWallet(ctx context.Context) {
	err := walletServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Wallet Server forced to shutdown:", err)
	} else {
		eslog.Log("Wallet Server stopped")
	}
}
