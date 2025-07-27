package tgbot

import (
	"compay/common/config"
	"compay/core/esrequest"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const telegramAPI = "https://api.telegram.org/bot%s/%s"

func SetWebHook() error {
	webHookUrl := constructWebHookUrl()

	type Request struct {
		URL                string `json:"url"`
		MaxConnections     int    `json:"max_connections,omitempty"`
		SecretToken        string `json:"secret_token,omitempty"`
		DropPendingUpdates bool   `json:"drop_pending_updates"`
	}
	webHookReq := Request{
		URL:                webHookUrl,
		MaxConnections:     200,
		SecretToken:        config.AppTelegramSecretToken,
		DropPendingUpdates: true,
	}

	reqBody, err := json.Marshal(webHookReq)
	if err != nil {
		return err
	}

	respBody, _, err := esrequest.Request(esrequest.RequestHelper{
		Url:    fmt.Sprintf(telegramAPI, config.AppTelegramToken, "setWebhook"),
		Method: "POST",
		IsJson: true,
		Body:   reqBody,
	})
	if err != nil {
		return err
	}

	type Response struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description"`
	}
	var resp Response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return errors.New("badResponse")
	}

	return nil
}

func constructWebHookUrl() string {
	path := strings.ReplaceAll(config.AppDomain+"/"+config.AppRouterTelegramPath, "//", "/")
	return "https://" + path
}
