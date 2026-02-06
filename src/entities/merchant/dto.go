package merchant

import (
	"github.com/gin-gonic/gin"
)

func DTO(model *Entity) gin.H {
	return gin.H{
		"name": model.Name,
		"settings": gin.H{
			"notification_settings": gin.H{
				"telegram_chat_id": model.Settings.NotificationSettings.TelegramChatId,
			},
		},
	}
}
