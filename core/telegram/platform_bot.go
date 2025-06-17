package telegram

import (
	"compay/core/config"
	"compay/core/eslog"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

var platformBot *gotgbot.Bot

func PlatformBotInit() {
	if !config.TelegramNotifyEnable {
		return
	}

	var err error
	platformBot, err = gotgbot.NewBot(config.PlatformBotToken, &gotgbot.BotOpts{
		BotClient: &gotgbot.BaseBotClient{
			Client: http.Client{},
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout, // Customise the default request timeout here
				APIURL:  gotgbot.DefaultAPIURL,  // As well as the Default API URL here (in case of using local bot API servers)
			},
		},
	})

	if err != nil {
		eslog.PanicLog(err.Error())
		return
	}
}

func SendPlatformMessage(chatId int64, message string) bool {
	if platformBot == nil {
		eslog.Log("restart platform bot")
		PlatformBotInit()
		if platformBot == nil {
			eslog.LogError("cant start platform bot")
			return false
		}
	}
	_, err := platformBot.SendMessage(chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		eslog.LogError(err.Error())
		return false
	}

	return true
}
