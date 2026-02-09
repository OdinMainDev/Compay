package efw

import "efw/core/efwnotify"

func (that *eframework) initProfiler() error {
	if that.options.ProfilerDisabled {
		return nil
	}

	if !that.options.NotifyTelegramDisabled {
		if err := efwnotify.InitTelegram(); err != nil {
			return err
		}
	}

	return nil
}
