package merchant

import (
	"compay/common/constants"
	"compay/src/models/merchant"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestMerchantSettingsProcess struct {
	NotificationTelegramChatId int64 `json:"notification_telegram_chat_id" binding:"omitempty"`
}

func MerchantSettingsGet(context *gin.Context) {
	merchantModel := context.MustGet(constants.GinMerchant).(*merchant.Model)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchant.DTO(merchantModel)})
}

func MerchantSettingsProcess(context *gin.Context) {
	var input requestMerchantSettingsProcess
	if !efwgin.ShouldBindBodyWithJSON(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)
	merchantModel := context.MustGet(constants.GinMerchant).(*merchant.Model)

	if !isUserCanUpdateMerchant(userModel, merchantModel) {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	merchantModel.Settings.PublicSettings.NotificationTelegramChatId = input.NotificationTelegramChatId
	merchant.Save(merchantModel, merchant.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserMerchantSettingsProcess)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchant.DTO(merchantModel)})
}

func MerchantSettingsActionGetToken(context *gin.Context) {
	merchantModel := context.MustGet(constants.GinMerchant).(*merchant.Model)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": merchantModel.Token})
}

func MerchantSettingsActionGenerateSecret(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)

	secretToken := uuid.NewString()
	merchantModel := context.MustGet(constants.GinMerchant).(*merchant.Model)
	merchantModel.SecretKey = entity.EncryptedString(secretToken)
	merchant.Save(merchantModel, merchant.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserMerchantSettingsGenerateSecret)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": secretToken})
}
