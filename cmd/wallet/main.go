package main

import (
	"compay/application/app"
	"context"
)

func main() {
	app.CreateAppInstance("wallet", []func(){app.StartWallet}, []func(context.Context){app.ShutdownWallet})
}
