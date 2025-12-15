package config

import "efw/core/efwlog"

func InitLoggers() {
	efwlog.NewEfwLog("panic").SetEnableConsoleOutput(true).SaveLogger()
	efwlog.NewEfwLog("gin_error").SetEnableConsoleOutput(true).SetDisableCallerInfo(true).SaveLogger()
}
