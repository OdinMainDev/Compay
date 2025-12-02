package efwnotify

import (
	"efw/config"
	"efw/core/efwrequest"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramClient struct {
	token  string
	client *http.Client
}

var telegramClient *TelegramClient

var ErrorTelegramInit = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramInit:", err)
}

var ErrorTelegramInitRequestFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramInitRequestFail:", err)
}

var ErrorTelegramInitDecodeFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramInitDecodeFail:", err)
}

var ErrorTelegramInitResponseFail = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramInitResponseFail:", err)
}

func InitTelegram() error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", config.NotifyTelegramToken)

	client, _, err := efwrequest.NewClient(efwrequest.RequestHelper{})
	if err != nil {
		return ErrorTelegramInit(err.Error())
	}

	var resp *http.Response
	for range 3 {
		resp, err = client.Get(url)
		if err == nil {
			break
		}
	}
	if err != nil {
		return ErrorTelegramInitRequestFail(err.Error())
	}

	defer resp.Body.Close()

	var result struct {
		Ok    bool   `json:"ok"`
		Error string `json:"description"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ErrorTelegramInitDecodeFail(err.Error())
	}

	if !result.Ok {
		return ErrorTelegramInitResponseFail(result.Error)
	}

	telegramClient = &TelegramClient{
		token:  config.NotifyTelegramToken,
		client: client,
	}
	return nil
}
