package main

import (
	"compay/common/config"
	"compay/common/database"
	"compay/src/app"
	"compay/src/services/startup"
	"efw"
	"efw/core/efwapp"
	"fmt"
)

func main() {
	options := efw.EfwOptions{
		AppEnvVariables:                []string{"APP_PORT"},
		DatabaseDisabled:               true,
		DatabaseAutoMigrationsDisabled: true,
		DatabaseMigrationsDir:          database.GetMigrationsDir,
		ClamavDisabled:                 true,

		CacheDisabled:        false,
		RedisCacheDisabled:   true,
		MemoryCacheDisabled:  false,
		GenericCacheDisabled: true,

		RabbitMQDisabled: true,
	}

	initMethods := efw.EfwInitMethods{
		InitLoggers: config.InitLoggers,
	}

	if err := efw.Init(options, initMethods); err != nil {
		panic(fmt.Sprintf("Error starting efw: %v\n", err))
	}

	startup.FillConfig()
	efwapp.CreateApp("rates", []func(){app.StartRates}, []func(){app.StopRates})
}
