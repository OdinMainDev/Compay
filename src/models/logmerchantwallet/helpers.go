package logmerchantwallet

import (
	"compay/common/types"
	"efw/components/entity"
)

func FastCreate(
	merchantId types.MerchantId,
	merchantWalletId types.MerchantWalletId,
	paymentId types.PaymentId,
	withdrawId types.PaymentId,
	paymentTrxId string,
	withdrawTrxId string,
	currency entity.CurrencyCodeAlpha,
	amount entity.Amount,
	amountUsd entity.Amount,
) {
	model := &Model{}
	model.MerchantId = merchantId
	model.MerchantWalletId = merchantWalletId
	model.PaymentId = paymentId
	model.WithdrawId = withdrawId
	model.PaymentTrxId = paymentTrxId
	model.WithdrawTrxId = withdrawTrxId
	model.Currency = currency
	model.Amount = amount
	model.AmountUsd = amountUsd
	Save(model, QueryHelper{})
}
