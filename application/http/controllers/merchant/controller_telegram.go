package merchant

import (
	"compay/application/service/tgbot"
	"compay/common/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot/models"
)

func ResponseTelegram(context *gin.Context) {
	if context.GetHeader("X-Telegram-Bot-Api-Secret-Token") != config.AppTelegramSecretToken {
		context.Status(http.StatusOK)
		return
	}

	var update models.Update
	if err := context.ShouldBindJSON(&update); err != nil {
		context.Status(http.StatusOK)
		return
	}

	context.Status(http.StatusOK)

	go func(context *gin.Context, update *models.Update) {
		bot, err := tgbot.CreateInstance(update)
		if err != nil {
			return
		}
		if _, err := bot.ValidateUser(); err != nil {
			return
		}
		bot.Response()
	}(context.Copy(), &update)
}
