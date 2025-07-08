package middleware

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/user"
	"compay/application/helpers"
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/core/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUserToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		userModel, err := checkUserToken(context)
		if err != nil {
			helpers.GinErrorResponse(context, err, http.StatusUnauthorized)
			context.Abort()
			return
		}

		context.Set(constants.GinUserId, userModel.GetId())
		context.Set(constants.GinUser, userModel)
		context.Next()
	}
}

func checkUserToken(context *gin.Context) (*user.Model, error) {
	userTokenEnc := context.Request.Header.Get("U-token")
	if userTokenEnc == "" {
		return nil, cperrors.ErrorAuthUserTokenIsNull
	}

	userToken, err := utils.DecryptString(userTokenEnc)
	if err != nil {
		return nil, cperrors.ErrorAuthUserTokenIsInvalid
	}

	userModel := user.GetModel(user.QueryHelper{
		Token: basemodel.EncryptedString(userToken),
	})

	if userModel == nil || userModel.GetId() == 0 {
		return nil, cperrors.ErrorAuthBadUserToken
	}

	if userModel.IsBanned {
		return nil, cperrors.ErrorAuthUserBanned
	}

	return userModel, nil
}
