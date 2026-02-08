package app

import (
	"context"
	"efw/core/efwgin"
	"efw/core/efwlog"
	"net/http"

	"compay/common/config"
	controller "compay/src/http/controllers/helloworld"
)

var helloworldServer *http.Server

func StartHelloWorld() {
	router := efwgin.CreateRouter(efwgin.RouterOptions{})

	router.GET("/hello", controller.Hello)
	router.GET("/hello2", controller.Hello2)

	helloworldServer = &http.Server{Addr: ":" + config.AppPort, Handler: router.Handler()}
	go func() {
		if err := helloworldServer.ListenAndServe(); err != nil {
			efwlog.LogError("HelloWorld server started with error: ", err.Error())
		}
	}()
}

func StopHelloWorld() {
	err := helloworldServer.Shutdown(context.Background())
	if err != nil {
		efwlog.LogError("HelloWorld server ended with error:", err.Error())
	}
}
