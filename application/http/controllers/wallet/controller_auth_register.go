package wallet

import (
	"compay/application/entities/basemodel"
	"compay/application/entities/user"
	"compay/application/helpers"
	"compay/application/service/cryptowallet"
	"compay/application/service/esvalidator"
	"compay/application/service/jwt"
	"compay/core/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AuthRegister(context *gin.Context) {
	var input requestAuth
	if err := context.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userModel := &user.Model{
		Token:       basemodel.EncryptedString(utils.GenerateRandomSalt()),
		PinCode:     basemodel.EncryptedString(input.PinCode),
		IsBanned:    false,
		SecretToken: utils.GenerateRandomSalt(),
	}
	userModel.Save(userModel)

	walletToken := ""

	if input.Seed != "" {
		walletModel, err := cryptowallet.RestoreWallet(input.Seed, userModel.GetId())
		if err != nil {
			helpers.GinErrorResponse(context, err)
			return
		}

		walletToken = fmt.Sprint(walletModel.GetId())
	} else {
		walletModel, err := cryptowallet.CreateWallet(userModel.GetId())
		if err != nil {
			helpers.GinErrorResponse(context, err)
			return
		}

		walletToken = fmt.Sprint(walletModel.GetId())
	}

	walletTokenEncrypted, err := utils.EncryptString(walletToken)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	authToken, err := jwt.GenerateAuthToken(userModel)
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	userToken, err := utils.EncryptString(string(userModel.Token))
	if err != nil {
		helpers.GinErrorResponse(context, err)
		return
	}

	helpers.GinSuccessResponse(context, gin.H{"user_token": userToken, "auth_token": authToken, "wallet_token": walletTokenEncrypted})
}
