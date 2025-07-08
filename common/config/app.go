package config

import "time"

var AppPort string = ""

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

	// DB
	"DB_HOST",
	"DB_NAME",
	"DB_USER",
	"DB_PORT",
	"DB_PASSWORD",
}
