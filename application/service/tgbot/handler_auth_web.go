package tgbot

import (
	"compay/common/config"
	"compay/core/utils"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (tgBot *EsTgBot) sendAuthWebMessage() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	hashString, err := utils.EncryptString(fmt.Sprintf("do_auth_user_as_admin_%d_%d", tgBot.UserModel.GetId(), time.Now().Unix()))
	if err != nil {
		tgBot.TgBot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: tgBot.UserModel.TgId,
			Text:   "Something goes bad\n\nSend this message to the developers: " + err.Error(),
		})
		return
	}

	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Get a new login link", CallbackData: "auth_login"},
			},
			{
				{Text: "Balances", CallbackData: "auth_login"},
			},
		},
	}

	urlStr := config.AppAdminDomain + "/auth/" + hashString

	tgBot.TgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      tgBot.UserModel.TgId,
		Text:        urlStr,
		ReplyMarkup: keyboard,
	})
}
