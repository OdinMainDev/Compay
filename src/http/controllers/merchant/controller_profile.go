package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwtotp"

	"github.com/gin-gonic/gin"
)

type requestProfileProcess struct {
	OldUsername string `json:"old_username" binding:"omitempty,min=6,max=255"`
	NewUsername string `json:"new_username" binding:"omitempty,required_with=OldUsername,min=6,max=255"`
	OldPassword string `json:"old_password" binding:"omitempty,min=6,max=255"`
	NewPassword string `json:"new_password" binding:"omitempty,required_with=OldPassword,min=6,max=255"`
}

func ProfileGet(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.DTO(userModel)})
}

func ProfileProcess(context *gin.Context) {
	var input requestProfileProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}
	userModel := context.MustGet(constants.GinUser).(*user.Model)

	if input.OldUsername != "" && input.OldUsername != userModel.Username {
		efwginhelpers.ResponseError(context, cperrors.ErrorProfileInvalidOldUsername)
		return
	}

	hasChanges := false
	if input.OldPassword != "" {
		if err := userModel.ValidatePassword(input.OldPassword); err != nil {
			efwginhelpers.ResponseError(context, err)
			return
		}
		userModel.NewPassword = input.NewPassword
		hasChanges = true
	}

	if input.NewUsername != "" && input.NewUsername != userModel.Username {
		checkUserExist := user.GetModel(user.QueryHelper{Username: input.NewUsername, DontCheckMerchantId: true})
		if checkUserExist != nil {
			efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorProfileNewUsernameIsTaken))
			return
		}
		userModel.Username = input.NewUsername
		hasChanges = true
	}

	if hasChanges {
		user.Save(userModel, user.QueryHelper{})
		createLogUser(context, userModel, constants.LogUserProfileUpdate)
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": user.DTO(userModel)})
}

func ProfileTotpUri(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	if userModel.IsTwofaEnabled {
		efwginhelpers.ResponseError(context, efwginhelpers.NotHiddenError(cperrors.ErrorProfileTotpAlreadyConfigured))
		return
	}

	issuer := "Compay"
	if userModel.TotpToken == "" {
		secret, err := efwtotp.GenerateSecret(userModel.Username, issuer)
		if err != nil {
			efwginhelpers.ResponseError(context, err)
			return
		}
		userModel.TotpToken = entity.EncryptedString(secret)
		user.Save(userModel, user.QueryHelper{})
	}

	efwginhelpers.ResponseSuccess(context, gin.H{"data": efwtotp.ConstructUri(string(userModel.TotpToken), userModel.Username, issuer)})
}

func ProfileTotpActivate(context *gin.Context) {
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

	userModel.IsTwofaEnabled = true
	user.Save(userModel, user.QueryHelper{})

	efwginhelpers.ResponseSuccess(context, nil)
}
