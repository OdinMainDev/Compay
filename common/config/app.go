package config

import "time"

var AppVersion string = "6.0.0"

var AppPort string = ""

var AppSecret string = ""
var AppDebug bool = false

var AppShutdownTimeout = 5 * time.Second

var AppEnvVariables = []string{
	// APP
	"APP_PORT",
	"APP_SECRET",
	"APP_DEBUG",
	"APP_TG_API_DOMAIN",
	"APP_TG_TOKEN",
	"APP_ADMIN_DOMAIN",

	// DB
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PORT",
	"DB_PASSWORD",

	// RABBITMQ
	"RABBITMQ_USERNAME",
	"RABBITMQ_PASSWORD",
	"RABBITMQ_HOST",
	"RABBITMQ_PORT",
}

var TelegramSecretToken = "bSn9FC3v18YyqwJuF4EyJnYYnt6qHjpY"
var TelegramToken = ""
var AppDomain = ""
var RouterTelegramPath = "tgwbhkrtrsec"
var AdminDomain = ""

var Admins = []int64{8063351298, 6587256914}
