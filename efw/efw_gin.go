package efw

import (
	"efw/core/efwgin/validator"
	"efw/core/efwlog"
)

func (that *eframework) initGin() error {
	validator.Init()

	efwlog.Log("[EFW: GIN] InitializedSuccess")
	return nil
}
