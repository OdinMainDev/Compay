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

	// REDIS
	config.RedisEnable = os.Getenv("REDIS_ENABLE") == "true"

	// RABBITMQ
	config.RabbitmqEnable = os.Getenv("RABBITMQ_ENABLE") == "true"

	// METRICS
	config.TelemetryCollectorHost = os.Getenv("TELEMETRY_COLLECTOR_HOST")
}
