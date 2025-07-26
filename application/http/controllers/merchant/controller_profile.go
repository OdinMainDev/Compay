package merchant

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/merchantuser"
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"compay/application/service/totp"
	"compay/common/constants"

	"github.com/gin-gonic/gin"
)

func GetMe(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser)
	helpers.GinSuccessResponse(context, gin.H{"data": userModel})
}

func GetMyHas2Fa(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*merchantuser.UserModel)
	helpers.GinSuccessResponse(context, gin.H{"result": userModel.IsTwofaEnabled})
}

func GetMy2FaUri(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*merchantuser.UserModel)

	if userModel.IsTwofaEnabled {
		helpers.GinErrorResponse(context, "AlreadyConfigured")
		return
	}

	issuer := "Compay(WIP)"

	secret := ""
	if userModel.TotpToken == "" {
		scr, err := totp.GenerateSecret(userModel.GetUserUsername(), issuer)
		if err != nil {
			helpers.GinErrorResponse(context, err.Error())
			return
		}
		secret = scr
		userModel.TotpToken = basemodel.EncryptedString(secret)
		userModel.Save(userModel)
	} else {
		secret = string(userModel.TotpToken)
	}
	helpers.GinSuccessResponse(context, gin.H{"uri": totp.OtpUri(secret, userModel.GetUserUsername(), issuer)})
}

func SetMy2Fa(context *gin.Context) {
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
	if userModel.IsTwofaEnabled {
		helpers.GinErrorResponse(context, "AlreadyConfigured")
		return
	}

	ok, err := totp.ValidateCode(string(userModel.TotpToken), input.Code)
	result_message := ""
	if err != nil {
		result_message = err.Error()
	}
	if ok {
		userModel.IsTwofaEnabled = true
		userModel.Save(userModel)
	}
	helpers.GinSuccessResponse(context, gin.H{"result": ok, "result_message": result_message})
}
