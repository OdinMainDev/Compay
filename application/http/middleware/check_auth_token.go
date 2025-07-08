package middleware

import (
	"compay/application/helpers"
	"compay/application/service/jwt"
	"compay/common/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuthToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModel, err := checkUserToken(context)
		if err != nil {
			helpers.GinErrorResponse(context, err, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Set(constants.GinUserId, userModel.GetId())
		context.Set(constants.GinUser, userModel)

		if err := checkAuthToken(context); err != nil {
			helpers.GinErrorResponse(context, err, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Next()
	}
}

func checkAuthToken(context *gin.Context) error {
	return jwt.ValidateJWT(context)
}
