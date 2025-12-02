package efwnotify

import (
	"bytes"
	"efw/config"
	"encoding/json"
	"fmt"
)

var ErrorTelegramInvalidChatId = fmt.Errorf("%s", "[EFW: NOTIFY] InvalidChatId")

var ErrorTelegramEncodeBodySendMessage = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramEncodeBodySendMessage:", err)
}

var ErrorTelegramSendMessageRequest = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramSendMessage:", err)
}

var ErrorTelegramDecodeResponseSendMessage = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramDecodeResponseSendMessage:", err)
}

var ErrorTelegramSendMessageResponse = func(err string) error {
	return fmt.Errorf("%s %s", "[EFW: NOTIFY] TelegramSendMessage:", err)
}

func NotifyTelegram(message string, args ...int64) error {
	chatId := config.NotifyTelegramChatId
	if len(args) == 1 && args[0] != 0 {
		chatId = args[0]
	}

	if chatId == 0 {
		return ErrorTelegramInvalidChatId
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramClient.token)
	body := map[string]any{
		"chat_id":    chatId,
		"text":       message,
		"parse_mode": "HTML",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return ErrorTelegramEncodeBodySendMessage(err.Error())
	}

	resp, err := telegramClient.client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return ErrorTelegramSendMessageRequest(err.Error())
	}

	defer resp.Body.Close()

	var result struct {
		Ok    bool   `json:"ok"`
		Error string `json:"description"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ErrorTelegramDecodeResponseSendMessage(err.Error())
	}

	if !result.Ok {
		return ErrorTelegramSendMessageResponse(result.Error)
	}
	return nil
}
