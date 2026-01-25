package logauth

import (
	"compay/common/types"
	"efw/components/entity"
)

func FastCreate(
	userId entity.UserId,
	merchantId types.MerchantId,
	action types.LogAuthAction,
	url entity.Url,
	clientData entity.ClientData,
) {
	model := &Model{}
	model.UserId = userId
	model.MerchantId = merchantId
	model.Action = action
	model.Url = url
	model.Ip = clientData.Ip
	model.FingerPrint = clientData.Fingerprint
	model.UserAgent = clientData.UserAgent
	model.ClientData = clientData

	Save(model, QueryHelper{})
}
