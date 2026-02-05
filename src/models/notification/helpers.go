package notification

import (
	"compay/common/types"
)

func FastCreate(merchantId types.MerchantId, action types.NotificationAction, text types.NotificationText) {
	model := &Model{}
	model.MerchantId = merchantId
	model.IsRead = false
	model.Action = action
	model.Text = text

	Save(model, QueryHelper{})
}
