package tgbot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func (tgBot *EsTgBot) handleCallbackQuery() {
	// Запрещаем в чатах использовать бота
	if tgBot.CallbackQuery.From.ID <= 0 {
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	tgBot.TgBot.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    tgBot.CallbackQuery.From.ID,
		MessageID: tgBot.CallbackQuery.Message.Message.ID,
	})

	switch {
	case tgBot.CallbackQuery.Data == "auth_login":
		tgBot.sendAuthWebMessage()
		return
	}
}
