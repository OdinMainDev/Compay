package wallet

import (
	"compay/application/entities/user"
	"compay/application/helpers"
	"compay/application/service/cryptowallet"
	"compay/application/service/esvalidator"
	"compay/common/constants"
	"compay/core/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type requestAddNewWallet struct {
	Seed string `json:"seed" binding:"omitempty,seedPhrase"`
}

func AddNewWallet(context *gin.Context) {
	var input requestAddNewWallet
	if err := context.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		helpers.GinErrorResponse(context, esvalidator.ValidatorResponse(err))
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)

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

	helpers.GinSuccessResponse(context, gin.H{"wallet_token": walletTokenEncrypted})
}
