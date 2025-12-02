package config

import "efw/core/efwlog"

func InitLoggers() {
	efwlog.NewEfwLog("panic").SetEnableConsoleOutput(true).SaveLogger()
	efwlog.NewEfwLog("info").SaveLogger()
	efwlog.NewEfwLog("error").SaveLogger()
	efwlog.NewEfwLog("debug").SaveLogger()
	efwlog.NewEfwLog("request").SaveLogger()
	efwlog.NewEfwLog("response").SaveLogger()
	efwlog.NewEfwLog("game_request").SaveLogger()
	efwlog.NewEfwLog("efw_request").SaveLogger()
	efwlog.NewEfwLog("efw_proxy_request").SaveLogger()
	efwlog.NewEfwLog("chat").SaveLogger()
	efwlog.NewEfwLog("gorm").SaveLogger()
	efwlog.NewEfwLog("gorm_error").SaveLogger()
	efwlog.NewEfwLog("gorm_debug").SaveLogger()
	efwlog.NewEfwLog("gorm_trace").SaveLogger()
	efwlog.NewEfwLog("gorm_info").SaveLogger()
	efwlog.NewEfwLog("gorm_warn").SaveLogger()
	efwlog.NewEfwLog("gin").SaveLogger()
}
