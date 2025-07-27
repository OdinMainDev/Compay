package tgbot

import (
	"compay/application/entities/merchantuser"
	"compay/common/config"
	"compay/core/utils"

	"fmt"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EsTgBot struct {
	TgBot     *bot.Bot
	UserModel *merchantuser.Model

	Message       *models.Message
	CallbackQuery *models.CallbackQuery
	InlineQuery   *models.InlineQuery
	MyChatMember  *models.ChatMemberUpdated
}

func CreateInstance(update *models.Update) (*EsTgBot, error) {
	tgBot, err := bot.New(config.AppTelegramToken)
	if err != nil {
		return nil, err
	}

	esTgBot := &EsTgBot{
		TgBot:         tgBot,
		Message:       update.Message,
		CallbackQuery: update.CallbackQuery,
		InlineQuery:   update.InlineQuery,
		MyChatMember:  update.MyChatMember,
	}

	if esTgBot.getMyId() == 0 {
		return nil, fmt.Errorf("BadConfiguration")
	}

	return esTgBot, nil
}

func (tgBot *EsTgBot) Response() {
	if tgBot.Message != nil {
		tgBot.handleMessage()
		return
	}
	if tgBot.CallbackQuery != nil {
		tgBot.handleCallbackQuery()
		return
	}
	if tgBot.MyChatMember != nil {
		tgBot.handleMyChatMember()
		return
	}
}

func (tgBot *EsTgBot) getMyId() int64 {
	parts := utils.Explode(":", config.AppTelegramToken)
	if len(parts) != 2 {
		return 0
	}

	botID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0
	}

	return botID
}
