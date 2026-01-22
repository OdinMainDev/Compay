package efw

import (
	"efw/core/efwgin/validator"
	"efw/core/efwlog"
	"efw/core/efwmaxmind"
)

func (that *eframework) initGin() error {
	validator.Init(that.initMethods.InitValidation)
	efwmaxmind.CreateInstance()

	efwlog.Log("[EFW: GIN] InitializedSuccess")
	return nil
}
