package middleware

import (
	"compay/application/entities/merchantuser"
	"compay/application/helpers"
	"compay/common/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckMerchantIsAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModelData, ok := context.Get(constants.GinUser)
		if !ok {
			helpers.GinErrorResponse(context, nil, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userModel, ok := userModelData.(*merchantuser.Model)
		if !ok {
			helpers.GinErrorResponse(context, nil, http.StatusUnauthorized)
			context.Abort()
			return
		}

		if *userModel.GetMerchantId() != 0 {
			helpers.GinErrorResponse(context, nil, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Next()
	}
}
