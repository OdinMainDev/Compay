package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/helpers"
	"compay/src/models/user"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/core/efwjwt"

	"github.com/gin-gonic/gin"
)

type requestLogin struct {
	Email    string `json:"email" binding:"required,email,min=5,max=255"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

func AuthLogin(context *gin.Context) {
	var input requestLogin
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := user.GetModel(user.QueryHelper{
		Username:            input.Email,
		DontCheckMerchantId: true,
		WithMerchant:        true,
	})

	if userModel == nil || userModel.IsBanned {
		efwginhelpers.ResponseError(context, cperrors.ErrorAuthUserNotFound)
		return
	}

	if err := userModel.ValidatePassword(input.Password); err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	jwtData := helpers.ConstructJwtData(userModel)
	jwt, err := efwjwt.GenerateJWT(jwtData, userModel.SecretToken, constants.JwtMerchantTTL)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	go asyncAuthLogin(context.Copy(), userModel)
	efwginhelpers.ResponseSuccess(context, gin.H{"token": jwt})
}

func asyncAuthLogin(context *gin.Context, userModel *user.Model) {
	createLogAuth(context, userModel, constants.LogAuthLogin)
}
