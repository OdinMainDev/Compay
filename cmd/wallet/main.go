package main

import (
	"compay/application/app"
	"compay/common/database"
	"context"
)

func main() {
	database.SetMigrationsType(true)
	app.CreateAppInstance("wallet", []func(){app.StartWallet}, []func(context.Context){app.ShutdownWallet})
}
