package wallet

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/user"
	"compay/application/helpers"
	"compay/application/service/esvalidator"
	"compay/application/service/jwt"
	"compay/common/constants"
	"compay/common/cperrors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AuthCheckUser(context *gin.Context) {
	helpers.GinSuccessResponse(context, nil)
}

func AuthCheckToken(context *gin.Context) {
	helpers.GinSuccessResponse(context, nil)
}

func AuthCheckPinCode(context *gin.Context) {
	var input requestAuth
	if err := context.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

	if userModel.PinCode != basemodel.EncryptedString(input.PinCode) {
		helpers.GinErrorResponse(context, cperrors.ErrorAuthJwtTokenMismatchPinCodes)
		return
	}

	authToken, err := jwt.GenerateAuthToken(userModel)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"auth_token": authToken})
}
