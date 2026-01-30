package middleware

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/merchant"
	"compay/src/models/user"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckMerchant() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModelAny, ok := context.Get(constants.GinUser)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckMerchantNilUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userModel, ok := userModelAny.(*user.Model)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckMerchantInvalidTypeUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		merchantModel := merchant.GetModel(merchant.QueryHelper{
			Id: userModel.MerchantId,
		})
		if merchantModel == nil || !merchantModel.IsEnabled {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckMerchantIsDisabled, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Set(constants.GinMerchant, merchantModel)
		context.Next()
	}
}
