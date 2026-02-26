package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/logauth"
	"compay/src/models/loguser"
	"compay/src/models/merchant"
	"compay/src/models/notification"
	"compay/src/models/user"
	"efw/components/entity"
	"efw/core/efwgin"
	"efw/core/efwgin/casts"
	"efw/utils"

	"github.com/gin-gonic/gin"
)

type requestSearch struct {
	Search string `json:"search" binding:"omitempty,max=255"`
}

type requestTotp struct {
	Code string `json:"code" binding:"required,len=6"`
}

func isUserRootAdmin(userId entity.UserId) bool {
	return utils.InArray(userId, constants.AdminIds)
}

func isAdminCanUpdateUser(adminUserModel, userModel *user.Model) bool {
	if isUserRootAdmin(adminUserModel.GetId()) {
		return true
	}
	return adminUserModel.GetId() == userModel.CreatorId
}

func isUserCanUpdateMerchant(userModel *user.Model, merchantModel *merchant.Model) bool {
	return userModel.GetId() == merchantModel.CreatorId
}

func createNotify(userModel *user.Model, action types.NotificationAction, text types.NotificationText) {
	var merchantId types.MerchantId = 0
	if userModel.Merchant != nil {
		merchantId = userModel.Merchant.GetId()
	}

	notification.FastCreate(merchantId, action, text)
}

func createLogAuth(context *gin.Context, userModel *user.Model, action types.LogAuthAction) {
	var merchantId types.MerchantId = 0
	if userModel.Merchant != nil {
		merchantId = userModel.Merchant.GetId()
	}

	logauth.FastCreate(
		userModel.GetId(),
		merchantId,
		action,
		casts.GetRequestUri(context),
		efwgin.ConstructClientData(context),
	)
}

func createLogUser(context *gin.Context, userModel *user.Model, action types.LogUserAction) {
	var merchantId types.MerchantId = 0
	if userModel.Merchant != nil {
		merchantId = userModel.Merchant.GetId()
	}

	loguser.FastCreate(
		userModel.GetId(),
		merchantId,
		action,
		casts.GetRequestUri(context),
		casts.GetRequestBodyString(context),
		efwgin.ConstructClientData(context),
	)
}
