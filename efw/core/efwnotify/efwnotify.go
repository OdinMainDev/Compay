package efwnotify

type Options struct {
	TelegramChatId int64
}

func Notify(message string, opt Options) error {
	if telegramClient != nil {
		if err := NotifyTelegram(message, opt.TelegramChatId); err != nil {
			return err
		}
	}

	return nil
}
