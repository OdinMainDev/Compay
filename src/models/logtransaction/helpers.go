package logtransaction

import (
	"compay/common/types"
	"compay/src/entities/logtransaction"
)

func FastCreate(
	paymentId types.PaymentId,
	paymentTrxId string,
	withdrawId types.WithdrawId,
	withdrawTrxId string,
	merchantId types.MerchantId,
	txData types.TxData,
	txWithdrawData types.TxWithdrawData,
) {
	model := &Model{}
	model.PaymentId = paymentId
	model.PaymentTrxId = paymentTrxId
	model.WithdrawId = withdrawId
	model.WithdrawTrxId = withdrawTrxId
	model.MerchantId = merchantId
	model.TxId = txData.TxId
	model.TxHash = txData.TxHash
	model.Address = txData.Address
	model.Data = logtransaction.Data{
		TxData:         txData,
		TxWithdrawData: txWithdrawData,
	}
	Save(model, QueryHelper{})
}
