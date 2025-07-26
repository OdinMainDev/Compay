package tgbot

import (
	"strings"
)

func (tgBot *EsTgBot) handleMessage() {
	// Запрещаем в чатах использовать бота
	if tgBot.Message.Chat.ID <= 0 {
		return
	}

	switch {
	case strings.HasPrefix(tgBot.Message.Text, "/start"):
		tgBot.sendStartMessage()
		return
	case tgBot.Message.Text == "/authweb":
		tgBot.sendAuthWebMessage()
		return
	}
}
