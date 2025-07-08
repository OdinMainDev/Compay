package main

import (
	"compay/application/app"
	"context"
)

func main() {
	app.CreateAppInstance("rates", []func(){app.StartRates}, []func(context.Context){app.ShutdownRates})
}
