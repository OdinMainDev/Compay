package main

import (
	"compay/common/config"
	"compay/common/database"
	"compay/src/app"
	"compay/src/services/startup"
	"compay/src/services/validator"
	"efw"
	"efw/core/efwapp"
	"fmt"
)

func main() {
	database.SetMigrationsType(false)

	options := efw.EfwOptions{
		AppEnvVariables:       []string{"APP_PORT", "BTC_API_URL"},
		DatabaseMigrationsDir: database.GetMigrationsDir,
		ClamavDisabled:        true,

		GenericCacheDisabled: true,
		RabbitMQDisabled:     true,

		ProfilerDisabled: true,
	}

	initMethods := efw.EfwInitMethods{
		InitLoggers:    config.InitLoggers,
		InitValidation: validator.CustomValidators,
	}

	if err := efw.Init(options, initMethods); err != nil {
		panic(fmt.Sprintf("Error starting efw: %v\n", err))
	}

	startup.FillConfig()
	startup.FillMerchantConfig()

	efwapp.CreateApp(
		"merchant",
		[]func(){app.StartMerchant, app.StartMetrics},
		[]func(){app.StopMerchant, app.StopMetrics},
	)
}
