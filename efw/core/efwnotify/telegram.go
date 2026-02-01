package efwnotify

import (
	"efw/core/efwrequest"
	"fmt"
	"net/http"
)

type TelegramClient struct {
	client *http.Client
}

var telegramClient *TelegramClient

var ErrorTelegramInit = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramInit:", err)
}

func InitTelegram() error {
	client, _, err := efwrequest.NewClient(efwrequest.RequestHelper{})
	if err != nil {
		return ErrorTelegramInit(err.Error())
	}
	telegramClient = &TelegramClient{
		client: client,
	}
	return nil
}
