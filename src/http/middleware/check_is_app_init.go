package middleware

import (
	"efw/config"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIsAppInit() gin.HandlerFunc {
	return func(context *gin.Context) {
		if config.Cryptkey == "" || config.Cryptiv == "" || config.CrypSalt == "" {
			efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
			context.Abort()
			return
		}

		context.Next()
	}
}
