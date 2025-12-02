package startup

import (
	"compay/common/config"
	"os"
)

func FillConfig() {
	config.AppPort = os.Getenv("APP_PORT")
}
