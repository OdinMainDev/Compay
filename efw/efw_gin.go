package efw

import (
	"efw/config"
	"efw/core/efwgin/validator"
	"efw/core/efwlog"
	"efw/core/efwmaxmind"

	"github.com/gin-gonic/gin"
)

func (that *eframework) initGin() error {
	if !config.AppDebug {
		gin.DefaultWriter = efwlog.NewEfwGinLogWriter()
	}
	validator.Init(that.initMethods.InitValidation)
	efwmaxmind.CreateInstance()

	efwlog.Log("[EFW: GIN] InitializedSuccess")
	return nil
}
