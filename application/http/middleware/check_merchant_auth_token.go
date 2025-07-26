package middleware

import (
	"compay/application/helpers"
	"compay/application/service/jwt"
	"compay/application/service/totp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckMerchantAuth(args ...bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := jwt.ValidateMerchantJWT(context); err != nil {
			helpers.GinErrorResponse(context, err, http.StatusUnauthorized)
			context.Abort()
			return
		}
		if len(args) > 0 && args[0] {
			context.Next()
			return
		}
		CheckTotp(context)
	}
}

func CheckTotp(context *gin.Context) {
	if err := totp.ValidateRequest(context); err != nil {
		helpers.GinErrorResponse(context, err)
		context.Abort()
		return
	}
	context.Next()
}
