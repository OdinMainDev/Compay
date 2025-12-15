package middleware

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwtotp"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAuthTotp() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModelAny, ok := context.Get(constants.GinUser)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckAuthTotpNilUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userModel, ok := userModelAny.(*user.Model)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckAuthTotpInvalidTypeUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		if !userModel.IsTwofaEnabled {
			context.Next()
			return
		}

		if err := efwtotp.Validate(context, userModel.Username); err != nil {
			if errors.Is(err, efwtotp.ErrorTotpTokenExpired) {
				efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorAuthTotpExpired))
			} else {
				efwginhelpers.ResponseError(context, err, http.StatusUnauthorized)
			}
			context.Abort()
			return
		}
		context.Next()
	}
}
