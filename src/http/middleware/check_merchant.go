package middleware

import "github.com/gin-gonic/gin"

func CheckMerchant() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
