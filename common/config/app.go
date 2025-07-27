package config

import (
	"time"

	"github.com/google/uuid"
)

var AppPort string = ""
var AppAdminDomain = ""
var AppTelegramToken = ""
var AppDomain = ""
var AppTelegramSecretToken = uuid.NewString()
var AppRouterTelegramPath = "tgwebhook"

var AppSecret string = ""
var AppDebug bool = false

var AppShutdownTimeout = 5 * time.Second

var AppEnvVariablesBasic = []string{
	// APP
	"APP_PORT",
	"APP_SECRET",
}

var AppEnvVariablesAll = []string{
	// APP
	"APP_PORT",
	"APP_SECRET",
	"APP_ADMIN_DOMAIN",
	"APP_DOMAIN",
	"APP_TELEGRAM_TOKEN",

	// DB
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PORT",
	"DB_PASSWORD",
}
