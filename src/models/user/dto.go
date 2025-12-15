package user

import (
	"github.com/gin-gonic/gin"
)

func DTO(model *Model) gin.H {
	merchantData := gin.H{}
	if model.Merchant != nil {
		// заполняем merchantData
	}
	return gin.H{
		"username":         model.Username,
		"is_twofa_enabled": model.IsTwofaEnabled,
		"merchant":         merchantData,
	}
}
