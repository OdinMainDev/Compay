package tgbot

import (
	"compay/application/entities/merchantuser"
	"compay/common/constants"
	"compay/core/utils"
	"fmt"

	"github.com/go-telegram/bot/models"
)

func (tgBot *EsTgBot) ValidateUser() (bool, error) {
	tgUser := tgBot.getTgUser()
	if tgUser.ID == 0 {
		return false, fmt.Errorf("BadRequestFromTelegram")
	}

	userModel := merchantuser.GetModel(merchantuser.QueryHelper{
		TelegramId: tgUser.ID,
	})

	return tgBot.setUser(userModel)
}

func (tgBot *EsTgBot) getTgUser() models.User {
	tgUser := models.User{}
	if tgBot.Message != nil {
		tgUser = *tgBot.Message.From
	}
	if tgBot.CallbackQuery != nil {
		tgUser = tgBot.CallbackQuery.From
	}
	if tgBot.InlineQuery != nil {
		tgUser = *tgBot.InlineQuery.From
	}
	if tgBot.MyChatMember != nil {
		tgUser = tgBot.MyChatMember.From
	}

	return tgUser
}

func (tgBot *EsTgBot) setUser(userModel *merchantuser.Model) (bool, error) {
	tgUser := tgBot.getTgUser()

	tgBot.UserModel = userModel
	needToSaveSecondModel := false
	if tgBot.UserModel.GetId() == 0 {
		tgBot.UserModel.TgId = tgUser.ID
		tgBot.UserModel.FirstName = tgUser.FirstName
		tgBot.UserModel.LastName = tgUser.LastName
		tgBot.UserModel.Username = tgUser.Username
		tgBot.UserModel.IsPremium = tgUser.IsPremium
		tgBot.UserModel.IsActive = true
		needToSaveSecondModel = true

		if utils.InArray(tgUser.ID, constants.AdminTelegramIds) {
			tgBot.UserModel.IsApproved = true
		}
	}

	if tgBot.UserModel.FirstName != tgUser.FirstName {
		tgBot.UserModel.FirstName = tgUser.FirstName
		needToSaveSecondModel = true
	}

	if tgBot.UserModel.LastName != tgUser.LastName {
		tgBot.UserModel.LastName = tgUser.LastName
		needToSaveSecondModel = true
	}

	if tgBot.UserModel.Username != tgUser.Username {
		tgBot.UserModel.Username = tgUser.Username
		needToSaveSecondModel = true
	}

	if tgBot.UserModel.IsPremium != tgUser.IsPremium {
		tgBot.UserModel.IsPremium = tgUser.IsPremium
		needToSaveSecondModel = true
	}

	if !tgBot.UserModel.IsActive {
		tgBot.UserModel.IsActive = true
		needToSaveSecondModel = true
	}

	if needToSaveSecondModel {
		tgBot.UserModel.Save(tgBot.UserModel)
	}

	return true, nil
}
