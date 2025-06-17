package main

import (
	"compay/application/app"
	"context"
)

func main() {
	app.CreateAppInstance([]func(){app.StartApi}, []func(context.Context){app.ShutdownApi})
}
