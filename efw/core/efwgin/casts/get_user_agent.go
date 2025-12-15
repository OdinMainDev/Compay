package casts

import "github.com/gin-gonic/gin"

func GetClientUserAgent(context *gin.Context) string {
	return context.GetHeader("User-Agent")
}
