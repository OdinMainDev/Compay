package efw

import (
	"efw/core/efwgin/validator"
	"efw/core/efwlog"
	"efw/core/efwmaxmind"

	"github.com/gin-gonic/gin"
)

// TODO: раскоментировать config.AppDebug
func (that *eframework) initGin() error {
	// if !config.AppDebug {
	gin.DefaultWriter = efwlog.NewEfwGinLogWriter()
	// }

	validator.Init(that.initMethods.InitValidation)
	efwmaxmind.CreateInstance()

	efwlog.Log("[EFW: GIN] InitializedSuccess")
	return nil
}
