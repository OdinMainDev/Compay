package merchantwallet

import (
	"compay/src/models/currency"

	"github.com/gin-gonic/gin"
)

func DTO(model *Model) any {
	return gin.H{
		"id":              model.Id,
		"merchant_id":     model.MerchantId,
		"payment_id":      model.PaymentId,
		"currency":        model.Currency,
		"address":         model.Address,
		"amount":          currency.AmountToFloatByCodeAlpha(model.Amount, model.Currency),
		"received_amount": currency.AmountToFloatByCodeAlpha(model.ReceivedAmount, model.Currency),
		"sent_amount":     currency.AmountToFloatByCodeAlpha(model.SentAmount, model.Currency),
		"created_at":      model.CreatedAt,
		"updated_at":      model.UpdatedAt,
	}
}
