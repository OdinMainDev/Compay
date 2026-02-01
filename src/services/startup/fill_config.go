package startup

import (
	"compay/common/config"
	configEfw "efw/config"
	"os"
)

func FillConfig() {
	config.AppPort = os.Getenv("APP_PORT")
}

func fillConfigMerchant() {
	configEfw.NotifyTelegramToken = "8287820061:AAGxEL9uA0-fnSx_EPJ8exxyAh2EpUrBe-M"
	configEfw.NotifyTelegramChatId = -1003675944523
}
