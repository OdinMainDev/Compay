package tgbot

func (tgBot *EsTgBot) handleMyChatMember() {
	// Запрещаем в чатах использовать бота
	if tgBot.MyChatMember.Chat.ID <= 0 {
		return
	}

	memberType := tgBot.MyChatMember.NewChatMember.Type
	if memberType == "kicked" || memberType == "left" {
		tgBot.UserModel.IsActive = false
		tgBot.UserModel.Save(tgBot.UserModel)
	}
}
