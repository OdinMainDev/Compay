package logmerchantwallet

import (
	"efw/core/efwamount"

	"github.com/gin-gonic/gin"
)

func DTO(model *Entity) gin.H {
	return gin.H{
		"merchant_id":        model.MerchantId,
		"merchant_wallet_id": model.MerchantWalletId,
		"payment_id":         model.PaymentId,
		"withdraw_id":        model.WithdrawId,
		"payment_trx_id":     model.PaymentTrxId,
		"withdraw_trx_id":    model.WithdrawTrxId,
		"currency":           model.Currency,
		"amount":             efwamount.ToFloatByCodeAlpha(model.Amount, model.Currency),
		"amount_usd":         efwamount.ToFloat(model.AmountUsd, 2),
		"created_at":         model.CreatedAt,
		"updated_at":         model.UpdatedAt,
	}
}
