package app

import (
	"compay/common/config"
	"compay/core/eslog"
	"context"
	"errors"
	"net/http"
)

var apiServer *http.Server

func StartApi() {
	router := generateRouter("api")

	apiServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			eslog.LogError(err)
		}
	}()
	eslog.Log("Api Server started on port:", config.AppPort)
}

func ShutdownApi(ctx context.Context) {
	err := apiServer.Shutdown(ctx)
	if err != nil {
		eslog.LogError("Server forced to shutdown:", err)
	} else {
		eslog.Log("Api Server stopped")
	}
}
