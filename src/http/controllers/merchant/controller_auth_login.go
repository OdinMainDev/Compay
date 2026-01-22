package merchant

import (
	"compay/common/constants"
	"compay/common/cperrors"
	"compay/common/types"
	"compay/src/helpers"
	"compay/src/models/logauth"
	"compay/src/models/notification"
	"compay/src/models/user"
	"efw/core/efwgin"
	"efw/core/efwgin/casts"
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

	if userModel == nil || userModel.GetId() == 0 {
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

	go afterAuthLogin(context.Copy(), userModel)
	efwginhelpers.ResponseSuccess(context, gin.H{"token": jwt})
}

func afterAuthLogin(context *gin.Context, userModel *user.Model) {
	var merchantId types.MerchantId = 0
	if userModel.Merchant != nil {
		merchantId = userModel.Merchant.GetId()
	}

	logauth.FastCreate(
		userModel.GetId(),
		merchantId,
		constants.LogAuthLogin,
		casts.GetRequestUri(context),
		efwgin.ConstructClientData(context),
	)

	notification.FastCreate(merchantId, constants.NotificationActionTest, constants.NotificationTextTest("login"))
}
