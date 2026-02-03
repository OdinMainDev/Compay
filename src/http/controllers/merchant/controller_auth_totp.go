package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwtotp"

	"github.com/gin-gonic/gin"
)

func AuthTotp(context *gin.Context) {
	var input requestTotp
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)
	if userModel.TotpToken == "" {
		efwginhelpers.ResponseError(context, cperrors.ErrorProfileTotpNotConfigured)
		return
	}
	if userModel.IsTwofaEnabled {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorProfileTotpAlreadyEnabled))
		return
	}

	err := efwtotp.ValidateCode(string(userModel.TotpToken), input.Code)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	token, err := efwtotp.GenerateHeaderToken(userModel.Username)
	if err != nil {
		efwginhelpers.ResponseError(context, err)
		return
	}

	go asyncAuthTotp(context.Copy(), userModel)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": token})
}

func asyncAuthTotp(context *gin.Context, userModel *user.Model) {
	createLogAuth(context, userModel, constants.LogAuthTotp)
}
