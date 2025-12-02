package efw

import (
	"efw/core/efwlog"
)

func (that *eframework) initLogger() error {
	if err := efwlog.Init(); err != nil {
		return err
	}

	efwlog.Log("[EFW: LOGGER] InitializedSuccess")
	return nil
}
