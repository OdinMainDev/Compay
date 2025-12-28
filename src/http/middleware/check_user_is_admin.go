package middleware

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIsUserAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModelAny, ok := context.Get(constants.GinUser)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckIsUserAdminNilUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userModel, ok := userModelAny.(*user.Model)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckIsUserAdminInvalidTypeUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		if userModel.MerchantId != 0 {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckIsUserAdminPermissionDenied, http.StatusUnauthorized)
			context.Abort()
			return
		}
		context.Next()
	}
}
