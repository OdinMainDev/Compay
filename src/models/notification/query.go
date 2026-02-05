package notification

import (
	"compay/common/types"
	"compay/src/entities/notification"
	"context"
	"efw"
	"efw/components/model"
	"efw/core/efwlog"

	"gorm.io/gorm"
)

type QueryHelper struct {
	model.QueryHelper
	Id         types.NotificationId
	MerchantId types.MerchantId
	IsRead     bool
	IsUnread   bool
}

func (helper QueryHelper) ApplyDBFilters(chain gorm.ChainInterface[*Model]) gorm.ChainInterface[*Model] {
	chain = chain.Where("merchant_id = ?", helper.MerchantId)

	if helper.Id != 0 {
		chain = chain.Where("id = ?", helper.Id)
	}
	if helper.IsRead {
		chain = chain.Where("is_read = ?", true)
	}
	if helper.IsUnread {
		chain = chain.Where("is_read = ?", false)
	}
	return chain.Order("id DESC")
}

func ReadAll(merchantId types.MerchantId) int {
	res, err := efw.DBTransaction[notification.Entity]().
		Where("merchant_id = ?", merchantId).
		Where("is_read = ?", false).
		Updates(context.Background(), notification.Entity{IsRead: true})

	if err != nil {
		efwlog.LogByLoggerName("gorm_error", 0, efwlog.FmtGormError(0, 0, err.Error(), "ReadAll"))
		return 0
	}
	return res
}
