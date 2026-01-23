package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/src/models/user"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"

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

	if input.NewUsername != "" {
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
