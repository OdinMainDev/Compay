package notification

import (
	"compay/common/types"
	"compay/src/entities/notification"
	"context"
	"efw"
	"efw/core/efwlog"
)

func FastCreate(merchantId types.MerchantId, action types.NotificationAction, text types.NotificationText) {
	model := &Model{}
	model.MerchantId = merchantId
	model.IsRead = false
	model.Action = action
	model.Text = text

	Save(model, QueryHelper{})
}

func ReadAll(merchantId types.MerchantId) int {
	res, err := efw.DBTransaction[any]().
		Where("merchant_id = ?", merchantId).
		Where("is_read = ?", false).
		Updates(context.Background(), Model{Entity: notification.Entity{IsRead: true}})

	if err != nil {
		efwlog.LogByLoggerName("gorm_error", 0, efwlog.FmtGormError(0, 0, err.Error(), "ReadAll"))
		return 0
	}
	return res
}
