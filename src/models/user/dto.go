package user

import (
	"github.com/gin-gonic/gin"
)

func DTO(model *Model) any {
	return gin.H{
		"username":         model.Username,
		"is_twofa_enabled": model.IsTwofaEnabled,
	}
}
