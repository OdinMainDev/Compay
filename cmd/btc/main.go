package main

import (
	"compay/common/config"
	"compay/src/app"
	"compay/src/services/startup"
	"compay/src/services/validator"
	"efw"
	"efw/core/efwapp"
	"fmt"
)

func main() {

	options := efw.EfwOptions{
		AppEnvVariables:       []string{"APP_PORT", "BTC_RPC_HOST", "BTC_RPC_USERNAME", "BTC_RPC_PASSWORD", "BTC_NETWORK_MAINNET"},
		DatabaseMigrationsDir: nil,
		DatabaseDisabled:      true,
		ClamavDisabled:        true,
		CacheDisabled:         true,
		GenericCacheDisabled:  true,
		RabbitMQDisabled:      true,
	}

	initMethods := efw.EfwInitMethods{
		InitLoggers:    config.InitLoggers,
		InitValidation: validator.CustomValidators,
	}

	if err := efw.Init(options, initMethods); err != nil {
		panic(fmt.Sprintf("Error starting efw: %v\n", err))
	}

	startup.FillConfig()

	efwapp.CreateApp(
		"btc",
		[]func(){app.StartBtc, app.StartMetrics},
		[]func(){app.StopBtc, app.StopMetrics},
	)
}
