package main

import (
	"compay/application/app"
	"compay/common/database"
	"context"
)

func main() {
	database.SetMigrationsType(false)
	app.CreateAppInstance("merchant", []func(){app.StartMerchant}, []func(context.Context){app.ShutdownMerchant})
}
