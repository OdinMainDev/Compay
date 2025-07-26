package tgbot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (tgBot *EsTgBot) sendStartMessage() {
	if tgBot.UserModel == nil || !tgBot.UserModel.IsApproved {
		return
	}

	tgBot.doSendStartMessage(tgBot.UserModel.TgId)
}

func (tgBot *EsTgBot) doSendStartMessage(chatId int64) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Log in", CallbackData: "auth_login"},
			},
			{
				{Text: "Balances", CallbackData: "auth_login"},
			},
		},
	}

	tgBot.TgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        "Main menu",
		ParseMode:   "HTML",
		ReplyMarkup: keyboard,
	})
}
