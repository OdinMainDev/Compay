package efw

import (
	"efw/core/efwnotify"
)

func (that *eframework) initNotify() error {
	if that.options.NotifyDisabled {
		return nil
	}

	if !that.options.NotifyTelegramDisabled {
		if err := efwnotify.InitTelegram(); err != nil {
			return err
		}
	}

	return nil
}
