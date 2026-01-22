package casts

import (
	"efw/components/entity"

	"github.com/gin-gonic/gin"
)

func GetClientUserAgent(context *gin.Context) entity.UserAgent {
	return entity.UserAgent(context.GetHeader("User-Agent"))
}
