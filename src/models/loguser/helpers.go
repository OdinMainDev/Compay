package loguser

import (
	"compay/common/types"
	"efw/components/entity"
)

func FastCreate(
	userId entity.UserId,
	merchantId types.MerchantId,
	action types.LogUserAction,
	url entity.Url,
	request string,
	clientData entity.ClientData,
) {
	model := &Model{}
	model.UserId = userId
	model.MerchantId = merchantId
	model.Action = action
	model.Url = url
	model.Request = request
	model.ClientData = clientData

	Save(model, QueryHelper{})
}
