package logmerchantwallet

import (
	"compay/common/types"
)

func FastCreate(
	merchantId types.MerchantId,
	paymentId types.PaymentId,
	withdrawId types.PaymentId,
	paymentTrxId string,
	withdrawTrxId string,
	currency types.CurrencyCodeAlpha,
	amount types.Amount,
	amountUsd types.Amount,
) {
	model := &Model{}
	model.MerchantId = merchantId
	model.PaymentId = paymentId
	model.WithdrawId = withdrawId
	model.PaymentTrxId = paymentTrxId
	model.WithdrawTrxId = withdrawTrxId
	model.Currency = currency
	model.Amount = amount
	model.AmountUsd = amountUsd
	Save(model, QueryHelper{})
}
