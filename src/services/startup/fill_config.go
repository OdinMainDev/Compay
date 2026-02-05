package startup

import (
	"compay/common/config"
	"os"
)

func FillConfig() {
	config.AppPort = os.Getenv("APP_PORT")

	// Bitcoin RPC Configuration
	config.BtcNodeHost = os.Getenv("BTC_RPC_HOST")
	config.BtcNodeUsername = os.Getenv("BTC_RPC_USERNAME")
	config.BtcNodePassword = os.Getenv("BTC_RPC_PASSWORD")
}
