package efw

import (
	"efw/config"
	"efw/core/efwlog"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var ErrorEnvLoadingFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: CONFIG] LoadingError:", err)
}

var ErrorEnvVariableMissing = func(v string) error {
	return fmt.Errorf("%s %s", "[EFW: CONFIG] BadEnvVariable:", v)
}

func (that *eframework) loadEnvironmentVariables() error {
	// load .env file (default name - .env)
	err := godotenv.Load()
	if err != nil {
		return ErrorEnvLoadingFail(err.Error())
	}

	envCheckKeys := []string{
		"APP_SECRET",
		"APP_DEBUG",
	}

	if len(that.options.AppEnvVariables) > 0 {
		envCheckKeys = append(envCheckKeys, that.options.AppEnvVariables...)
	}

	that.envDatabaseChecks(&envCheckKeys)
	that.envClamavChecks(&envCheckKeys)
	that.envRedisChecks(&envCheckKeys)
	that.envRabbitMQChecks(&envCheckKeys)
	that.envNotifyChecks(&envCheckKeys)
	that.envProfilerChecks(&envCheckKeys)

	for _, value := range envCheckKeys {
		if os.Getenv(value) == "" {
			return ErrorEnvVariableMissing(value)
		}
	}

	that.fillConfigs()

	for _, callback := range that.options.AppEnvCallbacks {
		callback()
	}

	efwlog.Log("[EFW: CONFIG] InitializedSuccess")
	return nil
}

// Validate Functions

func (that *eframework) envDatabaseChecks(envCheckKeys *[]string) {
	if !that.options.DatabaseDisabled {
		*envCheckKeys = append(*envCheckKeys,
			"DB_HOST",
			"DB_NAME",
			"DB_USER",
			"DB_PORT",
			"DB_PASSWORD",
		)
	}
}

func (that *eframework) envClamavChecks(envCheckKeys *[]string) {
	if !that.options.ClamavDisabled {
		*envCheckKeys = append(*envCheckKeys,
			"CLAMAV_URL",
			"CLAMAV_SKIP_VALIDATION",
		)
	}
}

func (that *eframework) envRedisChecks(envCheckKeys *[]string) {
	if !that.options.RedisCacheDisabled {
		*envCheckKeys = append(*envCheckKeys,
			"REDIS_HOST",
		)
	}
}

func (that *eframework) envRabbitMQChecks(envCheckKeys *[]string) {
	if !that.options.RabbitMQDisabled {
		*envCheckKeys = append(*envCheckKeys,
			"RABBITMQ_USERNAME",
			"RABBITMQ_PASSWORD",
			"RABBITMQ_HOST",
			"RABBITMQ_PORT",
		)
	}
}

func (that *eframework) envNotifyChecks(envCheckKeys *[]string) {
	if that.options.NotifyDisabled {
		return
	}

	if !that.options.NotifyTelegramDisabled {
		*envCheckKeys = append(*envCheckKeys,
			"NOTIFY_TELEGRAM_TOKEN",
		)
	}
}

func (that *eframework) envProfilerChecks(envCheckKeys *[]string) {
	if that.options.ProfilerDisabled {
		return
	}

	*envCheckKeys = append(*envCheckKeys,
		"PROFILER_TELEMETRY_COLLECTOR_HOST",
	)
}

// Fill Functions

func (that *eframework) fillConfigs() {
	that.fillConfigApp()
	that.fillConfigDatabase()
	that.fillConfigCache()
	that.fillConfigCrypt()
	that.fillConfigNotify()
	that.fillConfigProfiler()
}

func (that *eframework) fillConfigApp() {
	config.AppDebug = os.Getenv("APP_DEBUG") == "true"
	config.AppSecret = []byte(os.Getenv("APP_SECRET"))
}

func (that *eframework) fillConfigDatabase() {
	config.DbHost = os.Getenv("DB_HOST")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbName = os.Getenv("DB_NAME")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	if os.Getenv("DB_SSL_MODE") != "" {
		config.DbSSLMode = os.Getenv("DB_SSL_MODE")
	}
	if os.Getenv("DB_TIMEZONE") != "" {
		config.DbTimezone = os.Getenv("DB_TIMEZONE")
	}
}

func (that *eframework) fillConfigCrypt() {
	if os.Getenv("CRYPT_GENERATE_KEYS") == "true" {

	}
}

func (that *eframework) fillConfigCache() {
	config.RedisHost = os.Getenv("REDIS_HOST")
	config.RedisUsername = os.Getenv("REDIS_USERNAME")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")

	if value := os.Getenv("CACHE_SHARDS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			config.CacheShardCount = parsed
		}
	}

	if value := os.Getenv("CACHE_QUEUE_SIZE"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			config.CacheOperationQueueSize = parsed
		}
	}

	if value := os.Getenv("CACHE_OPERATION_TIMEOUT_SECONDS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			config.CacheOperationTimeout = time.Duration(parsed) * time.Second
		}
	}
}

func (that *eframework) fillConfigNotify() {
	config.NotifyTelegramToken = os.Getenv("NOTIFY_TELEGRAM_TOKEN")

	if value := os.Getenv("NOTIFY_TELEGRAM_CHAT_ID"); value != "" {
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			config.NotifyTelegramChatId = parsed
		}
	} else {
		fmt.Println("NOTIFY_TELEGRAM_CHAT_ID not set in environment file")
		efwlog.Log("NOTIFY_TELEGRAM_CHAT_ID not set in environment file")
	}
}

func (that *eframework) fillConfigProfiler() {
	config.ProfilerTelemetryCollectorHost = os.Getenv("PROFILER_TELEMETRY_COLLECTOR_HOST")
}
