package merchant

import (
	"github.com/gin-gonic/gin"
)

func DTO(model *Model) gin.H {
	return gin.H{
		"name": model.Name,
		"settings": gin.H{
			"notification_settings": gin.H{
				"notification_settings": model.Settings.NotificationSettings.TelegramChatId,
			},
		},
	}
}
