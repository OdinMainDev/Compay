package efw

import (
	"efw/core/efwlog"
)

func (that *eframework) initLogger() error {
	that.initBasicsLoggers()

	if that.initMethods.InitLoggers != nil {
		that.initMethods.InitLoggers()
	}

	if err := efwlog.Init(); err != nil {
		return err
	}

	efwlog.Log("[EFW: LOGGER] InitializedSuccess")
	return nil
}

func (that *eframework) initBasicsLoggers() {
	efwlog.NewEfwLog("panic").SaveLogger()
	efwlog.NewEfwLog("info").SaveLogger()
	efwlog.NewEfwLog("error").SaveLogger()
	efwlog.NewEfwLog("debug").SaveLogger()
	efwlog.NewEfwLog("request").SaveLogger()
	efwlog.NewEfwLog("response").SaveLogger()
	efwlog.NewEfwLog("efw_request").SaveLogger()
	efwlog.NewEfwLog("efw_proxy_request").SaveLogger()
	efwlog.NewEfwLog("gorm").SaveLogger()
	efwlog.NewEfwLog("gorm_error").SaveLogger()
	efwlog.NewEfwLog("gorm_debug").SaveLogger()
	efwlog.NewEfwLog("gorm_trace").SaveLogger()
	efwlog.NewEfwLog("gorm_info").SaveLogger()
	efwlog.NewEfwLog("gorm_warn").SaveLogger()
	efwlog.NewEfwLog("gin").SaveLogger()
	efwlog.NewEfwLog("gin_error").SaveLogger()
	efwlog.NewEfwLog("notifications_telegram_error").SaveLogger()
}
