package logpayment

import (
	"compay/common/types"
	"compay/src/entities/logpayment"
)

func FastCreate(
	paymentId types.PaymentId,
	paymentTrxId string,
	merchantId types.MerchantId,
	txData types.TxData,
) {
	model := &Model{}
	model.PaymentId = paymentId
	model.PaymentTrxId = paymentTrxId
	model.MerchantId = merchantId
	model.TxId = txData.TxId
	model.TxHash = txData.TxHash
	model.Address = txData.Address
	model.Data = logpayment.Data{
		TxData: txData,
	}
	Save(model, QueryHelper{})
}
