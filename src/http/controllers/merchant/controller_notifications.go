package merchant

import (
	"compay/common/constants"
	"compay/common/types"
	"compay/src/models/notification"
	"compay/src/models/user"
	"efw/components/model"
	"efw/core/efwgin"
	efwginhelpers "efw/core/efwgin/efwgin_helpers"
	"efw/services/efwcontroller"
	"efw/services/efwdt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type requestNotificationsList struct {
	Id     types.NotificationId `json:"id" efwdt:"id" binding:"omitempty,gt=0"`
	IsRead *bool                `json:"is_read" efwdt:"is_read" binding:"omitempty"`
	efwdt.WithCreatedAt
}

func NotificationsList(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	efwcontroller.Datable[notification.Model, requestNotificationsList](
		context,
		[]string{"id", "action", "is_read", "created_at"},
		[]efwdt.ChainFilter[notification.Model]{
			func(chain gorm.ChainInterface[notification.Model]) gorm.ChainInterface[notification.Model] {
				return chain.Where("merchant_id = ?", userModel.MerchantId)
			},
		},
		nil,
	)
}

func NotificationsListFilters(context *gin.Context) {
	efwcontroller.DatatableFilters(context, []efwdt.Filter{
		{Name: "id", Type: types.EFWDT_FilterTypeInteger},
		{Name: "is_read", Type: types.EFWDT_FilterTypeBoolean},
		{Name: "created_at", Type: types.EFWDT_FilterTypeDate},
	})
}

func NotificationsGetUnread(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	efwginhelpers.ResponseSuccess(context, gin.H{"data": notification.GetModels(notification.QueryHelper{
		IsUnread:   true,
		MerchantId: userModel.MerchantId,
		QueryHelper: model.QueryHelper{
			HasPagination: model.HasPagination{
				HasLimit: true,
			},
		},
	})})
}

func NotificationsRead(context *gin.Context) {
	var input efwcontroller.RequestWithId
	if !efwgin.ShouldBindUri(context, &input) {
		return
	}

	userModel := context.MustGet(constants.GinUser).(*user.Model)
	notificationModel := notification.GetModel(notification.QueryHelper{
		Id:         types.NotificationId(input.Id),
		MerchantId: userModel.MerchantId,
		IsUnread:   true,
	})

	if notificationModel == nil {
		efwginhelpers.ResponseError(context, nil, http.StatusNotFound)
		return
	}

	notificationModel.IsRead = true
	notification.Save(notificationModel, notification.QueryHelper{})

	createLogUser(context, userModel, constants.LogUserNotifyRead)
	efwginhelpers.ResponseSuccess(context, nil)
}

func NotificationsReadAll(context *gin.Context) {
	userModel := context.MustGet(constants.GinUser).(*user.Model)
	rowsAffected := notification.ReadAll(userModel.MerchantId)

	if rowsAffected != 0 {
		createLogUser(context, userModel, constants.LogUserNotifyReadAll)
	}

	efwginhelpers.ResponseSuccess(context, nil)
}
