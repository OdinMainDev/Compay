package efwmiddleware

import (
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwjwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const GinIsAuth = "_efw/gin/is_auth"

func CheckAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, ok := context.Get(GinIsAuth)
		if ok {
			context.Next()
			return
		}

		err := efwjwt.ValidateJWT(context)
		if err != nil {
			efwginhelpers.ResponseError(context, err, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Set(GinIsAuth, true)
		context.Next()
	}
}
