package middleware

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	"efw/components/entity"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwjwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, ok := context.Get(constants.GinUser)
		if ok {
			context.Next()
			return
		}

		jwtClaimsDataAny, ok := context.Get(efwjwt.GinJwtClaimsData)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserNilJwtData, http.StatusUnauthorized)
			context.Abort()
			return
		}

		jwtClaimsData, ok := jwtClaimsDataAny.(efwjwt.JwtClaimsData)
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserInvalidJwtData, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userIdStr, ok := jwtClaimsData[constants.JwtClaimsKeyUserId]
		if !ok {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserNilUserId, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserInvalidUserId, http.StatusUnauthorized)
			context.Abort()
			return
		}

		userModel := user.GetModel(user.QueryHelper{
			Id:                  entity.UserId(userId),
			DontCheckMerchantId: true,
			WithMerchant:        true,
		})

		if userModel == nil {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserInvalidUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		if userModel.IsBanned {
			efwginhelpers.ResponseError(context, cperrors.ErrorMiddlewareCheckUserBannedUserModel, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Set(constants.GinUser, userModel)
		context.Next()
	}
}
