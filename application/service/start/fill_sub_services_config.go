package start

import (
	"compay/core/config"
	"os"
)

func fillSubServicesConfig() {
	// DB SERVICE
	config.DbHost = os.Getenv("DB_HOST")
	config.DbName = os.Getenv("DB_NAME")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbDisabled = os.Getenv("DB_DISABLED") == "true"

	// CDN
	config.CdnUploaderDisabled = true

	// REDIS
	config.RedisEnable = false

	// RABBITMQ
	config.RabbitmqEnable = false

	// METRICS
	config.TelemetryCollectorHost = os.Getenv("TELEMETRY_COLLECTOR_HOST")

	//  TELEGRAM NOTIFY
	config.TelegramNotifyEnable = false
}
