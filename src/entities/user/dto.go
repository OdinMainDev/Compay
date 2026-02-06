package user

import "github.com/gin-gonic/gin"

func DTO(model *Entity) gin.H {
	return gin.H{
		"username":         model.Username,
		"is_twofa_enabled": model.IsTwofaEnabled,
	}
}
