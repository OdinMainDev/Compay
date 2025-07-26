package merchant

import (
	"compay/application/entities/merchantuser"
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"compay/application/service/totp"
	"compay/common/constants"

	"github.com/gin-gonic/gin"
)

type requestTotp struct {
	Code string `json:"code" binding:"required,len=6"`
}

func AuthTotp(context *gin.Context) {
	var input requestTotp

	if err := context.ShouldBindJSON(&input); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userModel := context.MustGet(constants.GinUser).(*merchantuser.UserModel)
	if userModel.TotpToken == "" {
		helpers.GinErrorResponse(context, "NotConfigured")
		return
	}
	if !userModel.IsTwofaEnabled {
		helpers.GinErrorResponse(context, "NotConfigured")
		return
	}

	ok, err := totp.ValidateCode(string(userModel.TotpToken), input.Code)
	if !ok || err != nil {
		helpers.GinErrorResponse(context, "InvalidCode")
		return
	}

	token, err := totp.GenerateHeaderToken(userModel)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"token": token})
}
