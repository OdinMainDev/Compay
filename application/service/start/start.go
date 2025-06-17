package start

import "compay/application/service/esvalidator"

func StartupApplication() []func() {
	return []func(){
		fillAppConfig,
		fillSubServicesConfig,
		fillTelegramBotConfig,
		fillCryptoConfig,
	}
}

func StartupServices() {
	esvalidator.Initialize()
	registerCachableTypes()
}
